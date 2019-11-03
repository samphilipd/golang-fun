package main

import "strings"

func defangIPaddr(address string) string {
	var defanged strings.Builder

	for _, rn := range address {
		if rn == '.' {
			defanged.WriteString("[.]")
		} else {
			defanged.WriteRune(rn)
		}
	}

	return defanged.String()
}
