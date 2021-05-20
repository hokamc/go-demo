package iteration

func Repeat(character string, repeat int) string {
	var repeated string
	for i := 0; i < repeat; i++ {
		repeated = repeated + character
	}
	return repeated
}
