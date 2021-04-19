package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancelSignal := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("Waiting for close signal")
			}

		}
	}()

	sig := <-sigs
	cancelSignal()
	fmt.Println("Received SIGTERM", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancel()
	<-ctx.Done()
	fmt.Println("Closed")
}
