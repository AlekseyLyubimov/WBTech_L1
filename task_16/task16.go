package main

func main() {
	array := []int{0, 99, 3, 1, 756, 765, 123, 2, 999, 61, 33, 11, 975, 7, 1, 0, 0, 542}

	quickSort(array, 0, len(array)-1)

	println("Sorted array: ")
	for _, val := range array {
		println(val)
	}
}

func step(arr []int, start, end int) (int) {

	pivot_val := arr[end]
	left_ptr := start

	for right_ptr := start; right_ptr < end; right_ptr++ {
		if arr[right_ptr] < pivot_val {
			arr[left_ptr], arr[right_ptr] = arr[right_ptr], arr[left_ptr]
			left_ptr++
		}
	}

	arr[left_ptr], arr[end] = arr[end], arr[left_ptr]
	return left_ptr
}



func quickSort(arr []int, low, high int) {

	if low < high {
		p := step(arr, low, high)
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}
