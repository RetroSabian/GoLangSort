package main

func MultiMergeSort(data []int64, res chan []int64) {
	if len(data) < 2 {
		res <- data
		return
	}

	leftChan := make(chan []int64)
	rightChan := make(chan []int64)
	middle := len(data) / 2

	go MultiMergeSort(data[:middle], leftChan)
	go MultiMergeSort(data[middle:], rightChan)
	ldata := <-leftChan
	rdata := <-rightChan

	close(leftChan)
	close(rightChan)
	res <- Merge(ldata, rdata)
}

func RunMultiMergeSort(data []int64) (multiResult []int64) {
	res := make(chan []int64)
	go MultiMergeSort(data, res)
	multiResult = <-res
	return
}
