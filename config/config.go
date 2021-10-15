package config

import "strings"

func ToCamelCase(s string) string {
	s2 := strings.ToLower(s)
	s1 := string(s2[0])
	for i := 1; i < len(s); i++ {
		if string(s2[i-1]) == "_" {
			s1 = s1[:len(s1)-1]
			s1 += strings.ToUpper(string(s2[i]))
		} else {
			s1 += string(s2[i])
		}
	}
	return s1
}

func Format(m map[string]string, options... func(string)string) map[string]string {
	if m == nil {
		return m
	}
	transform1 := ToCamelCase
	if len(options) > 0 {
		if options[0] == nil {
			return m
		} else {
			transform1 = options[0]
		}
	}
	x := make(map[string]string)
	for k, v := range m {
		k2 := transform1(k)
		x[k2] = v
	}
	return x
}
