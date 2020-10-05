package signal

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)
var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}
var onlyOneSignalHandler = make(chan struct{})
var shutdownHandler chan os.Signal

func SetupSignalContext() context.Context {
	close(onlyOneSignalHandler) // panics when called twice

	shutdownHandler = make(chan os.Signal, 2)

	ctx, cancel := context.WithCancel(context.Background())
	signal.Notify(shutdownHandler, shutdownSignals...)
	go func() {
		<-shutdownHandler
		cancel()
		<-shutdownHandler
		os.Exit(1) // second signal. Exit directly.
	}()

	return ctx
}
