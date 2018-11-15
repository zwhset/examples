package other

import (
	"fmt"
	"qianbao.com/examples/logs/logger"
	"time"
)

func Printlog() {
	i := 0
	for {
		logger.Logger.Info(fmt.Sprintf("PrintLog goroutine: %d", i))
		i++
		time.Sleep(time.Second * 2)
	}
}
