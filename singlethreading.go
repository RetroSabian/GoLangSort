package main

func SingleMergeSort(data []int64) []int64 {
	if len(data) < 2 {
		return data
	}
	middle := len(data) / 2
	return Merge(SingleMergeSort(data[:middle]), SingleMergeSort(data[middle:]))
}
