package main

import "fmt"

func main() {
	fmt.Println("hello world")
	PrintAnotherThing("Sasha")
	fmt.Println(Multiply(11))
}

func PrintAnotherThing(name string) {
	fmt.Println(name + " is the goodest girl")
}

func Multiply(val int) int {
	return val * val
}
