package minilog

import (
	"fmt"
	"os"
	"time"
)

// see https://en.wikipedia.org/wiki/ANSI_escape_code
var colors = map[string]string{
	"FATAL": "\033[0;31m%s\033[0m", // RED
	"ERROR": "\033[1;31m%s\033[0m", // ORANGE
	"WARN":  "\033[0;33m%s\033[0m", // YELOW
	"INFO":  "\033[0;32m%s\033[0m", // GREEN
	"DEBUG": "\033[0;34m%s\033[0m", // BLUE
}

const (
	debug = iota
	info
	warn
	err
	fatal
)

var (
	// Verbose log level
	// 0 - debug log
	// 4 - Fatal
	Level = warn

	// Output method
	Output = os.Stdout

	//Colored specifies whether to colorize logs
	Colored = true

	// TimeFormat https://golang.org/src/time/format.go
	// now is DD/MM hh:mm:ss.mil
	TimeFormat = "02/01 15:04:05.000"
)

func print(message interface{}, level string) {
	//format datetime [level] mesage
	format := "%s [%s] %s\n"

	var levelFormat string
	if Colored {
		levelFormat = fmt.Sprintf(colors[level], level)
	} else {
		levelFormat = level
	}

	fmt.Fprintf(Output, format, time.Now().Format(TimeFormat), levelFormat, message)
}

func Debug(message interface{}) {
	if Level <= debug {
		print(message, "DEBUG")
	}
}

func Info(message interface{}) {
	if Level <= info {
		print(message, "INFO")
	}
}

func Warning(message interface{}) {
	if Level <= warn {
		print(message, "WARN")
	}
}

func Error(message interface{}) {
	if Level <= err {
		print(message, "ERROR")
	}
}

func Fatal(message interface{}) {
	if Level <= fatal {
		print(message, "FATAL")
	}
	os.Exit(1)
}
