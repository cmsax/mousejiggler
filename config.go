package main

import (
	"context"
	"flag"
	"time"
)

func loadConfig() (context.Context, context.CancelFunc, time.Duration) {
	interval := flag.Int("interval", 1, "jiggling interval seconds")
	dur := flag.Int("duration", -1, "jiggling duration seconds")
	flag.Parse()

	if *interval <= 0 {
		*interval = 1
	}
	itv := time.Second * time.Duration(*interval)

	if *dur <= 0 {
		ctx, cancel := context.WithCancel(context.Background())
		return ctx, cancel, itv
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(*dur))
	return ctx, cancel, itv
}
