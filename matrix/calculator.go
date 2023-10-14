package matrix

import "fmt"

type Matrix [][]int

/*
addition
덧셈
*/
func (m Matrix) Add(md Matrix) Matrix {
	if !m.CheckEqualSize(md) {
		return nil
	}
	tmp := Matrix{}
	tmp.Create(len(m), m.J()).Write(m)
	for i := range tmp {
		for j := range tmp[i] {
			tmp[i][j] += md[i][j]
		}
	}
	return tmp
}

/*
subtraction Matrix
행렬의 뺄셈
*/
func (m Matrix) Sub(md Matrix) Matrix {
	if !m.CheckEqualSize(md) {
		return nil
	}

	tmp := Matrix{}
	tmp.Create(len(m), m.J()).Write(m)
	for i := range tmp {
		for j := range tmp[i] {
			tmp[i][j] -= md[i][j]
		}
	}
	return tmp
}

/*
transpose
전치행렬
*/
func (m Matrix) Transpose() Matrix {
	tmp := Matrix{}
	tmp.Create(m.J(), len(m))
	for i := range m {
		for j := range m[i] {
			tmp[j][i] = m[i][j]
		}
	}
	return tmp
}

/*
scalar multiplication
스칼라곱
*/
func (m Matrix) SMul(s int) Matrix {
	tmp := Matrix{}
	tmp.Create(len(m), m.J()).Write(m)
	for i := range tmp {
		for j := range tmp[i] {
			tmp[i][j] *= s
		}
	}
	return tmp
}

/*
Matrix multiplication
행렬의 곱셈
*/
func (m Matrix) Mul(md Matrix) Matrix {
	if m.J() != m.I() {
		return nil
	}
	res := Matrix{}
	res.Create(len(m), len(md[0]))
	for i := range m {
		for j := range md[0] {
			for k := range md {
				res[i][j] += m[i][k] * md[k][j]
			}
		}
	}
	return res
}

/*
matrix division
행렬대 행렬 나누기
*/
func (m Matrix) Div(md Matrix) Matrix {
	if !m.CheckEqualSize(md) {
		return nil
	}
	d, e := md.Inverse()
	if e != nil {
		panic(e)
	}
	return m.Mul(d)
}

/*
matrix division int
행렬대 int 나누기
*/
func (m Matrix) DivInt(s int) Matrix {
	for i := range m {
		for j := range m[i] {
			m[i][j] /= s
		}
	}
	return m
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
			// fmt.Println(xx, yy, x, y, m[x][y])
			tmp[xx][yy] = m[x][y]
			yy++
		}
		yy = 0
		xx++
	}
	return tmp
}

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

	return m.Adj().DivInt(m.Det()), nil
}

func square(d, i int) int {
	if i == 0 {
		return 1
	}
	return d * square(d, i-1)
}
