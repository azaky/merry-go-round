package main

import (
	"fmt"

	a "github.com/azaky/merry-go-round/a"
	b "github.com/azaky/merry-go-round/b"
)

func main() {
	serviceA := new(a.Service)
	serviceB := new(b.Service)
	serviceA.ServiceB = serviceB
	serviceB.ServiceA = serviceA

	fmt.Printf("a: %s\n", serviceA.GetSomething())
	fmt.Printf("b: %s\n", serviceB.GetName())
}
