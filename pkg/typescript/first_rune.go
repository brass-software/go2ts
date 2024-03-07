package typescript

func firstRune(s string) rune {
	for _, r := range s {
		return r
	}
	panic("empty string")
}
