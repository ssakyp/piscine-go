package piscine

func StrRev(s string) string {
	sChangeable := []byte(s)
	i := 0

	for range sChangeable {
		i++
	}

	for j := 0; j < i/2; j++ {
		temp := sChangeable[j]
		sChangeable[j] = sChangeable[i-j-1]
		sChangeable[i-j-1] = temp
	}

	sFinalized := string(sChangeable)
	return sFinalized
}
