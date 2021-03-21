package matrix

import (
	"errors"
	"math"
	"strconv"

	"github.com/AlexeyRyabichev/CompCompexity4RussiansAlgo/pkg/logger"
)

type Matrices struct {
	Left        Matrix `json:"left"`
	Right       Matrix `json:"right"`
	size        int
	blockSize   int
	blocksTotal int
}

type Matrix []Row
type Row []bool

func (m *Matrix) String() string {
	var ans string
	for i := 0; i < len(*m); i++ {
		for j := 0; j < len((*m)[i]); j++ {
			ans += strconv.FormatBool((*m)[i][j]) + "\t"
		}
		ans += "\n"
	}
	return ans
}

func (m *Matrices) String() string {
	return m.Left.String() + "\n" + m.Right.String()
}

func (m *Matrices) Validate() error {
	if len(m.Left) != len(m.Right) {
		err := errors.New("matrices of different sizes")
		logger.Sugar.Errorf("canot validate matrices. err: %v", err.Error())
		return err
	}
	return nil
}

func (m *Matrices) Prepare() error {
	m.size = len(m.Left)
	m.blockSize = int(math.Log2(float64(m.size)))
	m.blocksTotal = int(math.Ceil(float64(m.size / m.blockSize)))
	return nil
}

func NewSquareMatrix(size int) *Matrix {
	return NewMatrix(size, size)
}

func NewMatrix(row, col int) *Matrix {
	ans := make(Matrix, row)
	for i := 0; i < row; i++ {
		ans[i] = make(Row, col)
	}
	return &ans
}

func (m *Matrices) Multiply() (*Matrix, error) {
	logger.Sugar.Infof("matrix size: %v, blocks total: %v, block size: %v", m.size, m.blocksTotal, m.blockSize)

	ans := NewSquareMatrix(m.size)

	for i := 0; i < m.blocksTotal; i++ {
		left := m.Left.getColumns(i*m.blockSize, m.blockSize)
		right := m.Right.getRows(i*m.blockSize, m.blockSize)

		tmp := multiply(left, right, m.blockSize, m.size)

		ans.add(&tmp)
	}

	return ans, nil
}

func multiply(left, right *Matrix, blockSize, matrixSize int) Matrix {
	ans := NewSquareMatrix(matrixSize)

	maxSize := int(math.Pow(2, float64(blockSize)))
	rowsSums := NewMatrix(maxSize, matrixSize)

	betweenPowers := 1
	k := 0

	for i := 1; i < maxSize; i++ {
		(*rowsSums)[i] = or(
			rowsSums,
			i-int(math.Pow(2, float64(k))),
			right,
			len(*right)-k-1)

		if betweenPowers == 1 {
			betweenPowers = i + 1
			k += 1
		} else {
			betweenPowers -= 1
		}
	}

	for i := 0; i < matrixSize; i++ {
		(*ans)[i] = (*rowsSums)[toNum((*left)[i])]
	}

	return *ans
}

func toNum(row Row) int {
	ans := 0

	for i, v := range row {
		if v {
			ans += int(math.Pow(2, float64(len(row)-i-1)))
		}
	}

	return ans
}

func or(self *Matrix, self_index int, other *Matrix, other_index int) Row {
	row := make(Row, len((*self)[self_index]))
	copy((*self)[self_index], row)
	for i, v := range (*other)[other_index] {
		row[i] = row[i] || v
	}
	return row
}

func (m *Matrix) getRows(start, num int) *Matrix {
	ans := (*m)[start : start+num]
	return &ans
}

func (m *Matrix) getColumns(start, num int) *Matrix {
	actualNum := int(math.Min(float64(num), float64(len(*m))))
	rows := len(*m)
	ans := NewMatrix(rows, num)
	for i := 0; i < rows; i++ {
		for j := 0; j < actualNum; j++ {
			(*ans)[i][j] = (*m)[i][start+j]
		}
	}
	return ans
}

func (m *Matrix) toNum(i int) int {
	ans := 0

	for i, v := range (*m)[i] {
		if v {
			ans += int(math.Pow(2, float64(len(*m)-i-1)))
		}
	}

	return ans
}

func (m *Matrix) add(additional *Matrix) {
	for i, row := range *m {
		for j, _ := range row {
			(*m)[i][j] = (*m)[i][j] || (*additional)[i][j]
		}
	}
}
