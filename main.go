package main

import (
	"fmt"
	"gmsl/matrix"
)

func main() {
	m := matrix.Matrix{}
	m.Create(2, 2).Write([][]int{{1, 2}, {3, 4}})
	m2 := matrix.Matrix{}
	m2.Create(2, 2).Write([][]int{{1, 3}, {3, 31}})
	m3 := matrix.Matrix{}
	m3.Create(3, 3).Write([][]int{{1, 1, 1}, {0, 1, 1}, {2, 2, 1}})

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

	fmt.Println("\nm3 cut")
	m3.Cut(0, 0).Print()

	fmt.Println("\nm determinant")
	fmt.Println(m3.Det())

	fmt.Println("\nm3 inverse")
	d, e := m3.Inverse()
	if e != nil {
		panic(e)
	}
	d.Print()
}
