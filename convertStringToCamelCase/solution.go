package convertStringToCamelCase

import (
	"strings"
)

func ToCamelCase(s string) string {
	var res string

	for index,value:=range strings.Split(strings.NewReplacer("_"," ","-"," ").Replace(s)," "){
		if index==0{
			res+=value
		}else{
			res+=strings.ToUpper(string(value[0]))+value[1:]
		}
	}

	return res
}