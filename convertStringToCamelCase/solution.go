package convertStringToCamelCase

import (
	"strings"
)

func ToCamelCase(s string) string {
	words:=strings.Split(strings.NewReplacer("_"," ","-"," ").Replace(s)," ")
	return words[0]+strings.ReplaceAll(strings.Title(strings.Join(words[1:]," "))," ","")
}