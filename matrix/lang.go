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

/*
delete i and j
특정 행과 열을 제거
*/
func (m Matrix) Cut(i, k int) Matrix {
	tmp := Matrix{}
	tmp.Create(m.I()-1, m.J()-1)
	xx, yy := 0, 0
	for x := range m {
		if x == i {
			continue
		}
		for y := range m[x] {
			if y == k {
				continue
			}
			tmp[xx][yy] = m[x][y]
			yy++
		}
		yy = 0
		xx++
	}
	return tmp
}
