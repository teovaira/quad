package main

import (
	"fmt"

	"quad"
)

func main() {
	quad.QuadA(5, 3)
	quad.QuadA(5, 1)
	quad.QuadA(1, 1)
	quad.QuadA(1, 5)

	fmt.Println()

	quad.QuadB(5, 3)
	quad.QuadB(5, 1)
	quad.QuadB(1, 1)
	quad.QuadB(1, 5)

	fmt.Println()

	quad.QuadC(5, 3)
	quad.QuadC(5, 1)
	quad.QuadC(1, 1)
	quad.QuadC(1, 5)

	fmt.Println()

	quad.QuadD(5, 3)
	quad.QuadD(5, 1)
	quad.QuadD(1, 1)
	quad.QuadD(1, 5)

	fmt.Println()

	quad.QuadE(5, 3)
	quad.QuadE(5, 1)
	quad.QuadE(1, 1)
	quad.QuadE(1, 5)
}
