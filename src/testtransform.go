package main

/** right now this i useless **/
import (
	"fmt"
	"pkgs"
)

func main() {
	var t = transform.Transform{"Transform"}
	//fmt.Println(t.Name)
	fmt.Println(t.Doit())

	var c = transform.Tchild{transform.Transform{"blah"}}
	//fmt.Println(c.Name)
	fmt.Println(c.Doit())
}
