package main

import (
	"fmt"
	"io"
	"time"

	"github.com/adshao/go-binance/v2"
)

type Writer struct {
	File     io.Writer
	Input    chan *binance.WsBookTickerEvent
	Sequence int
}

func (w *Writer) Header() {
	fmt.Fprintln(w.File, "ToB,msg_seq_num,md_seq_num,ts_transact,bid_price,ask_price,bid_volume,ask_volume")
}

func (w *Writer) Row(x *binance.WsBookTickerEvent) {
	timestamp := time.Now().UnixNano()
	fmt.Fprintf(
		w.File,
		"ToB,%d,%d,%d,%s,%s,%s,%s\n",
		x.UpdateID,
		w.Sequence,
		timestamp,
		x.BestBidPrice,
		x.BestAskPrice,
		x.BestBidQty,
		x.BestAskQty,
	)
	w.Sequence++
}

func (w *Writer) Loop() {
	for x := range w.Input {
		w.Row(x)
	}
}
