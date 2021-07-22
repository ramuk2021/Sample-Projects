package arraysort

func (h *sortHandler) QuickSort(a []int) {
	quicksort(a, 0, len(a)-1)
}

func quicksort(a []int, left, right int) {
	if left < right {
		pivot := partition(a, left, right)
		quicksort(a, left, pivot-1)
		quicksort(a, pivot+1, right)
	}
}

func partition(a []int, left, right int) int {
	pivot := a[right]
	i := left - 1
	for j := left; j < right; j++ {
		if a[j] <= pivot {
			i = i + 1
			swap(a, i, j)
		}
	}
	swap(a, i+1, right)
	return i + 1
}
