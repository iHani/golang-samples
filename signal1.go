package main

/**
 * to test this build and fun the exe file, don't use go run
 *
 * when ctl-c is detect this will cal shutdown
 *
 * this works as expected with
 * ctrl-c
 *THIS DOES NOT TRAP os.kill
 **/

import (
	"fmt"
	"os"
	"os/signal"
	//"time"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(2)
	go signalListener()

	// for empty for loop (for {}) we need runtime.GOMAXPROCS(2) so the main func yields
	// or do some sleep thing or I/O
	for {
		//time.Sleep(10000)
	}
}

func signalListener() {
	fmt.Println("in signalListener()")
	interrupt := make(chan os.Signal, 100)
	signal.Notify(interrupt, os.Interrupt)
	defer shutdown()
	// infinite loop
	for {
		s := <-interrupt
		_ = s
		fmt.Print("interrupt detected: ")
		panic(" a panic")
	}
}

func shutdown() {
	fmt.Println("In shutdown()")
	os.Exit(1)
}
