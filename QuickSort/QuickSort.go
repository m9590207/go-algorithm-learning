package QuickSort

func QuickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	pivot := array[0]
	less := []int{}
	greater := []int{}
	for _, value := range array[1:] {
		if pivot > value {
			less = append(less, value)
		} else {
			greater = append(greater, value)
		}
	}
	less = append(QuickSort(less), pivot)
	greater = QuickSort(greater)
	return append(less, greater...)
}
