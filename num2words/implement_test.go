package num2words

import (
	"testing"
)

func TestZero(t *testing.T){
	if ToWords(0) != "zero"{
		t.Log("0 is not zero")
		t.Fail()
	}
}

func TestThirteen(t *testing.T){
	if ToWords(13) != "thirteen"{
		t.Log("13 is not thirteen")
		t.Fail()
	}
}

func TestSomeNumbers(t *testing.T){
	if ToWords(239) != "not found"{
		t.Log("239 is not unknown")
	}

	if ToWords(1111) != "not found"{
		t.Log("1111 is not unknown")
	}
}

//pass array index as input and output compared to testcases
func TestNumbers(t *testing.T){
	testcases := map[int]string{
		1:"one",
		2:"two",
		111:"one hundred eleven",
	}
	//testcases := [...]string{"zero","one","two","three"}

	for i, val := range testcases{
		if r := ToWords(i); r != val{
			t.Logf("Input %d, expected %s, got %s\n", i, val, r)
			t.Fail()
		}
	}
}