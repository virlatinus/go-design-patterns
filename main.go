package main

import (
	"go-design-patterns/patterns"
)

func main() {
	// SOLID principles
	srp := new(patterns.SingleResponsibility)
	srp.Run()

	ocp := new(patterns.OpenClose)
	ocp.Run()
}
