package prose

import "strings"

func JoinWithCommas(phrases []string) string {
	if len(phrases) == 0 {
		return ""
	}
	result := strings.Join(phrases[:len(phrases)-1], ", ")
	if len(phrases) == 1 {
		return phrases[0]
	}
	if len(phrases) == 2 {
		result += " and "
	} else {
		result += ", and "
	}

	result += phrases[len(phrases)-1]
	return result
}
