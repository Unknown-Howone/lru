package lru

import (
	"fmt"
	"math/rand/v2"
)

type Node[T any] struct {
	Key        int
	Value      T
	Prev, Next *Node[T]
}

type Cache[T any] struct {
	Keys       map[int]*Node[T]
	Len, Cap   int
	Head, Tail *Node[T]
}

func newNode[T any](key int, value T) *Node[T] {
	return &Node[T]{
		Key:   key,
		Value: value,
	}
}

func NewCache[T any](cap int) *Cache[T] {
	c := &Cache[T]{
		Keys: make(map[int]*Node[T]),
		Len:  0,
		Cap:  cap,
		Head: newNode[T](0, *new(T)),
		Tail: newNode[T](0, *new(T)),
	}
	c.Head.Next = c.Tail
	c.Tail.Prev = c.Head
	return c
}

func (c *Cache[T]) newKey() int {
	if c == nil {
		fmt.Println("Cache is not initialized")
		return -1
	}

	for {
		key := rand.N(c.Cap+1) + 1
		if _, ok := c.Keys[key]; !ok {
			return key
		}
	}
}

func (c *Cache[T]) RemoveNode(node *Node[T]) {
	if c == nil {
		fmt.Println("Cache is not initialized")
		return
	}

	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (c *Cache[T]) AddToHead(node *Node[T]) {
	if c == nil {
		fmt.Println("Cache is not initialized")
		return
	}

	node.Prev = c.Head
	node.Next = c.Head.Next
	c.Head.Next.Prev = node
	c.Head.Next = node
}

func (c *Cache[T]) MoveToHead(node *Node[T]) {
	if c == nil {
		fmt.Println("Cache is not initialized")
		return
	}

	c.RemoveNode(node)
	c.AddToHead(node)
}

func (c *Cache[T]) RemoveTail() int {
	if c == nil {
		fmt.Println("Cache is not initialized")
		return -1
	}

	key := c.Tail.Prev.Key
	c.RemoveNode(c.Tail.Prev)
	return key
}

func (c *Cache[T]) Get(key int) (T, bool) {
	if c == nil {
		fmt.Println("Cache is not initialized")
		return *new(T), false
	}

	if node, ok := c.Keys[key]; ok {
		c.MoveToHead(node)
		return node.Value, ok
	}
	return *new(T), false
}

func (c *Cache[T]) Put(value T) int {
	if c == nil {
		fmt.Println("Cache is not initialized")
		return -1
	}

	key := c.newKey()
	if node, ok := c.Keys[key]; ok {
		node.Value = value
		c.MoveToHead(node)
	} else {
		node = newNode[T](key, value)
		c.Keys[key] = node
		c.AddToHead(node)
		c.Len++
		if c.Len > c.Cap {
			k := c.RemoveTail()
			delete(c.Keys, k)
			c.Len--
		}
	}
	return key
}

func (c *Cache[T]) Show() {
	if c == nil {
		fmt.Println("Cache is not initialized")
		return
	}

	pNode := c.Head.Next
	l := 1
	for l <= c.Len {
		fmt.Print(pNode, " > ")
		pNode = pNode.Next
		l++
	}
	fmt.Print("\nshow done\n")
}
