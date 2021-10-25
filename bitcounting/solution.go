package bitcounting

import (
	"fmt"
	"strings"
)

func CountBits(input uint) int {
	return strings.Count(fmt.Sprintf("%b",input),"1")
}
