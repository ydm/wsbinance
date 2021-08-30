package main

import (
	"context"
	"os"

	"github.com/adshao/go-binance/v2"
	"github.com/joho/godotenv"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	GoInterrupt(ctx, cancel)

	_ = godotenv.Load()

	symbol := os.Getenv("SYMBOL")
	if symbol == "" {
		symbol = "BTCUSDT"
	}

	events := make(chan *binance.WsBookTickerEvent, 128)
	defer close(events)

	writer := Writer{
		File:     os.Stdout,
		Input:    events,
		Sequence: 0,
	}
	go writer.Loop()

	done, stop, err := binance.WsBookTickerServe(
		symbol,
		func(x *binance.WsBookTickerEvent) {
			events <- x
		},
		func(err error) { panic(err) },
	)
	if err != nil {
		panic(err)
	}

	go func() {
		<-ctx.Done()
		close(stop)
	}()

	<-done
}
