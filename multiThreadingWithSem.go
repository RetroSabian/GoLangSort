package main

import (
	"sync"
)

func MultiMergeSortWithSem(data []int64, sem chan struct{}) []int64 {
	if len(data) < 2 {
		return data
	}

	middle := len(data) / 2

	wg := sync.WaitGroup{}
	wg.Add(2)

	var ldata []int64
	var rdata []int64

	select {
	case sem <- struct{}{}:
		go func() {
			ldata = MultiMergeSortWithSem(data[:middle], sem)
			<-sem
			wg.Done()
		}()
	default:
		ldata = SingleMergeSort(data[:middle])
		wg.Done()
	}

	select {
	case sem <- struct{}{}:
		go func() {
			rdata = MultiMergeSortWithSem(data[middle:], sem)
			<-sem
			wg.Done()
		}()
	default:
		rdata = SingleMergeSort(data[middle:])
		wg.Done()
	}

	wg.Wait()
	return Merge(ldata, rdata)
}

func RunMultiMergesortWithSem(data []int64) []int64 {
	// amount of semephores that I could play around with
	sem := make(chan struct{}, 128)
	return MultiMergeSortWithSem(data, sem)
}
