package main

import (
	"fmt"
)


type Node struct {
	Left *Node
	Right *Node
	Val string
}

type Queue struct {
	Head *Node
	Tail *Node
	Length int
}

type Cache struct {
	Queue Queue
	Hash Hash
}

type Hash map[string]*Node


func NewCache() Cache{
	return Cache{Queue: NewQueue() , Hash: Hash{}}
}

func (c *Cache) Check(str string){
	node := &Node{}

	if val , ok := c.Hash[str]; ok {
		node = c.Remove(val)
	}else{
		node = &Node{Val:str}
	}
	c.Add(node)
	c.Hash[str] = node
}


func NewQueue() Queue{
	head:= &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head 

	queue := Queue{}
	queue.Head = head
	queue.Tail = tail
	queue.Length = 0
	return queue
}



func main(){
	fmt.Println("START CACHE")
	cache := NewCache()
	for _,word := range []string{""}{
		cache.Check(word)
		cache.Display()
	}
}