package Matrix

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
	tmp.Create(len(m), len(m[0])).Write(m)
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
	tmp.Create(len(m), len(m[0])).Write(m)
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
	tmp.Create(len(m[0]), len(m))
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
	tmp.Create(len(m), len(m[0])).Write(m)
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
	if len(m[0]) != len(md) {
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
Matrix inverse
역행렬
*/
func (m Matrix) Inverse() Matrix {
	if !m.CheckSquare() {
		return nil
	}
	return nil
}
