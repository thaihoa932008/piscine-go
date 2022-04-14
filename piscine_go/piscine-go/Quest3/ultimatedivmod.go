package piscine

func UltimateDivMod(a *int, b *int) {
	c := *a
	*a = c / *b
	*b = c % *b
}
