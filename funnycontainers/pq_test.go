package funnycontainers

import "testing"

func TestHeapifyPush(t *testing.T) {
	var maxIntLess Less = func(a interface{}, b interface{}) bool {
		return a.(int) < b.(int)
	}
	pq := NewPriorityQueue(maxIntLess)
	pq.Push(3)
	pq.Push(7)
	pq.Push(21)
	pq.Push(2)
	if pq.Peek() != 21 {
		t.Fail()
	}
}
func TestHeapifyWithPop(t *testing.T) {
	var maxIntLess Less = func(a interface{}, b interface{}) bool {
		return a.(int) < b.(int)
	}
	pq := NewPriorityQueue(maxIntLess)
	pq.Push(3)
	pq.Push(7)
	pq.Push(21)
	pq.Push(2)
	pq.Pop()
	if pq.Pop() != 7 {
		t.Fail()
	}
}
