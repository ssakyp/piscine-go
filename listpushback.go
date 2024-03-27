package piscine

type NodeM struct { // NodeL structure represents a node in a linked list. It contains two fields:Data, which is of type interface{}, is used to store the data associated with the node.
	// The interface{} type allows Data to hold values of any type.
	// Next, which is a pointer to the next node in the linked list. This field allows nodes to be linked together, forming a chain of nodes.
	Data interface{}
	Next *NodeM
}

type ListA struct { // This List structure is used to implement a linked list. It contains two fields:
	Head *NodeM // Head is a pointer to a NodeL struct. ; NodeL itself is a structure that holds data and a reference to the next node in the linked list. ;Head points to the first node (or the head) of the linked list.
	Tail *NodeM // Tail is a pointer to a NodeL struct.; NodeL is a structure that represents a node in a linked list.; Tail points to the last node (or the tail) of the linked list.
}

func ListPushBackA(l *ListA, data interface{}) { // The ListPushBack function you're referring to is designed to append a new node containing the provided data to the end of the linked list.
	// It takes a pointer to a List structure (l) and the data to be inserted (data), which can be of any type due to the use of interface{}
	newNode := &NodeM{Data: data, Next: nil} // creates a new node (newNode) for the linked list. Let's break down what this line does:
	// &NodeL{Data: data, Next: nil}: This part initializes a new instance of the NodeL structure.
	// The Data field is set to the provided data, and the Next field is set to nil.
	// newNode := ...: This part assigns the address of the newly created node to the variable newNode.
	// So, overall, this line creates a new node with the provided data and sets its Next pointer to nil, indicating that it is the last node in the list (as it will be appended to the end).
	if l.Head == nil { // The condition l.Head == nil is checking if the Head pointer of the list (l) is nil, which indicates whether the list is empty or not.

		l.Head = newNode // This line is executed when the list is empty (l.Head == nil). In this case, since there are no nodes in the list yet, the new node (newNode)
		// becomes both the first node (head) and the last node (tail) of the list.
		// Therefore, the Head and Tail pointers of the list (l) are both assigned to point to the new node.

		l.Tail = newNode

	} else {

		l.Tail.Next = newNode // In the code l.Tail.Next = newNode, newNode is being assigned to the Next pointer of the current tail node in the list l.
		// This line is executed when the list is not empty (l.Head != nil). In this case, the new node (newNode) needs to be appended to the end of the list.
		l.Tail = newNode
	}
}
