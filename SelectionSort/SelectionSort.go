package SelectionSort

func findSmallestIndex(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}

func Sort(arr []int) []int {
	size := len(arr)
	newArr := make([]int, size)
	for i := 0; i < size; i++ {
		smallest := findSmallestIndex(arr)
		newArr[i] = arr[smallest]
		arr = append(arr[:smallest], arr[smallest+1:]...)
	}
	return newArr
}
