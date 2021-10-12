package stringCase

import "strings"

func toWeirdCase(str string)string{
	var res string

	for _,value:=range strings.Split(str," "){
		for index,runes:=range value{
			if index%2==0{
				res+=strings.ToUpper(string(runes))
			}else{
				res+=strings.ToLower(string(runes))
			}
		}
		res+=" "
	}
	return strings.TrimSpace(res)
}