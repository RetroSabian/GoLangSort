package main

import (
	"fmt"
	"time"
)

func main() {
	// CreateInputFile()
	fmt.Println("Start of file read")
	start := time.Now()
	fileName := "input.txt"
	s := ReadFromFile(fileName)
	startSort := time.Now()
	multiResultWithSem := RunMultiMergesortWithSem(s)
	fmt.Println(time.Since(startSort))
	outputFile := "output.txt"
	WriteToFile(outputFile, multiResultWithSem)
	fmt.Println("End of file write")
	fmt.Println(time.Since(start))

}
