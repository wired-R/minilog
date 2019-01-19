# minilog
Golang minimalistic package for logging

The possibility color output logs

The possibility write log to file

```
go get github.com/wired-R/minilog

```

Example

```go
package main

import (
	"os"

	log "github.com/wired-R/minilog"
)

func main() {
	var err error

	log.Level = 0                         // default 2 - warning
	log.TimeFormat = "02/01 15:04:05.000" // default
	log.Colored = true                    // default

	// Write log to file
	log.Output, err = os.OpenFile("file.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.Debug("Debug message")
	log.Info("Info message")
	log.Warning("warning message")
	log.Error("Error message")
	log.Fatal("Fatal message")

}

```
![](images/view.png)
