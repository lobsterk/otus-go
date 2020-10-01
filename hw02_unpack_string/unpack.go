package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var saveChar, l string
	var res strings.Builder

	for i, r := range s {
		if unicode.IsDigit(r) {
			if (i+1) <= len(s)-1 && unicode.IsDigit(int32(s[i+1])) {
				return "", ErrInvalidString
			}
			if (i-1) >= 0 && i < len(s) {
				num, _ := strconv.ParseInt(string(r), 10, 64)
				l = strings.Repeat(saveChar, int(num)-1)
				res.WriteString(l)
			} else {
				return "", ErrInvalidString
			}
		} else {
			saveChar = string(r)
			res.WriteString(string(r))
		}
	}
	if len(res.String()) > 0 {
		return res.String(), nil
	}
	return "", nil
}
