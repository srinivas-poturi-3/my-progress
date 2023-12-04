package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"templates/internal/logs"
	"templates/internal/pkg/app"
)

var (
	log = &logs.Logger{}
)

func main() {
	log = logs.NewLogger(logs.Debug)
	wg := sync.WaitGroup{}
	a := app.Make(&wg)
	if err := a.Start(); err != nil {
		log.Error("error starting application")
	}

	log.Infof("started server : %s\n", a.Name())
	go forceStop(a)

	// Handle unexpected panics
	defer func() {
		if r := recover(); r != nil {
			log.Warnf("Recovered from panic:", r)
		}
	}()

	wg.Wait()
	log.Info("Stopped server")
}

func forceStop(worker app.App) {
	// Handle OS termination signals
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		log.Warn("Received termination signal. Shutting down gracefully...")
		worker.Stop()
	}()
}
