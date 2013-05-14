package main

import "strings"
import "fmt"

func main() {
	var a,b,c string =  "A", "B", "C"
	fmt.Println(a,b,c)


	myStrArray := []string{a, b, c}
	var myStr = strings.Join(myStrArray, "\t")
	fmt.Println(myStr)
	fmt.Println(strings.Join([]string{
		"X", "Y"},":"))
	