package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	index := 0
	count := 0
	h := l

	for h != nil {
		index++
		h = h.Next
	}

	if pos <= index {
		for l != nil {
			if count == pos {
				return l
			}
			count++
			l = l.Next
		}
	}
	return nil
}
