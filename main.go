package main

import (
	"context"
	"os"

	"github.com/adshao/go-binance/v2/futures"
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

	events := make(chan *futures.WsBookTickerEvent, 128)
	defer close(events)

	output := os.Getenv("OUTPUT")
	if output == "" {
		output = "/tmp"
	}

	writer := Writer{
		File:     makeWriter(output),
		Input:    events,
		Sequence: 0,
	}
	go writer.Loop()

	done, stop, err := futures.WsBookTickerServe(
		symbol,
		func(x *futures.WsBookTickerEvent) {
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
