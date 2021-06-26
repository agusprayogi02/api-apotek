package utils

import "strings"

func ExtactToken(str string) (r string) {
	const BearerSchema = "Tokenku"
	arrStr := strings.Split(str, " ")
	if len(arrStr) == 2 {
		if arrStr[0] != BearerSchema {
			r = ""
			return
		}
		r = arrStr[1]
		return
	}
	r = ""
	return
}
