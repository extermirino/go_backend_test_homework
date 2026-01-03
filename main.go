package main

import (
	"fmt"
)

func Add(a, b int) int {
	return a + b
}

func hello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func main() {
	fmt.Println("Я домашка")
	hello("Nikola")
}
