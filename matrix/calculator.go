package matrix

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
행렬을 행렬로 나누기
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
행렬을 int로 나누기
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
matrix lu decomposition
행렬을 lu 분해
*/
func (m Matrix) LU() (Matrix, Matrix) {
	if !m.CheckSquare() {
		return nil, nil
	}
	l := Matrix{}
	l.Create(len(m), len(m))
	u := Matrix{}
	u.Create(len(m), len(m))

	n := len(m)
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if j == i {
				l[i][j] = 1
			} else {
				sum := 0
				for k := 0; k < j; k++ {
					sum += l[i][k] * u[k][j]
				}
				l[i][j] = (m[i][j] - sum) / u[j][j]
			}
		}

		// Compute the elements of the U matrix
		for j := i; j < n; j++ {
			sum := 0
			for k := 0; k < i; k++ {
				sum += l[i][k] * u[k][j]
			}
			u[i][j] = m[i][j] - sum
		}
	}

	return l, u
}
