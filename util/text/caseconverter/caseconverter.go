package caseconverter

import (
	"strings"
	"unicode"
)

// 驼峰式写法转为下划线写法
func Camel2Underline(str string) string {
	var buf strings.Builder
	for i, c := range str {
		if i > 0 && unicode.IsUpper(c) {
			buf.WriteRune('_')
		}
		buf.WriteRune(unicode.ToLower(c))
	}
	return buf.String()
}

// 下划线转首字母大写驼峰
func Underline2Camel(str string) string {
	parts := strings.Split(str, "_")
	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = strings.ToUpper(part[0:1]) + strings.ToLower(part[1:])
		}
	}
	return strings.Join(parts, "")
}

// 下划线转首字母小写驼峰
func Underline2camel(str string) string {
	parts := strings.Split(str, "_")
	for i, part := range parts {
		if len(part) > 0 {
			if i == 0 {
				parts[i] = strings.ToLower(part)
			} else {
				parts[i] = strings.ToUpper(part[0:1]) + strings.ToLower(part[1:])
			}
		}
	}
	return strings.Join(parts, "")
}

// 首字母大写
func Ucfirst(str string) string {
	if len(str) > 0 {
		return strings.ToUpper(str[0:1]) + str[1:]
	}

	return ""
}

// 首字母小写
func Lcfirst(str string) string {
	if len(str) > 0 {
		return strings.ToLower(str[0:1]) + str[1:]
	}

	return ""
}
