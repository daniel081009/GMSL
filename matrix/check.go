package Matrix

/*
check square Matrix
정방행렬인지 확인
*/
func (m Matrix) CheckSquare() bool {
	return len(m) == len(m[0])
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
