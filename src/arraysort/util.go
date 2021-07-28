package arraysort

func swap(a []int, i, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}
