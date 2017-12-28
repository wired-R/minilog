# minilog
Golang minimalistic package for logging

Color output logs

The possibility of logging to a file

```
go get github.com/wired-R/minilog

```

Example

```go
package main

import log "github.com/wired-R/minilog"

func main() {
    //set debug log level
    log.Level = 0
    log.Debug("This is debug message")
}
