package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func jiggling(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("exit")
			return
		case <-ticker.C:
			robotgo.MoveRelative(10, 10)
			time.Sleep(time.Millisecond * 200)
			robotgo.MoveRelative(-10, -10)
		}
	}

}
