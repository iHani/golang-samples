package main

import (
	"fmt"
)

var a int64 = 0

func main() {

	for i := 0; i < 3; i++ {
		incA()
		fmt.Println("Outside a:", a)
	}

	var b int64 = 0
	for i := 0; i < 3; i++ {
		incB(&b)
		fmt.Println("Outside b:", b)
	}

	var c int64 = 0
	for i := 0; i < 3; i++ {
		c = incC(c)
		fmt.Println("Outside c:", c)
	}
}

func incA() {
	/* when a is declared outside of main() it is visible by ref inside of function
	 * so you can get static like behavior with this variable
	 */

	fmt.Println("Inside a:", a)
	a = a + 1
}

func incB(b *int64) {
	/* when b is passed by ref it can be incremented in the outter scope
	 * so you can get static like behavior with this variable
	 */

	fmt.Println("Inside b:", *b)
	*b = *b + 1
}

func incC(c int64) int64 {
	/* of course you can pass it in and return it out
	 */

	fmt.Println("Inside c:", c)
	return c + 1
}
