package convertStringToCamelCase

import "testing"

type dataTest struct {
	input string
	res string
	want string
}

func TestToCamelCase(t *testing.T) {
	tests:=[]dataTest{
		{input: "The_Stealth_Warrior",want:"TheStealthWarrior"},
		{input: "the-Stealth-Warrior",want: "theStealthWarrior"},
	}
	for _,test:=range tests{
		test.res=ToCamelCase(test.input)
		if test.res!=test.want{
			t.Errorf("\nInput:%s\nRes:%s\nWant:%s",test.input,test.res,test.want)
		}
	}
}