package main

import (
	"io"
	"path"

	"gopkg.in/natefinch/lumberjack.v2"
)

func makeWriter(dir string) io.Writer {
	// 1MB == ~11200 lines, ~2:15min
	return &lumberjack.Logger{
		Filename:   path.Join(dir, "bookTicker.csv"),
		MaxSize:    50,
		MaxAge:     0, // Do not remove old files based on age.
		MaxBackups: 0, // Retain all old files.
		LocalTime:  false,
		Compress:   false,
	}
}
