package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

func CreateInputFile() {
	numLines := 400000
	inputFile := "input.txt"
	in, err := os.Create(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numLines; i++ {
		fmt.Fprintf(in, "%d\n", rand.Intn(math.MaxInt64))
	}
	in.Close()
}
