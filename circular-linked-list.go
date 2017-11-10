package cqlnls

type Node struct {
	Value          interface{}
	Next, Previous *Node
}

type CircularLinkedList struct {
	head, tail, current *Node
	len                 int
}

func (c *CircularLinkedList) addInit(n *Node) {
	c.head = n
	c.tail = n
	c.current = n
}

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

func (c *CircularLinkedList) Length() int {
	return c.len
}

func (c *CircularLinkedList) Current() *Node {
	if c.current != nil {
		return c.current
	}

	return c.head
}

func (c *CircularLinkedList) Next() *Node {
	if c.current != nil {
		c.current = c.current.Next
		return c.current
	}

	return nil
}

func (c *CircularLinkedList) Previous() *Node {
	if c.current != nil {
		c.current = c.current.Previous
		return c.current
	}

	return nil
}

func (c *CircularLinkedList) AddBefore(n *Node) *Node {
	if c.len == 0 {
		c.addInit(n)
	} else {
		c.current.Previous.Next = n
		c.current.Previous = n
		n.Previous = c.current.Previous
		n.Next = c.current
		c.current = n
	}

	c.len++

	return c.current
}

func (c *CircularLinkedList) AddAfter(n *Node) *Node {
	if c.len == 0 {
		c.addInit(n)
	} else {
		c.current.Next.Previous = n
		c.current.Next = n
		n.Previous = c.current
		n.Next = c.current.Next
		c.current = n
	}

	c.len++

	return c.current
}

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
