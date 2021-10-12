package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findIan(in string) {
	in = strings.ToLower(in)
	in = strings.ReplaceAll(in, " ", "")

	if strings.HasPrefix(in, "i") && strings.Index(in, "a") > 0 && strings.HasSuffix(in, "n") {
		fmt.Println("Found!!")
	} else {
		fmt.Println("Not Found!")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Insert your string:")
	scanner.Scan()
	input := scanner.Text()
	findIan(input)
}
