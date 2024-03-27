package piscine

func BasicJoin(elems []string) string {
	final_string := ""

	for _, string := range elems {
		final_string += string
	}
	return final_string
}
