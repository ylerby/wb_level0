package main

import (
	"os"
	"os/signal"
	"subscriber/internal/app"
	"syscall"
)

func main() {
	application := app.New()
	go application.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.Stop()
}
