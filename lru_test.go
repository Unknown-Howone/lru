package lru

import (
	"fmt"
	"testing"
)

func TestNewNode(t *testing.T) {
	fmt.Println(newNode[string](1, "t1"))
	fmt.Println(newNode[int](1, 2))
	fmt.Println(newNode[map[int]int](1, make(map[int]int)))
}

func TestLRU(t *testing.T) {
	cache := NewCache[string](3)
	key := cache.Put("t1")
	fmt.Println(key)
	fmt.Println(cache.Put("t2"))
	fmt.Println(cache.Put("t3"))
	cache.Show()
	key4 := cache.Put("t4")
	cache.Show()
	fmt.Println(cache.Get(key))
	fmt.Println(cache.Get(key4))
}
