package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func ReadFromFile(fileName string) []int64 {
	response := make([]int64, 0)
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("cannot able to read the file", err)
	}

	defer file.Close() //close after checking err

	var wg sync.WaitGroup
	linesPool := sync.Pool{New: func() interface{} {
		lines := make([]byte, 250*1024)
		return lines
	}}

	intPool := sync.Pool{New: func() interface{} {
		lines := ""
		return lines
	}}

	r := bufio.NewReader(file)
	for {
		buf := linesPool.Get().([]byte)

		n, err := r.Read(buf)
		buf = buf[:n]

		if n == 0 {
			if err != nil {
				fmt.Println(err)
				break
			}
			if err == io.EOF {
				break
			}
		}

		nextUntillNewline, err := r.ReadBytes('\n')

		if err != io.EOF {
			buf = append(buf, nextUntillNewline...)
		}

		wg.Add(1)
		go func() {
			res := ProcessChunk(buf, &linesPool, &intPool)
			response = append(response, res...)
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println(len(response))
	return response
}

func ProcessChunk(chunk []byte, linesPool *sync.Pool, intPool *sync.Pool) []int64 {
	response := make([]int64, 0)

	var wg2 sync.WaitGroup

	logs := intPool.Get().(string)
	logs = string(chunk)

	linesPool.Put(chunk)

	logsSlice := strings.Split(logs, "\n")

	intPool.Put(logs)

	chunkSize := 350
	n := len(logsSlice)
	noOfThread := n / chunkSize

	if n%chunkSize != 0 {
		noOfThread++
	}
	for i := 0; i < (noOfThread); i++ {

		wg2.Add(1)
		go func(s int, e int) {
			defer wg2.Done() //to avaoid deadlocks
			for i := s; i < e; i++ {
				text := logsSlice[i]
				if len(text) == 0 {
					continue
				}
				value, err := strconv.ParseInt(text, 0, 64)
				if err != nil {
					fmt.Println(err)
					break
				}
				if value == 0 {
					fmt.Println("how")
				}
				response = append(response, value)
			}

		}(i*chunkSize, int(math.Min(float64((i+1)*chunkSize), float64(len(logsSlice)))))
		// }(i*chunkSize, int(float64(len(logsSlice))))
	}

	wg2.Wait()
	logsSlice = nil
	return response
}
