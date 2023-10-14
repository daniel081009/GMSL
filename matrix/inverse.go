package matrix

import "fmt"

/*
Matrix determinant
행렬식
*/
func (m Matrix) Det() int {
	if len(m) != m.J() {
		return 0
	}
	if len(m) == 1 {
		return m[0][0]
	}

	var res int
	for i := range m {
		res += m[i][0] * m.Cut(i, 0).Det() * square(-1, i+0)
	}
	return res
}

/*
adjugate matrix
여인수
*/
func (m Matrix) Adj() Matrix {
	tmp := Matrix{}
	tmp.Create(len(m), len(m))
	for i := range m {
		for j := range m[i] {
			tmp[i][j] = m.Cut(i, j).Det() * square(-1, i+j)
		}
	}
	return tmp
}

/*
Matrix inverse
역행렬
*/
func (m Matrix) Inverse() (Matrix, error) {
	if !m.CheckSquare() {
		return nil, fmt.Errorf("not square matrix")
	}
	if m.Det() == 0 {
		return nil, fmt.Errorf("determinant is zero")
	}
	tmp := Matrix{}
	tmp.Create(m.I(), m.J())

	return m.Adj().DivInt(m.Det()).Transpose(), nil
}

func square(d, i int) int {
	if i == 0 {
		return 1
	}
	return d * square(d, i-1)
}
