package iteration

func Repeat(character string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}
func Divide(one, two int) int {
	return one / two
}
