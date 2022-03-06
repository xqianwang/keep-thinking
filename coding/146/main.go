package main

import "fmt"

func main()  {
	lru := Constructor(2)
	lru.Put(1,1)
	lru.Put(2,2)
	fmt.Println(lru.Get(1))
	lru.Put(3,3)
	fmt.Println(lru.Get(2))
}

type LRUCache struct {
	capacity int
	values map[int]*Node
	head *Node
	tail *Node
}


func Constructor(capacity int) LRUCache {
	lru := LRUCache{capacity: capacity}
	lru.values = make(map[int]*Node)
	lru.head = &Node{key:0, val:0}
	lru.tail = &Node{key:0, val:0}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}


func (this *LRUCache) Put(key int, value int)  {
	if node, ok := this.values[key]; ok {
		node.val = value
		this.moveToHead(node)
	} else {
		if len(this.values) == this.capacity {
			this.removeTail()
			delete(this.values, key)
		}

		node := &Node{key: key, val: value}
		this.values[key] = node
		this.addNode(node)
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.values[key]; ok {
		this.moveToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) moveToHead(node *Node) {
	if node.next != nil {
		node.next.prev = node.prev
	}
	node.prev.next = node.next
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
	node.prev = this.head
}

func (this *LRUCache) addNode(node *Node) {
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
	node.prev = this.head
}

func (this *LRUCache) removeTail() {
	this.tail.prev.prev.next = this.tail
	this.tail.prev = this.tail.prev.prev
}


type Node struct {
	key int
	val int
	next *Node
	prev *Node
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
