package matrixlib

import (
	"errors"
	"fmt"
)

type matrixop func(int,int) int

type matrix struct {
	size   int     "Size of matrix"
	values [][]int "Values of matrix"
}

type other matrix

func main() {
	var m1 matrix
	fmt.Println(m1)

	var m2 other = other(m1)
	fmt.Println(m2)

	m1 = matrix{3,
		[][]int{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8},
		},
	}
	fmt.Println(m1.values[2][1])
}

func NewMatrix(size int) (*matrix, error) {
	if size < 1{
		err := errors.New("Invalid size")
		return &matrix{}, err
	}

	result := &matrix{}
	result.size = size
	result.values = make([][]int, size)
	for i, _ := range result.values{
		result.values[i] = make([]int, size)
	}
	return result, nil
}


func Operate(lhs, rhs *matrix, operation matrixop) (*matrix, error) {
	var err error

	if lhs.size != rhs.size {
		err = errors.New("size should be same")
	}
	if lhs.values == nil || rhs.values == nil{
		err = errors.New("Nil values")
	}
	result, _ := NewMatrix(lhs.size)

	for i, row := range lhs.values{
		for j, data := range row{
			result.Set(i,j, operation(data, rhs.values[i][j]))
		}
	}

	return result, err
}

func addop(lvalue, rvalue int) int{
	return lvalue + rvalue
}

func subop(lvalue, rvalue int) int{
	return lvalue - rvalue
}

func sumsqrval(lvalue, rvalue int) int{
	return lvalue*lvalue + rvalue*rvalue
}

func Subtract(lhs, rhs *matrix) (*matrix, error) {
	return Operate(lhs, rhs, subop)
}

func Add(lhs, rhs *matrix) (*matrix, error) {
	return Operate(lhs, rhs, addop)
}

func (m *matrix) Set(row, col, val int) error {
	m.values[row][col] = val
 	return nil
}

func Get(m *matrix, row, col int) (int, error){
	return m.values[row][col], nil
}

func (m matrix) GetSize() int{
	return m.size
}

func (m matrix) Get(row, col int) (int, error){
	return m.values[row][col], nil
}