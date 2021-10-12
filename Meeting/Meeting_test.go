package Meeting

import "testing"

type dataTable struct {
	input string
	res string
	want string
}

func TestMeeting(t *testing.T){
	tests:=[]dataTable{
		{input: "Alexis:Wahl;John:Bell;Victoria:Schwarz;Abba:Dorny;Grace:Meta;Ann:Arno;Madison:STAN;Alex:Cornwell;Lewis:Kern;Megan:Stan;Alex:Korn",want: "(ARNO, ANN)(BELL, JOHN)(CORNWELL, ALEX)(DORNY, ABBA)(KERN, LEWIS)(KORN, ALEX)(META, GRACE)(SCHWARZ, VICTORIA)(STAN, MADISON)(STAN, MEGAN)(WAHL, ALEXIS)"},
	}
	for _,test:=range tests{
		test.res=Meeting(test.input)
		if test.res!=test.want{
			t.Errorf("\n Input:%s;\n Res:%s;\n Want:%s",test.input,test.res,test.want)
		}
	}
}