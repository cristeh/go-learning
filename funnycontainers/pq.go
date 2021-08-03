package funnycontainers

// Dirty heap-based implementation of a priority queue, using 1-indexed array
//for easier reheapification operations.
type Less func(interface{}, interface{}) bool
type HeapPriorityQueue struct {
	data []interface{}
	size int
	less Less
}

func NewPriorityQueue(less Less) *HeapPriorityQueue {
	pq := &HeapPriorityQueue{
		data: make([]interface{}, 0, 16),
		less: less,
	}
	pq.data = append(pq.data, "")
	return pq
}
func (pq *HeapPriorityQueue) Push(element interface{}) {
	pq.size++
	pq.data = append(pq.data, element)
	pq.heapifyUp(pq.size)
}
func (pq *HeapPriorityQueue) Pop() interface{} {
	head := pq.data[1]
	pq.swap(1, pq.size)
	pq.size--
	pq.data = pq.data[:pq.size+1]
	pq.heapifyDown(1)
	return head
}
func (pq *HeapPriorityQueue) Peek() interface{} {
	return pq.data[1]
}
func (pq *HeapPriorityQueue) IsEmpty() bool {
	return pq.size == 0
}
func (pq *HeapPriorityQueue) Size() int {
	return pq.size
}

func (pq *HeapPriorityQueue) swap(i1 int, i2 int) {
	pq.data[i1], pq.data[i2] = pq.data[i2], pq.data[i1]
}

func (pq *HeapPriorityQueue) heapifyUp(k int) {
	for k > 1 && pq.less(pq.data[k/2], pq.data[k]) {
		pq.swap(k/2, k)
		k = k / 2
	}
}

func (pq *HeapPriorityQueue) heapifyDown(k int) {
	for 2*k <= pq.size {
		j := 2 * k
		if j < pq.size && pq.less(pq.data[j], pq.data[j+1]) {
			j++
		}
		if !pq.less(pq.data[k], pq.data[j]) {
			break
		}
		pq.swap(k, j)
		k = j
	}
}
