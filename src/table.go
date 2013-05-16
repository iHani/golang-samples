package main

//import "pkgs"
import "fmt"
func main() {
	var r1 = []string{"r1A", "r1B", "r1B"}
	var r2 = []string{"r2A", "r2B", "r2B"}
	var data = [][]string{r1, r2}
	fmt.Println("In Main")
	fmt.Println(data)
	tbl := new(TextTable)
	tbl.Init(data)
	tbl.Dump()
	tbl.AddRow([]string{"r3A", "r3B", "r3B"})
	tbl.Dump()
}

/** for now forget the import pkgs and just do this locally **/

type TextTable struct {
	Data *[][]string
}

func (t *TextTable) Init(data [][]string) {
	fmt.Println("Inside Init")
	fmt.Println(data)
	t.Data = &data
}
func (t *TextTable) AddRow(row []string) {
	fmt.Println("Inside Add Row")
	fmt.Println(row)
	t.Data = append(t.(&Data), row)
}

func (t TextTable) Dump() {
	fmt.Println("Inside Dump")
	fmt.Println(t.Data)
}
