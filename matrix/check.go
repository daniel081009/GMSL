package matrix

/*
check square Matrix
정방행렬인지 확인
*/
func (m Matrix) CheckSquare() bool {
	return m.I() == m.J()
}

/*
check Matrix equal
행렬이 같은지 확인
*/
func (m Matrix) CheckEqual(md Matrix) bool {
	if !m.CheckEqualSize(md) {
		return false
	}
	for i := range m {
		for j := range m[i] {
			if m[i][j] != md[i][j] {
				return false
			}
		}
	}
	return true
}

func (m Matrix) I() int {
	return len(m)
}
func (m Matrix) J() int {
	return len(m[0])
}

/*
check equal Matrix size
행렬 size가 같은지 확인 (주의 : 완벽하게 행과 열이 같아야함)
*/
func (m Matrix) CheckEqualSize(md Matrix) bool {
	if len(m) != len(md) || len(m[0]) != len(md[0]) {
		return false
	}
	return true
}
