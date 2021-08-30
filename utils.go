package main

import (
	"context"
	"os"
	"os/signal"
)

func GoInterrupt(ctx context.Context, cancel context.CancelFunc) {
	go func() {
		if Interrupt(ctx) {
			cancel()
		}
	}()
}

// Interrupt returns when either (1) interrupt signal is received by
// the OS or (2) the given context is done.
func Interrupt(ctx context.Context) bool {
	appSignal := make(chan os.Signal, 1)
	signal.Notify(appSignal, os.Interrupt)
	select {
	case <-appSignal:
		return true
	case <-ctx.Done():
		return false
	}
}
