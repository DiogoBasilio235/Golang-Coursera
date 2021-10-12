package main

import "fmt"

func BubbleSort(items []int) {
	n := len(items)
	sorted := false

	for !sorted {
		swapped := false

		for i := 0; i < n-1; i++ {

			if items[i] > items[i+1] {
				Swap(items, i)
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}

func Swap(sliceIn []int, index int) {
	temp := sliceIn[index]
	sliceIn[index] = sliceIn[index+1]
	sliceIn[index+1] = temp
}

func main() {
	slice := make([]int, 0)
	var insertedNum int

	for len(slice) < 10 {
		fmt.Println("Enter your number to be sorted:")
		fmt.Scan(&insertedNum)
		slice = append(slice, insertedNum)
		BubbleSort(slice)
		fmt.Println("Sorted list", slice)
	}

}
