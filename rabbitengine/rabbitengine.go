package rabbitengine

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/sync/semaphore"
)

// MessageBody is the struct for the body passed in the AMQP message. The type will be set on the Request header
type MessageBody struct {
	Data []byte
}

// Message is the amqp request to publish
type Message struct {
	Queue         string
	ReplyTo       string
	ContentType   string
	CorrelationID string
	Priority      uint8
	Body          MessageBody
	Header        map[string]interface{}
}

// Exchange Config
type ExchangeConfig struct {
	Name string
	Type string
}

// Connection config
type ConnectionConfig struct {
	User    string
	Pass    string
	Address string
}

// Queue config
type QueueConfig struct {
	Name    string
	BindKey string
}

// Connection is the connection created
type Connection struct {
	name     string
	config   *ConnectionConfig
	conn     *amqp.Connection
	channel  *amqp.Channel
	exchange *ExchangeConfig
	queues   []*QueueConfig
	err      chan error
	args     map[string]interface{}
}

type RabbitMQ struct {
	Host         string `envconfig:"RABBITMQ_HOST" default:"localhost"`
	User         string `envconfig:"RABBITMQ_USER" default:""`
	Password     string `envconfig:"RABBITMQ_PASSWORD" default:""`
	Port         string `envconfig:"RABBITMQ_PORT" default:"5672"`
	Publisher    string `envconfig:"RABBITMQ_PUBLISHER" default:"KEKASI"`
	Exchange     string `envconfig:"RABBITMQ_EXCHANGE" default:"KEKASI_QUEUE"`
	Queue        string `envconfig:"RABBITMQ_QUEUE" default:"KEKASI_QUEUE"`
	QueueBindKey string `envconfig:"RABBITMQ_QUEUE_BIND_KEY" default:"KEKASI_QUEUE"`
}

func (p RabbitMQ) ConnConfig() *ConnectionConfig {
	data := ConnectionConfig{
		User:    p.User,
		Pass:    p.Password,
		Address: p.Host + ":" + p.Port,
	}
	return &data
}

// NewConnection returns the new connection object
func NewConnection(name string, conf RabbitMQ, exchange *ExchangeConfig, queues []*QueueConfig, args map[string]interface{}) *Connection {
	c := &Connection{
		name:     name,
		exchange: exchange,
		config:   conf.ConnConfig(),
		queues:   queues,
		err:      make(chan error),
		args:     args,
	}
	return c
}

// Connect create a new connection to RabbitMQ
func (c *Connection) Connect() error {
	var err error

	connName := os.Getenv("RABBIT_CONNECTION_NAME")

	if connName != "" {
		cfg := amqp.Config{
			Properties: amqp.Table{
				"connection_name": connName,
			},
		}

		c.conn, err = amqp.DialConfig(fmt.Sprintf("amqp://%s:%s@%s/", c.config.User, c.config.Pass, c.config.Address), cfg)
		if err != nil {
			return fmt.Errorf("error in creating rabbitmq connection with %s : %s", c.config.Address, err.Error())
		}
	} else {
		c.conn, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/", c.config.User, c.config.Pass, c.config.Address))
		if err != nil {
			return fmt.Errorf("error in creating rabbitmq connection with %s : %s", c.config.Address, err.Error())
		}
	}

	go func() {
		<-c.conn.NotifyClose(make(chan *amqp.Error)) //Listen to NotifyClose
		c.err <- errors.New("Connection Closed")
	}()
	c.channel, err = c.conn.Channel()
	if err != nil {
		return fmt.Errorf("channel: %s", err)
	}
	if err := c.channel.ExchangeDeclare(
		c.exchange.Name, // name
		c.exchange.Type, // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // noWait
		c.args,          // arguments
	); err != nil {
		return fmt.Errorf("error in Exchange Declare: %s", err)
	}
	return nil
}

// BindQueue do declare all needed queue, qos, and will bind the queue with selected BindKey to Exchange
func (c *Connection) BindQueue() error {
	for _, q := range c.queues {
		if _, err := c.channel.QueueDeclare(q.Name, true, false, false, false, nil); err != nil {
			return fmt.Errorf("error in declaring the queue %s", err)
		}
		if err := c.channel.Qos(1, 0, false); err != nil {
			return fmt.Errorf("error in set qos the queue %s", err)
		}
		if err := c.channel.QueueBind(q.Name, q.BindKey, c.exchange.Name, false, nil); err != nil {
			return fmt.Errorf("queue  Bind error: %s", err)
		}
	}
	return nil
}

// Reconnect reconnects the connection
func (c *Connection) Reconnect() error {
	if err := c.Connect(); err != nil {
		return err
	}
	if err := c.BindQueue(); err != nil {
		return err
	}
	return nil
}

// Consume consumes the messages from the queues and passes it as map of chan of amqp.Delivery
func (c *Connection) Consume() (map[string]<-chan amqp.Delivery, error) {
	m := make(map[string]<-chan amqp.Delivery)
	for _, q := range c.queues {
		deliveries, err := c.channel.Consume(q.Name, "", false, false, false, false, nil)
		if err != nil {
			return nil, err
		}
		m[q.Name] = deliveries
	}
	return m, nil
}

// HandleConsumedDeliveries handles the consumed deliveries from the queues. Should be called only for a consumer connection
func (c *Connection) HandleConsumedDeliveries(q string, delivery <-chan amqp.Delivery, fn func(Connection, string, <-chan amqp.Delivery)) {
	for {
		go fn(*c, q, delivery)
		if err := <-c.err; err != nil {
			c.Reconnect()
			deliveries, err := c.Consume()
			if err != nil {
				panic(err) //raising panic if consume fails even after reconnecting
			}
			delivery = deliveries[q]
		}
	}
}

// DelayProcess handle consumption with semaphore and delay in second.
// The purpose is to achieve exclusive periodic delay on each process of consumption with limitation related to semaphore size.
// fn as sub-function, is the main logic that can be inserted in this method to customize the outcome of delay process
func (c *Connection) DelayProcess(q string, deliveries <-chan amqp.Delivery, semaphoreSize int, delaySec int, slowDownMilisec int, keyBinding string, fn func(*Connection, Message)) {
	sem := semaphore.NewWeighted(int64(semaphoreSize))
	for d := range deliveries {
		m := Message{
			Queue:       q,
			Body:        MessageBody{Data: d.Body},
			ContentType: d.ContentType,
			Header:      d.Headers,
		}

		v := d

		// hanlde when semaphore doesn't have any room, avoid block state, slow down the loop then republish to broker
		if ok := sem.TryAcquire(1); !ok {
			log.Println("full semaphore detected")
			err := v.Ack(false)
			if err != nil {
				log.Println("error on ack #1: ", err)
			}
			time.Sleep(time.Duration(slowDownMilisec) * time.Millisecond)
			c.Publish(m, keyBinding)
			continue
		}

		go func() {
			err := v.Ack(false)
			if err != nil {
				fmt.Println("error on ack #0: ", err)
			}
			time.Sleep(time.Duration(delaySec) * time.Second)
			fn(c, m)
			sem.Release(1)
		}()
	}
}

// Publish publishes a request to the amqp queue
func (c *Connection) Publish(m Message, queueBindKey string) error {
	select { //non blocking channel - if there is no error will go to default where we do nothing
	case err := <-c.err:
		if err != nil {
			c.Reconnect()
		}
	default:
	}

	p := amqp.Publishing{
		DeliveryMode:  amqp.Persistent,
		Headers:       m.Header,
		ContentType:   m.ContentType,
		CorrelationId: m.CorrelationID,
		Body:          m.Body.Data,
		ReplyTo:       m.ReplyTo,
	}

	if err := c.channel.Publish(c.exchange.Name, queueBindKey, false, false, p); err != nil {
		return fmt.Errorf("error in Publishing: %s", err)
	}
	return nil
}
