package iteration

func Repeat(character string, rotations int) string {
	var repeated string

	for i := 0; i < rotations; i++ {
		repeated += character
	}

	return repeated
}
