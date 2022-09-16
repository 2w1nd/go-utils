package strs

func NextAsciiStr(c rune) rune {
	i := int(c)
	i++
	c = rune(i)
	return c
}
