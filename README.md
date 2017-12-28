# minilog
golang minimalistic package for logging

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