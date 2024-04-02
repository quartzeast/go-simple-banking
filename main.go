package main

import (
	"github.com/quartzeast/go-simple-banking/app"
	"github.com/quartzeast/go-simple-banking/logger"
)

func main() {
	logger.Info("Application is starting")
	app.Start()
}
