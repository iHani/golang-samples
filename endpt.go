package main

import "fmt"

type EndPt struct {
	x string
}

func (self EndPt) IsConsumer() bool {
	return false
}

func (self EndPt) IsProducer() bool {
	return false
}

type Consumer struct {
	_name string
	a     string
	EndPt
}

func (self Consumer) IsConsumer() bool {
	return true
}

type Producer struct {
	_name string
	a     string
	EndPt
}

func (self Producer) IsProducer() bool {
	return true
}

/*func (e EndPt) dump() {
	fmt.Println("isConsumer =", e.IsConsumer())
	fmt.Println("isProducer =", e.IsProducer())
}*/

func (e Consumer) dump() {
	fmt.Println(e._name)
	fmt.Println(e.a)
	fmt.Println("isConsumer =", e.IsConsumer())
	fmt.Println("isProducer =", e.IsProducer())
	fmt.Println("")

}

func (e Producer) dump() {
	fmt.Println(e._name)
	fmt.Println(e.a)
	fmt.Println("isConsumer =", e.IsConsumer())
	fmt.Println("isProducer =", e.IsProducer())
	fmt.Println("")
}

func (e *Consumer) change(str string) {
	e.a = str
}
func (e *Producer) change(str string) {
	e.a = str
}

func main() {
	c1 := new(Consumer)
	c1._name = "c1"
	p1 := new(Producer)
	p1._name = "p1"
	c2 := new(Consumer)
	c2._name = "c2"
	p2 := new(Producer)
	p2._name = "p2"

	_ = c2
	_ = p2

	fmt.Println("isConsumer =", c1.IsConsumer())
	fmt.Println("isProducer =", c1.IsProducer())

	fmt.Println()

	fmt.Println("isProducer =", p1.IsProducer())
	fmt.Println("isConsumer =", p2.IsConsumer())

	fmt.Println("\nNow use dump()")
	fmt.Println("\nConsumers")
	c1.dump()
	c2.dump()
	fmt.Println("\nProducers")

	p1.dump()
	p2.dump()

	c1.change("consumer value")
	p1.change(("producer value"))
	c1.dump()
	p1.dump()

}
