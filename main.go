package main

import (
	"fmt"
	"gmsl/Matrix"
)

func main() {
	m := Matrix.Matrix{}
	m.Create(2, 2).Write([][]int{{1, 2}, {3, 4}})
	m2 := Matrix.Matrix{}
	m2.Create(2, 2).Write([][]int{{1, 2}, {3, 4}})

	fmt.Println("m + m2")
	m.Add(m2).Print()
	fmt.Println("\nm - m2")
	m.Sub(m2).Print()
	fmt.Println("\nm * m2")
	m.Mul(m2).Print()
	fmt.Println("\nm * 2")
	m.SMul(2).Print()
	fmt.Println("\nm transpose")
	m.Transpose().Print()
}
