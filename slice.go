package main
import "fmt"
func main() {
	var s string
	s = "1234567890"
	fmt.Println(s)
	y := s[2:3]
	fmt.Println(y)
	fmt.Println(s[2:])
	fmt.Println("\nMore advanced slices")
	var b = []string{"A", "B", "C"}
	var c = []string{"1", "2", "3"}
	
	z := make([][]string, 2, 10)

	fmt.Println(b)
	fmt.Println(len(z), cap(z))
	z[0] = b
	fmt.Println(z[:1])
	b[1] = "x"
	fmt.Println(z[:1])
	z[1] = c
	fmt.Println(z[:2])
	w := make([][]string, 2, 3)
	j := []int{0,1}
	for i := range j {
		fmt.Println(i)
		w[i] = z[i]
	}
	fmt.Println(w[:2])
	b[1] = "P"
	fmt.Println(w[:2])
	fmt.Println(len(w), cap(w))
	w = append(w, []string{"4", "5", "6"})
	w = append(w, []string{"4", "5", "6"})
	w = append(w, []string{"4", "5", "6"})
	fmt.Println(w)
	fmt.Println(len(w), cap(w))
	
}