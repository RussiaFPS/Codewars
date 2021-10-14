package meeting

import (
	"fmt"
	"sort"
	"strings"
)

func Meeting(s string) string {
	var persons[]string

	for _,value:=range strings.Split(s,";"){
		buf:=strings.Split(value,":")
		persons=append(persons,fmt.Sprintf("(%s, %s)",strings.ToUpper(buf[1]),strings.ToUpper(buf[0])))
	}
	sort.Strings(persons)
	return strings.Join(persons,"")
}