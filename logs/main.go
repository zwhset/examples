package main

import (
	"fmt"
	"qianbao.com/examples/logs/logger"
	"qianbao.com/examples/logs/other"
	"time"
)

func main() {
	go tLog()        // goroutine 1
	go tLog()        // goroutine 2
	other.Printlog() // main goroutine
}

func tLog() {
	i := 0
	for {
		logger.Logger.Error(fmt.Sprintf("Info: main: %d", i))
		i++
		time.Sleep(time.Second)
		if i > 100 {
			break
		}
	}
}
