package kekasigohelper

import (
	"fmt"
	"log"
)

const (
	colorRed    string = "\033[31m"
	colorGreen  string = "\033[32m"
	colorYellow string = "\033[33m"
	colorBlue   string = "\033[34m"
	colorPurple string = "\033[35m"
	colorCyan   string = "\033[36m"
	colorWhite  string = "\033[37m"
)

//this is for logger info level
func LoggerInfo(message interface{}) {
	log.Println(colorCyan, "[ INFO ] => "+fmt.Sprint(message)+".")
}

//this is for logger warning level
func LoggerWarning(message interface{}) {
	log.Println(colorYellow, "[ WARNING ] => "+fmt.Sprint(message)+".")
}

//this is for logger success level
func LoggerSuccess(message interface{}) {
	log.Println(colorGreen, "[ SUCCESS ] => "+fmt.Sprint(message)+".")
}

//this is for logger error level
func LoggerError(err error) {
	if err != nil {
		log.Println(colorRed, "[ ERROR ] => "+err.Error()+".")
	}
}

//this is for logger debug level
func LoggerDebug(msg interface{}) {
	log.Println(colorPurple, "[ DEBUG ] => "+fmt.Sprint(msg)+".")
}

//this is for logger fatal level
func LoggerFatal(msg interface{}) {
	log.Fatal(colorRed, "[ FATAL ] => "+fmt.Sprint(msg)+".")
}
