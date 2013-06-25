package main

import "fmt"

type EndPt struct {
}

func (self EndPt) IsConsumer() bool {
	return false
}

func (self EndPt) IsProducer() bool {
	return false
}

type Consumer struct {
	EndPt
}

func (self Consumer) IsConsumer() bool {
	return true
}

type Producer struct {
	EndPt
}

func (self Producer) IsProducer() bool {
	return true
}

func main() {
	consumer := new(Consumer)
	producer := new(Producer)

	fmt.Println("isConsumer =", consumer.IsConsumer())
	fmt.Println("isProducer =", consumer.IsProducer())

	fmt.Println()

	fmt.Println("isProducer =", producer.IsProducer())
	fmt.Println("isConsumer =", producer.IsConsumer())
}
