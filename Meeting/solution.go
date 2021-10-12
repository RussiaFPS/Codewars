package Meeting

import (
	"fmt"
	"sort"
	"strings"
)

func Meeting(s string) string {
	var res string
	var persons[]string

	for _,value:=range strings.Split(s,";"){
		buf:=strings.Split(value,":")
		persons=append(persons,fmt.Sprint(strings.ToUpper(buf[1])," ",strings.ToUpper(buf[0])))
	}
	sort.Strings(persons)
	for _,value:=range persons{
		buf:=strings.Split(value," ")
		res+=fmt.Sprintf("(%s, %s)",buf[0],buf[1])
	}

	return res
}