package main

/**
 * okay the panic causes the defered routine to be called
 *
 * this works as expected
 **/

import (
	"fmt"
	"os"
)

func main() {
	defer shutdown()
	i := 0
	for {
		i++

		if i > 100000 {
			panic("a panic")
		}
	}
}

func shutdown() {
	fmt.Println("In shutdown()")
	os.Exit(1)
}
