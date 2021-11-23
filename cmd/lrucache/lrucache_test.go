package lrucache

import(
	"testing"
) 


func Test_EmptyCache(t *testing.T){
	c := NewLRUCache(1000)
	invalid := c.Get(5)
	if invalid != -1 {
		t.Fail()
	}
}

func Test_ValidSingleEntry(t *testing.T){
	c := NewLRUCache(10)
	c.Set(1,4)
	valid := c.Get(1)
	if valid != 4 {
		t.Fail()
	}
	invalid := c.Get(2)
	if invalid != -1 {
		t.Fail()
	}
}

func Test_OverCapacity(t *testing.T){
	c := NewLRUCache(5)
	c.Set(1,5)
	c.Set(2,8)
	c.Set(3,11)
	c.Set(4,14)
	c.Set(5,17)
	c.Set(6,20)
	// test latest entry
	// leaving it here will not disrupt the order of the LRU cache
	latest := c.Get(6)
	if latest != 20 {
		t.Fail()
	}
	// test oldest entry
	oldest := c.Get(2)
	if oldest != 8 {
		t.Fail()
	}
	// test expired cache entry
	expired := c.Get(1)
	if expired != 1 {
		t.Fail()
	}
}
