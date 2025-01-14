package logicol

import (
	"fmt"
	"log"
	"runtime"
)

type Level string

const (
	Info    Level = "Info"
	Debug         = "Debug"
	Warning       = "Warning"
	Error         = "Error"
	Event         = "Event"
)

type Color string

const (
	Reset  Color = "\033[0m"
	Red          = "\033[31m"
	Green        = "\033[32m"
	Yellow       = "\033[33m"
	Blue         = "\033[34m"
	Purple       = "\033[35m"
	Cyan         = "\033[36m"
	Grey         = "\033[37m"
	White        = "\033[97m"
)

// Write writes a new log to the default logger set by the native log package. The output has added color highlighting and the name of the function that called the Write function as well as the line number of the function call.
func Write(logLevel Level, message interface{}) {
	var caller string
	if logLevel == Event {
		caller = ""
	} else {
		caller = getCaller()
	}

	color := getColor(&logLevel)

	// output is adjusted for different log levels
	switch logLevel {
	case Info, Debug:
		// reset the color wrapping BEFROE the message content
		log.Printf("%s[%s] %s%s %s", color, logLevel, caller, Reset, message)
	case Warning, Error:
		// reset the color wrapping AFTER the message content
		log.Printf("%s[%s] %s %s%s", color, logLevel, caller, message, Reset)
	case Event:
		log.Printf("%s[%s]%s %s", color, logLevel, Reset, message)
	default:
		log.Println(message)
	}
}

// getCaller retrieves the function that called the Write function.
func getCaller() string {
	pc, _, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	name := runtime.FuncForPC(pc).Name()
	return fmt.Sprintf("%s : %d ", name, line)
}

// getColor determines what color to wrap log content with
func getColor(l *Level) Color {
	switch *l {
	case Info:
		return Blue
	case Debug:
		return Cyan
	case Warning:
		return Yellow
	case Error:
		return Red
	case Event:
		return Purple
	default:
		return White
	}
}
