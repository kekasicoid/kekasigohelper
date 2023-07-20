package kekasigohelper

import (
	"fmt"
	"strconv"
	"strings"
)

func Calculator(expression string) (float64, error) {
	tokens := Tokenize(expression)
	// Reverse Polish Notation (RPN) algorithm
	// Using two stacks: one for operators and another for operands
	operators := []string{}
	operands := []float64{}

	// Operator precedence map
	precedence := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}

	for _, token := range tokens {
		if IsNumber(token) {
			operand, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, err
			}
			operands = append(operands, operand)
		} else if token == "(" {
			operators = append(operators, token)
		} else if token == ")" {
			for len(operators) > 0 && operators[len(operators)-1] != "(" {
				if err := Evaluate(&operators, &operands); err != nil {
					return 0, err
				}
			}
			// Remove the "(" from the operators stack
			if len(operators) > 0 && operators[len(operators)-1] == "(" {
				operators = operators[:len(operators)-1]
			}
		} else {
			for len(operators) > 0 && precedence[operators[len(operators)-1]] >= precedence[token] {
				if err := Evaluate(&operators, &operands); err != nil {
					return 0, err
				}
			}
			operators = append(operators, token)
		}
	}

	for len(operators) > 0 {

		if err := Evaluate(&operators, &operands); err != nil {
			return 0, err
		}
	}
	if len(operands) != 1 {
		return 0, fmt.Errorf("invalid expression")
	}

	return operands[0], nil
}

func Tokenize(expression string) []string {
	// Remove whitespaces and split into tokens
	expression = strings.ReplaceAll(expression, " ", "")
	tokens := []string{}
	currentToken := ""

	for _, char := range expression {
		switch char {
		case '+', '-', '*', '/', '(', ')':
			if currentToken != "" {
				tokens = append(tokens, currentToken)
				currentToken = ""
			}
			tokens = append(tokens, string(char))
		default:
			currentToken += string(char)
		}
	}

	if currentToken != "" {
		tokens = append(tokens, currentToken)
	}

	return tokens
}

func IsNumber(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

func Evaluate(operators *[]string, operands *[]float64) error {
	if len(*operands) < 2 {
		return fmt.Errorf("invalid expression")
	}

	operator := (*operators)[len(*operators)-1]
	*operators = (*operators)[:len(*operators)-1]

	operand2 := (*operands)[len(*operands)-1]
	*operands = (*operands)[:len(*operands)-1]

	operand1 := (*operands)[len(*operands)-1]
	*operands = (*operands)[:len(*operands)-1]

	result, err := PerformOperation(operand1, operand2, operator)
	if err != nil {
		return err
	}

	*operands = append(*operands, result)
	return nil
}

func PerformOperation(operand1, operand2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return operand1 + operand2, nil
	case "-":
		return operand1 - operand2, nil
	case "*":
		return operand1 * operand2, nil
	case "/":
		if operand2 == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return operand1 / operand2, nil
	default:
		return 0, fmt.Errorf("invalid operator: %s", operator)
	}
}
