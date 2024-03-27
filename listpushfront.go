package piscine

type NodeN struct {
	Data interface{}
	Next *NodeN
}

type ListB struct {
	Head *NodeN
	Tail *NodeN
}

// ListPushFront inserts a new element at the beginning of the list
func ListPushFrontB(l *ListB, data interface{}) {
	// Create a new node with the provided data and a Next pointer set to nil
	newNode := &NodeN{Data: data, Next: nil}

	// Check if the list is empty
	if l.Head == nil {
		// If the list is empty, set both Head and Tail pointers to the new node
		l.Head = newNode
		l.Tail = newNode
	} else {
		// If the list is not empty, set the Next pointer of the new node to the current head of the list
		newNode.Next = l.Head
		// Update the Head pointer of the list to point to the new node
		l.Head = newNode
	}
}
