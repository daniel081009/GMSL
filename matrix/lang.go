package matrix

import "fmt"

func (m *Matrix) Create(i, j int) *Matrix {
	for k := 0; k < i; k++ {
		*m = append(*m, make([]int, j))
		for l := 0; l < j; l++ {
			(*m)[k][l] = 0
		}
	}
	return m
}

func (m *Matrix) Write(d [][]int) {
	if len(*m) != len(d) || len((*m)[0]) != len(d[0]) {
		panic("Matrix size is not equal to data size")
	} else if len(*m) == 0 || len((*m)[0]) == 0 {
		panic("Matrix is empty")
	}
	for i := range *m {
		for j := range (*m)[i] {
			(*m)[i][j] = d[i][j]
		}
	}
}

/*
print Matrix
행열을 출력함
*/
func (m Matrix) Print() {
	for i := range m {
		for j := range (m)[i] {
			fmt.Printf("%d ", (m)[i][j])
		}
		fmt.Println()
	}
}
