package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// not that unmarshal seems to choke on \n in the json string
func main() {
	fmt.Println("we will unmarshal json without knowing the data structure")

	var f interface{}
	err := json.Unmarshal(jsonByte(), &f)
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
	prtKeys(f, 0)
}

func jsonByte() []byte {

	return []byte(`{"parent":{"name":"a parent name","child1":{"type":"child","name":"child1","desc":"some desc of child 1"},"child2":{"type":"child","name":"child3","desc":"child2 has children","gchild1":{"prop1":"value1-a","prop2":"value2-a"},"gchild2":{"prop1":"value1-b","prop2":"value2-b"}},"child3":{"type":"child","name":"child3","desc":"some desc of child 3"}}}`)

}
func prtKeys(f interface{}, level int) {

	myvalue, ok := f.(string)
	if ok {
		fmt.Println("value: ", myvalue)
	} else {
		mymap, ok := f.(map[string]interface{})
		if !ok {
			//fmt.Println("mymap, ok := f.(map[string]interface{}) NOT OK")
			level = level - 1
			return
		} else {
			//fmt.Println(mymap["parent"])
		}

		for k, n := range mymap {
			_ = n
			fmt.Println(level)
			if level-1 < 0 {
				fmt.Println(k)
			} else {
				fmt.Println(strings.Repeat(" ", level-1) + k)

			}
			level = level + 1
			prtKeys(n, level)

		}
	}

}
