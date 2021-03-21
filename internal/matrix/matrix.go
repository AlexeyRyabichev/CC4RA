package matrix

import (
	"errors"
	"math"
	"strconv"

	"github.com/AlexeyRyabichev/CompCompexity4RussiansAlgo/pkg/logger"
)

type Matrices struct {
	Left  Matrix `json:"left"`
	Right Matrix `json:"right"`
	size  int
	logN  int
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
	m.logN = int(math.Log(float64(m.size)))
	return nil
}

func NewMatrix(size int) *Matrix {
	ans := make(Matrix, size)
	for i := 0; i < size; i++ {
		ans[i] = make(Row, size)
	}
	return &ans
}

func (m *Matrices) Multiply() (*Matrix, error) {
	ans := NewMatrix(m.size)

	return ans, nil
}
