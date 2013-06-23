package main

import (
	"fmt"
	"reflect"
)

type Ctest struct {
	NameStrDef string
	Blah       int
	Flag       bool
}

func main() {
	c := Ctest{"avalue", 123, true}
	s := reflect.ValueOf(&c).Elem()
	typeOfC := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfC.Field(i).Name, f.Type(), f.Interface())
		if typeOfC.Field(i).Name == "Blah" {
			fmt.Println("setting the balue of: " + typeOfC.Field(i).Name)
			fmt.Println("settability of s:", f.CanSet())
			f.SetInt(200)
		}
	}
	fmt.Println(c)
}
