package funnycontainers

import "testing"

var maxIntLess Less = func(a interface{}, b interface{}) bool {
	return a.(int) < b.(int)
}
var minIntLess Less = func(a interface{}, b interface{}) bool {
	return a.(int) > b.(int)
}

func TestHeapifyPush(t *testing.T) {

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

func TestIsEmpty(t *testing.T) {
	pq := NewPriorityQueue(maxIntLess)
	pq.Push(2)
	pq.Push(3)
	pq.Pop()
	pq.Pop()
	if !pq.IsEmpty() {
		t.Fail()
	}
}

func TestTypicalTopKProblem(t *testing.T) {
	elements := []int{2, 3, 11, 91, 4, 7, 33, 55}
	k := 3
	pq := NewPriorityQueue(minIntLess)
	for _, element := range elements {
		pq.Push(element)
		if pq.Size() > k {
			pq.Pop()
		}
	}
	result := make([]int, 0, k)
	for i := 0; i < k; i++ {
		result = append(result, pq.Pop().(int))
	}
	if result[0] != 33 {
		t.Fail()
	}
	if result[1] != 55 {
		t.Fail()
	}
	if result[2] != 91 {
		t.Fail()
	}
}
