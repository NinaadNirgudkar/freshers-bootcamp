package dayone

import "fmt"

func MakeMatrix() {
	var matrix1 = Matrix{2, 2, [][]int{{1, 2}, {5, 2}}}
	var matrix2 = Matrix{2, 2, [][]int{{0, 3}, {8, 22}}}
	fmt.Println("matrix1", matrix1)
	fmt.Println("matrix2", matrix2)
	fmt.Println("addition of 2 metrices", matrix1.add(matrix2))
}

type Matrix struct {
	rows int
	cols int
	ele  [][]int
}

func (m *Matrix) getRows() int {
	return m.rows
}

func (m *Matrix) getCols() int {
	return m.rows
}

func (m *Matrix) setElement(row, col, val int) {
	m.ele[row][col] = val
}

func (m *Matrix) add(m2 Matrix) (result Matrix) {
	result.cols = m.cols
	result.rows = m.rows
	result.ele = [][]int{{0, 0}, {0, 0}}
	for i := 0; i < result.rows; i++ {
		for j := 0; j < result.cols; j++ {
			fmt.Println(i, j, result.ele[i][j])
			result.ele[i][j] = m.ele[i][j] + m2.ele[i][j]
		}
	}
	return
}
