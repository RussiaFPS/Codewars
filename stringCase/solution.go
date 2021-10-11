package stringCase

import "strings"

func toWeirdCase(str string)string{
	var res string
	s:=strings.Split(str," ")
	for i:=0;i<len(s);i++{
		for j:=0;j<len(s[i]);j++{
			buf:=s[i]
			if j%2==0{
				res+=strings.ToUpper(string(buf[j]))
			}else{
				res+=strings.ToLower(string(buf[j]))
			}
		}
		if i<len(s)-1{
			res+=" "
		}
	}
	return res
}