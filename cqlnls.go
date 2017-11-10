package cqlnls

// Node is an element block that contains the data and has pointers to other
// Nodes that it's connected
type Node struct {
	Value          interface{}
	Next, Previous *Node
}

// CircularLinkedList provides the interface for Nodes to connect
// add and remove nodes from the structure
type CircularLinkedList struct {
	head, tail, current *Node
	len                 int
}

func (c *CircularLinkedList) addInit(n *Node) {
	c.head = n
	c.tail = n
	c.current = n
}

// Prepend is used to add a Node before the head of the list
func (c *CircularLinkedList) Prepend(n *Node) *Node {
	if c.len == 0 {
		c.addInit(n)
	} else {
		c.head.Previous = n
		c.tail.Next = n
		n.Next = c.head
		n.Previous = c.tail
		c.head = n
		c.current = n
	}

	c.len++

	return c.current
}

// Append is used to add a Node after the tail of the list
func (c *CircularLinkedList) Append(n *Node) *Node {
	if c.len == 0 {
		c.addInit(n)
	} else {
		c.head.Previous = n
		c.tail.Next = n
		n.Next = c.head
		n.Previous = c.tail
		c.tail = n
		c.current = n
	}

	c.len++

	return c.current
}

// Length returns the size of the list (How many Nodes are in the list)
func (c *CircularLinkedList) Length() int {
	return c.len
}

// Current returns the Node that is active in the list
func (c *CircularLinkedList) Current() *Node {
	if c.current != nil {
		return c.current
	}

	return c.head
}

// Next returns the following Node of the current one
func (c *CircularLinkedList) Next() *Node {
	if c.current != nil {
		c.current = c.current.Next
		return c.current
	}

	return nil
}

// Previous returns the Node before the current one
func (c *CircularLinkedList) Previous() *Node {
	if c.current != nil {
		c.current = c.current.Previous
		return c.current
	}

	return nil
}

// AddBefore adds a new Node before the current one
func (c *CircularLinkedList) AddBefore(n *Node) *Node {
	switch {
	case c.len == 0:
		c.addInit(n)
	case c.current == c.head:
		return c.Prepend(n)
	default:
		c.current.Previous.Next = n
		c.current.Previous = n
		n.Previous = c.current.Previous
		n.Next = c.current
		c.current = n
	}

	c.len++

	return c.current
}

// AddAfter adds a new Node after the current one
func (c *CircularLinkedList) AddAfter(n *Node) *Node {
	switch {
	case c.len == 0:
		c.addInit(n)
	case c.current == c.tail:
		return c.Append(n)
	default:
		c.current.Next.Previous = n
		c.current.Next = n
		n.Previous = c.current
		n.Next = c.current.Next
		c.current = n
	}

	c.len++

	return c.current
}

// Remove is disconnecting the current node from the list
func (c *CircularLinkedList) Remove() *Node {
	if c.len == 0 {
		return nil
	}

	c.current.Previous.Next = c.current.Next
	c.current.Next.Previous = c.current.Previous

	c.current = c.current.Next
	c.len--

	return c.current
}
