package helpers

//IntHeap initialise Heap
type IntHeap []int

//Len len of heap
func (h IntHeap) Len() int {
	return len(h)
}

//Less less elements from heap
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

//Swap swap elements in heap
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

//Push push elements into heap
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

//Pop pop elements from heap
func (h *IntHeap) Pop() interface{} {
	oldh := *h
	x := oldh[len(oldh)-1]
	newh := oldh[0 : len(oldh)-1]
	*h = newh
	return x
}
