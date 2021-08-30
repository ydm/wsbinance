package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/joho/godotenv"
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

// func newClient() {
// 	apikey := os.Getenv("BINANCE_APIKEY")
// 	secret := os.Getenv("BINANCE_SECRET")
// 	client := binance.NewClient(apikey, secret)
// 	return client
// }

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	GoInterrupt(ctx, cancel)

	_ = godotenv.Load()

	symbol := os.Getenv("SYMBOL")
	if symbol == "" {
		symbol = "BTCUSDT"
	}

	fmt.Println("ToB,msg_seq_num,md_seq_num,ts_transact,bid_price,ask_price,bid_volume,ask_volume")
	seq := 0
	f := func(x *binance.WsBookTickerEvent) {
		timestamp := time.Now().UnixNano() / 1000000
		fmt.Printf(
			"ToB,%d,%d,%d,%s,%s,%s,%s\n",
			x.UpdateID,
			seq,
			timestamp,
			x.BestBidPrice,
			x.BestAskPrice,
			x.BestBidQty,
			x.BestAskQty,
		)
		seq++
	}

	g := func(err error) {
		panic(err)
	}

	done, stop, err := binance.WsBookTickerServe(symbol, f, g)
	if err != nil {
		panic(err)
	}

	go func() {
		<-ctx.Done()
		close(stop)
	}()

	<-done
}
