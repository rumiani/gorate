package helpers

func Capitalize(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		if r >= 'a' && r <= 'z' {
			runes[i] = r - 32 // convert to uppercase
			break
		}
	}
	return string(runes)
}
