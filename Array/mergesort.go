package main

import "log"

func MergeSort(arr []int) {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}

	mergeSort(arr, 0, arrLen-1)
}

func mergeSort(arr []int, start int, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	tmp := make([]int, end-start+1)

	i := start
	j := mid + 1
	k := 0
	for ; i <= mid && j <= end; k++ {
		if arr[i] < arr[j] {
			tmp[k] = arr[i]
			i++
		} else {
			tmp[k] = arr[j]
			j++
		}
	}

	for ; i <= mid; i++ {
		tmp[k] = arr[i]
		k++
	}

	for ; j <= end; j++ {
		tmp[k] = arr[j]
		k++
	}

	copy(arr[start:end+1], tmp)
}

func main() {
	test := []int{0, 1, 2, 3, 4, 5, 6}

	MergeSort(test)

	log.Println(test)
}
