package dot

import (
	"strings"
	"unicode"
)

// SplitCamelCase разбивает строку по смене регистра
func SplitCamelCase(s string) []string {
	if s == "" {
		return nil
	}

	var result []string
	runes := []rune(s)
	start := 0

	for i := 1; i < len(runes); i++ {
		if unicode.IsUpper(runes[i]) {
			// если текущая буква заглавная
			if unicode.IsLower(runes[i-1]) {
				// переход строчная -> заглавная
				result = append(result, string(runes[start:i]))
				start = i
			} else if i+1 < len(runes) && unicode.IsLower(runes[i+1]) {
				// случай: несколько заглавных, затем строчная
				// пример: "DBMSKey" -> "DBMS", "Key"
				result = append(result, string(runes[start:i]))
				start = i
			}
		}
	}

	// добавляем последний кусок
	result = append(result, string(runes[start:]))

	return result
}

func toJoinedLower(src string, joiner string) string {
	split := SplitCamelCase(src)
	for idx := range split {
		split[idx] = strings.ToLower(split[idx])
	}

	return strings.Join(split, joiner)
}

func ToSnakeCase(s string) string {
	return toJoinedLower(s, "_")
}

func ToKebabCase(s string) string {
	return toJoinedLower(s, "-")
}
