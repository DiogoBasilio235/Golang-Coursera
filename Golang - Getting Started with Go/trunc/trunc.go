package main

import (
	"fmt"
)

func main() {
	var decimal float64

	fmt.Printf("Enter your number?\n")
	fmt.Scan(&decimal)
	fmt.Printf("Your truncated number = %.0f", decimal)

}
