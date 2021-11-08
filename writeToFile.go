package main

import (
	"fmt"
	"os"
	"strings"
)

func WriteToFile(fileName string, values []int64) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	s := strings.Trim(strings.Replace(fmt.Sprint(values), " ", "\n", -1), "[]")
	fmt.Fprint(f, s)
	return nil
}
