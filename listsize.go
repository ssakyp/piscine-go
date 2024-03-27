package piscine

type NodeO struct {
	Data interface{}
	Next *NodeO
}

type ListC struct {
	Head *NodeO
	Tail *NodeO
}

func ListSize(l *ListC) int {
	// Initialize a counter variable to count the number of elements
	count := 0

	// Traverse the linked list starting from the head
	current := l.Head
	for current != nil {
		// Increment the counter for each node encountered
		count++
		// Move to the next node
		current = current.Next
	}

	// Return the count, which represents the number of elements in the list
	return count
}

func ListPushFrontC(l *ListC, data interface{}) {
	newNode := &NodeO{Data: data, Next: nil}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
}
