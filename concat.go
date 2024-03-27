package piscine

func Concat(str1 string, str2 string) string {
	counter := 0
	length := len(str1) + len(str2)
	concat := make([]rune, length)

	for _, char := range str1 {
		concat[counter] = char
		counter++
	}

	for _, char := range str2 {
		concat[counter] = char
		counter++
	}
	return string(concat[:counter])
}
