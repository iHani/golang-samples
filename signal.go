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
)

func main() {
	// set buffer up to 100 just in case
	interrupt := make(chan os.Signal, 100)
	signal.Notify(interrupt, os.Interrupt, os.Kill)
	defer shutdown()
	// infinite loop
	for {
		s := <-interrupt
		fmt.Println("interrupt detected: ", s)
		panic(" a panic")
	}
}

func shutdown() {
	fmt.Println("\nIn shutdown()")
	os.Exit(1)
}
