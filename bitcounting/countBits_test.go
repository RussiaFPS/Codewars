package bitcounting

import "testing"

type dataTest struct {
	input uint
	res int
	want int
}

func TestCountBits(t *testing.T){
	tests:=[]dataTest{
		{input: 0,want: 0},
		{input: 4,want: 1},
		{input: 7,want: 3},
		{input: 10,want: 2},
	}
	for _,test:=range tests{
		test.res=CountBits(test.input)
		if test.res!=test.want{
			t.Errorf("Input:%d; Res==%d; Want==%d;",test.input,test.res,test.want)
		}
	}
}