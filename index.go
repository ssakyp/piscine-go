package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}

	for index, char := range s {
		if char == rune(toFind[0]) {
			substringMatch := true
			for i, subChar := range toFind {
				if index+i >= len(s) || rune(s[index+i]) != subChar {
					substringMatch = false
					break
				}
			}
			if substringMatch {
				return index
			}
		}
	}
	return -1
}
