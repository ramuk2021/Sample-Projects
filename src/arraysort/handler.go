package arraysort

type sortHandler struct{}

type SortHandler interface {
	BubbleSort([]int)
	QuickSort([]int)
}

func NewSortHandler() SortHandler {
	return &sortHandler{}
}
