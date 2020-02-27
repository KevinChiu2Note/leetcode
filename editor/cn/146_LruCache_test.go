package cn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//运用你所掌握的数据结构，设计和实现一个 LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
//
// 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
//写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新
//的数据值留出空间。
//
// 进阶:
//
// 你是否可以在 O(1) 时间复杂度内完成这两种操作？
//
// 示例:
//
// LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
//
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // 返回  1
//cache.put(3, 3);    // 该操作会使得密钥 2 作废
//cache.get(2);       // 返回 -1 (未找到)
//cache.put(4, 4);    // 该操作会使得密钥 1 作废
//cache.get(1);       // 返回 -1 (未找到)
//cache.get(3);       // 返回  3
//cache.get(4);       // 返回  4
//
// Related Topics 设计

//leetcode submit region begin(Prohibit modification and deletion)
type LRUCache struct {
	capacity     int
	doubleLinked linkedArray
	hashMap      map[int]*node
}

// node
type node struct {
	preNode  *node
	nextNode *node
	data     int
	key      int
}

// linked array
type linkedArray struct {
	headNode *node
	tailNode *node
	size     int
}

func (linkedArray *linkedArray) AddFirst(n *node) {
	n.nextNode = linkedArray.headNode.nextNode
	n.preNode = linkedArray.headNode
	linkedArray.headNode.nextNode.preNode = n
	linkedArray.headNode.nextNode = n
	linkedArray.size += 1
}

func (linkedArray *linkedArray) AddLast(n *node) {
	n.preNode = linkedArray.tailNode.preNode
	n.nextNode = linkedArray.tailNode
	linkedArray.tailNode.preNode.nextNode = n
	linkedArray.tailNode.preNode = n
	linkedArray.size += 1
}

func (linkedArray *linkedArray) Remove(n *node) {
	n.preNode.nextNode = n.nextNode
	n.nextNode.preNode = n.preNode
	linkedArray.size -= 1
}

func (linkedArray *linkedArray) RemoveLast() *node {
	tmp := linkedArray.tailNode.preNode
	linkedArray.Remove(tmp)
	return tmp
}

func (linkedArray *linkedArray) RemoveFirst() *node {
	tmp := linkedArray.headNode.nextNode
	linkedArray.Remove(tmp)
	return tmp
}

func (linkedArray *linkedArray) Size() int {
	return linkedArray.size
}

func Constructor(capacity int) LRUCache {
	headNode := &node{}
	tailNode := &node{}
	headNode.nextNode = tailNode
	tailNode.preNode = headNode
	return LRUCache{
		capacity: capacity,
		doubleLinked: linkedArray{
			headNode: headNode,
			tailNode: tailNode,
			size:     0,
		},
		hashMap: make(map[int]*node),
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.hashMap[key]; ok {
		// 优化，将查询的节点直接移到最前端
		//BenchMark 520 ns/op -> 320 ns/op
		//this.Put(key, v.data)
		v.preNode.nextNode = v.nextNode
		v.nextNode.preNode = v.preNode
		v.nextNode = this.doubleLinked.headNode.nextNode
		v.preNode = this.doubleLinked.headNode
		this.doubleLinked.headNode.nextNode = v
		v.nextNode.preNode = v
		return v.data
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	n := &node{
		data: value,
		key:  key,
	}
	// map 中已经有该节点，删除
	if node, ok := this.hashMap[key]; ok {
		this.doubleLinked.Remove(node)
		delete(this.hashMap, key)
	} else {
		// 满了
		if this.doubleLinked.Size() == this.capacity {
			delete(this.hashMap, this.doubleLinked.RemoveLast().key)
		}
	}
	this.doubleLinked.AddFirst(n)
	this.hashMap[key] = n
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestLruCache(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	assert.Equal(t, cache.Get(1), 1)

	cache.Put(3, 3)
	t.Log(cache.Get(2))
	assert.Equal(t, cache.Get(2), -1)

	cache.Put(4, 4)
	assert.Equal(t, cache.Get(1), -1)
	assert.Equal(t, cache.Get(3), 3)
	assert.Equal(t, cache.Get(4), 4)

}

func BenchmarkLruCache(b *testing.B) {
	cache := Constructor(2)
	for i := 0; i < b.N; i++ {
		cache.Put(1, 1)
		cache.Put(2, 2)
		cache.Get(1)
		cache.Put(3, 3)
		cache.Get(2)
		cache.Put(4, 4)
		cache.Get(1)
		cache.Get(3)
		cache.Get(4)
	}
}
