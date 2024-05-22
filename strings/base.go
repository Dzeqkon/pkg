package strings

import (
	"strings"
	"unicode/utf8"

	"github.com/asaskevich/govalidator"
)

/*
first = false: "aaa_bbb_ccc" -> "aaaBbbCcc"
first = true: "aaa_bbb_ccc" -> "AaaBbbCcc"
*/
func FirstCaseToUpper(str string, first bool) string {
	temp := strings.Split(str, "_")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		if y == 0 && !first {
			continue
		}
		for i := 0; i < len(vv); i++ {
			if i == 0 {
				vv[i] -= 32
				upperStr += string(vv[i])
			} else {
				upperStr += string(vv[i])
			}
		}
	}
	if first {
		return upperStr
	} else {
		return temp[0] + upperStr
	}
}

func CamelCaseToUnderscore(str string) string {
	return govalidator.CamelCaseToUnderscore(str)
}

func UnderscoreToCamelCase(str string) string {
	return govalidator.UnderscoreToCamelCase(str)
}

func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}
