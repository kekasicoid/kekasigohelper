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
	fmt.Println(colorCyan)
	log.Println(colorCyan, "[ INFO ] => "+fmt.Sprint(message)+".")
	log.Println(colorCyan, "========== End Of Info Message ==========")
	fmt.Println(colorCyan)
}

//this is for logger warning level
func LoggerWarning(message interface{}) {
	fmt.Println(colorYellow)
	log.Println(colorYellow, "[ WARNING ] => "+fmt.Sprint(message)+".")
	fmt.Println(colorYellow)
}

//this is for logger success level
func LoggerSuccess(message interface{}) {
	fmt.Println(colorGreen)
	log.Println(colorGreen, "[ SUCCESS ] => "+fmt.Sprint(message)+".")
	fmt.Println(colorGreen)
}

//this is for logger error level
func LoggerError(err error) {
	if err != nil {
		fmt.Println(colorRed)
		log.Println(colorRed, "[ SUCCESS ] => "+err.Error()+".")
		fmt.Println(colorRed)
	}
}

//this is for logger debug level
func LoggerDebug(msg interface{}) {
	fmt.Println(colorPurple)
	log.Println(colorPurple, "[ SUCCESS ] => "+fmt.Sprint(msg)+".")
	fmt.Println(colorPurple)
}
