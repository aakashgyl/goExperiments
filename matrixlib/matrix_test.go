package matrixlib_test

import (
	"testing"
)

import "github.com/Infoblox-CTO/matrixlib"

func TestNewMatrix(t *testing.T) {
	size := 3
	m1, err := matrixlib.NewMatrix(size)

	if err != nil{
		t.Log("Valid Size, error returned")
		t.Fail()
	}

	if m1.GetSize() != size{
		t.Log("Size not properly set")
		t.Fail()
	}

	if val, err:= m1.Get(size-1,size-1); err != nil || val != 0 {
		t.Log("Could not correctly access last value in matrix")
		t.Fail()
	}
}

func TestNewMatrixNegSize(t *testing.T) {
	_, err := matrixlib.NewMatrix(-1)
	if err == nil{
		t.Log("Invalid Size should be an error")
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	size := 2
	m1, _ := matrixlib.NewMatrix(size)
	m2, _ := matrixlib.NewMatrix(size)

	var datam1 [4]int = [4]int{2,2,4,4}
	var datam2 [4]int = [4]int{1,1,1,1}
	var result [4]int = [4]int{3,3,5,5}

	currind := 0
	for i:=0; i < size; i++{
		for j := 0; j < size; j++{
			m1.Set(i,j,datam1[currind])
			m2.Set(i,j,datam2[currind])
			currind++
		}
	}

	currind=0
	m3, _ := matrixlib.Add(m1,m2)
	for i:=0; i < size; i++{
		for j := 0; j < size; j++{
			if val, _ := m3.Get(i,j); val != result[currind]{
				t.Log("Add failed")
				t.Fail()
			}
			currind++
		}
	}
}

func TestSub(t *testing.T) {
	size := 2
	m1, _ := matrixlib.NewMatrix(size)
	m2, _ := matrixlib.NewMatrix(size)

	var datam1 [4]int = [4]int{2,2,4,4}
	var datam2 [4]int = [4]int{1,1,1,1}
	var result [4]int = [4]int{1,1,3,3}

	currind := 0
	for i:=0; i < size; i++{
		for j := 0; j < size; j++{
			m1.Set(i,j,datam1[currind])
			m2.Set(i,j,datam2[currind])
			currind++
		}
	}

	currind=0
	m3, _ := matrixlib.Subtract(m1,m2)
	for i:=0; i < size; i++{
		for j := 0; j < size; j++{
			if val, _ := m3.Get(i,j); val != result[currind]{
				t.Log("Add failed")
				t.Fail()
			}
			currind++
		}
	}
}

func TestSub2(t *testing.T) {
	size := 2
	m1, _ := matrixlib.NewMatrix(size)
	m2, _ := matrixlib.NewMatrix(size)

	var datam1 [4]int = [4]int{-1,-2,-3,-4}
	var datam2 [4]int = [4]int{-1,0,-1,-8}
	var result [4]int = [4]int{-0,-2,-2,4}

	currind := 0
	for i:=0; i < size; i++{
		for j := 0; j < size; j++{
			m1.Set(i,j,datam1[currind])
			m2.Set(i,j,datam2[currind])
			currind++
		}
	}

	currind=0
	m3, _ := matrixlib.Subtract(m1,m2)
	for i:=0; i < size; i++{
		for j := 0; j < size; j++{
			if val, _ := m3.Get(i,j); val != result[currind]{
				t.Log("Add failed")
				t.Fail()
			}
			currind++
		}
	}
}