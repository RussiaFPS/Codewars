package consonantValue

import "testing"

type dataTest struct {
	input string
	res int
	want int
}

func TestSolve(t *testing.T){
	tests:=[]dataTest{
		{input: "a",want: 0},
		{input: "bcd",want: 9},
		{input: "zea",want: 26},
		{input: "aeiou",want: 0},
		{input: "codewars",want: 37},
	}
	for _,test:=range tests{
		test.res=solve(test.input)
		if test.res!=test.want{
			t.Errorf("Input:%s; Res==%d; Want==%d;",test.input,test.res,test.want)
		}
	}
}