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
		switch value {
		case 32:
			now=0
		default:
			now+=int(value-96)
			if max<now{ max=now }
		}
	}
	return max
}