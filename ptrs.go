package main

import "fmt"

func main() {
	a := [2]string{"a", "b"}
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(a[0])
	fmt.Println(&a[0])
	fmt.Println(a[1])
	fmt.Println(&a[1])
	ptr := 	&a[1]
	fmt.Println(*ptr)
	b := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(b)
	fmt.Println(&b)
	c := b[1:2]
	fmt.Println(c)
	fmt.Println(&c)

}