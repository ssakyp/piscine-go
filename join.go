package piscine

func Join(strs []string, sep string) string {
	joinedString := ""
	for i, str := range strs {
		joinedString += str
		if i < len(strs)-1 {
			joinedString += sep
		}
	}
	return joinedString
}
