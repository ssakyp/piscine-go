package main

import "github.com/01-edu/z01"

type ptr struct {
	x string
	y string
}

func setPoint(pointer *ptr, a string, b string) {
	pointer.x = a
	pointer.y = b
}

func main() {
	points := ptr{}

	setPoint(&points, "42", "21")

	text_x := "x = "
	text_x = text_x + points.x + ", "
	text_y := "y = "
	text_y += points.y
	for _, char := range text_x {
		z01.PrintRune(char)
	}
	for _, char := range text_y {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
