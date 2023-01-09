package mymethods

import "strings"

func RemoveElement(s string, element byte) string {
	for i := 0; i < len(s); i++ {
		if s[i] == element {
			s = s[:i] + s[i+1:]
		}
	}
	return s
}

func RemoveElementFast(str, element string) string {
	for true {
		slashIndex := strings.LastIndex(str, element)
		if slashIndex != -1 {
			str = str[:slashIndex] + str[slashIndex+1:]
		} else {
			break
		}
	}
	return str
}
