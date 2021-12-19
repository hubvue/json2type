package util

import (
	"regexp"
	"strings"
)

// SnakeToCamel snake case to camel case
func SnakeToCamel(str string, firstUpper bool) string {
	if len(str) == 0 {
		return str
	}
	keyMatchRegexp := regexp.MustCompile(`_([a-z|0-9])`)
	converted := keyMatchRegexp.ReplaceAllFunc([]byte(str), func(match []byte) []byte {
		matchStr := string(match)
		key := matchStr[1:]
		return []byte(strings.ToUpper(key))
	})
	// first uppercase
	convertedStr := string(converted)
	if firstUpper {
		convertedStr = strings.ToUpper(string(convertedStr[0])) + string(convertedStr[1:])
	}

	return convertedStr
}
