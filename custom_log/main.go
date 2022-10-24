package main

import (
	"os"

	"github.com/WeixinX/examples/custom_log/logger"
)

func main() {
	file := "./test.log"
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	l := logger.New(logger.INFO, f)
	l.Errorln("%s, hello world!", "This is a Error log")
	l.Warnln("%s, hello world!", "This is a Warn log")
	l.Infoln("%s, hello world!", "This is a Info log")
	l.Debugln("%s, hello world!", "This is a Debug log") // 因为等级设置为 INFO，所以 Debug 日志忽略
}
