package main

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */

// @lc code=start

type pageListNode struct {
	key, value int
	prev, next *pageListNode
}

type LRUCache struct {
	capacity   int
	cache      map[int]*pageListNode
	head, tail *pageListNode
}

func Constructor(capacity int) LRUCache {
	head := &pageListNode{}
	tail := &pageListNode{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*pageListNode),
		head:     head,
		tail:     tail,
	}
}

// moveToTail 移动节点到链表尾部表示最近访问
func (this *LRUCache) moveToTail(node *pageListNode) {
	// 先删除节点
	node.prev.next = node.next
	node.next.prev = node.prev
	// 再插入到尾部
	this.insertBeforeTail(node)
}

func (this *LRUCache) insertBeforeTail(node *pageListNode) {
	node.prev = this.tail.prev
	node.next = this.tail
	this.tail.prev.next = node
	this.tail.prev = node
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.cache[key]; exists {
		this.moveToTail(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.cache[key]; exists {
		node.value = value
		this.moveToTail(node)
		return
	}
	if len(this.cache) == this.capacity {
		// 删除头节点的下一个节点（最久未使用）
		removed := this.head.next
		delete(this.cache, removed.key)
		this.head.next = removed.next
		removed.next.prev = this.head
	}
	node := &pageListNode{key: key, value: value}
	this.cache[key] = node
	this.insertBeforeTail(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
