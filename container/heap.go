package container

import "fmt"

type Heap []int

func up(h Heap) {
	l := len(h) - 1

	for h[l] < h[(l-1)/2] {
		h[l], h[(l-1)/2] = h[(l-1)/2], h[l]
		l = (l - 1) / 2
	}
}

func down(h Heap) {
	e, start := len(h)-1, 0
	flag := true

	for flag == true {
		if 2*start+1 <= e && h[start] > h[2*start+1] && 2*start+2 <= e && h[start] > h[2*start+2] {
			//fmt.Println(h[start], h[2*start+1], h[2*start+2])
			i := map[bool]int{true: 2*start + 2, false: 2*start + 1}[h[2*start+1] > h[2*start+2]]
			h[start], h[i] = h[i], h[start]
			start = i
		} else if 2*start+1 <= e && h[start] > h[2*start+1] {
			h[start], h[2*start+1] = h[2*start+1], h[start]
			start = 2*start + 1
		} else if 2*start+2 <= e && h[start] > h[2*start+2] {
			h[start], h[2*start+2] = h[2*start+2], h[start]
			start = 2*start + 2
		} else {
			flag = false
		}
	}
}

func (h *Heap) Push(v int) {
	*h = append(*h, v)
	up(*h)
}

func (h *Heap) Pop() int {
	elem := (*h)[len(*h)-1]
	(*h)[0], (*h)[len(*h)-1] = (*h)[len(*h)-1], (*h)[0]

	*h = (*h)[0 : len(*h)-1]
	down(*h)

	return elem
}

func (h Heap) Trace() {
	fmt.Println(h)
}
