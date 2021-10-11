package stringCase

import (
	"testing"
)

type testTable struct {
	input string
	want string
}

func TestToWeirdCase(t *testing.T) {
	tests:=[]testTable{
		{input:"String",want:"StRiNg" },
		{input: "Weird string case",want: "WeIrD StRiNg CaSe"},
		{input: "abc def",want: "AbC DeF" },
		{input: "ABC",want: "AbC"},
	}
	for _,test:=range tests{
		res:= toWeirdCase(test.input)
		if res!=test.want {
			t.Errorf("Input == %s ; Res == %s; Want == %s \n", test.input,res, test.want)
		}
	}
}