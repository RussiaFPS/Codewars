package consonantValue

import (
	"strings"
)

func solve(str string) int {
	var max,now int

	for _,alp:=range []string{"a", "e", "i", "o", "u"}{
		str=strings.ReplaceAll(str,alp," ")
	}

	for _,value:=range str{
		if value==32{
			now=0
		}else{
			now+=int(rune(value)-96)
		}
		if max<now{
			max=now
		}
	}

	return max
}
