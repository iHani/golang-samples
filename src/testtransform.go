package main
import (
    "fmt"
    "pkgs" 
    )

func main() {
	var t = transform.Transform{"Transform"}
	//fmt.Println(t.Name)
	fmt.Println(t.Doit())

	var c = transform.Tchild{Transform{"blah"}}
	//fmt.Println(c.Name)
	fmt.Println(c.Doit())
}