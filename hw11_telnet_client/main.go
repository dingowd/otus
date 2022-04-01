package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "timeout duration")
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		log.Fatal("not right flags")
	}
	addr := net.JoinHostPort(args[0], args[1])
	client := NewTelnetClient(addr, *timeout, os.Stdin, os.Stdout)
	if err := client.Connect(); err != nil {
		log.Fatalf("unable to connect within %s", *timeout)
	}
	defer client.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	go func() {
		err := client.Receive()
		if err != nil {
			fmt.Fprintln(os.Stderr, "receive error: ", err)
		}
		cancel()
	}()
	go func() {
		err := client.Send()
		if err != nil {
			fmt.Fprintln(os.Stderr, "send error: ", err)
		}
		cancel()
	}()

	<-ctx.Done()
}
