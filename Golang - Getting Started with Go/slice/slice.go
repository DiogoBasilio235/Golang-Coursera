package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	slc := make([]int, 3)
	slc = slc[:0]

	var input string

	for true {
		fmt.Println("Insert your number (Press X to quit):")
		fmt.Scan(&input)
		if input == "X" || input == "x" {
			break
		} else {
			val, _ := strconv.Atoi(input)
			slc = append(slc, val)
			sort.Ints(slc)
			fmt.Println(slc)
		}
	}
}
