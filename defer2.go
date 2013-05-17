package main

/**
 * when the loop in main ends the defered func shutdown() is called
 *
 * this works as expected
 **/

import (
	"fmt"
	"os"
)

func main() {
	defer shutdown()
	for i := 0; i < 1000; i++ {
		fmt.Print(".")
		if i%10 == 0 {
			fmt.Print(i)
		}
	}
}

func shutdown() {
	fmt.Println("\nIn shutdown()")
	os.Exit(1)
}
