package arraysort

func (h *sortHandler) BubbleSort(a []int) {
	bubbleSort(a)
}

func bubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[i] < a[j] {
				swap(a, i, j)
			}
		}
	}
}
