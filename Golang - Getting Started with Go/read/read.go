package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FullName struct{ fname, lname string }

func main() {

	var fileNameIN string

	temp := make([]FullName, 0)

	fmt.Println("Enter the name of the file:")
	fmt.Scan(&fileNameIN)

	file, err := os.Open(fileNameIN)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := strings.Split(scanner.Text(), " ")
		var name FullName
		name.fname, name.lname = str[0], str[1]
		temp = append(temp, name)
	}

	file.Close()
	fmt.Println(" ")

	for _, val := range temp {
		fmt.Println(val.fname, val.lname)
	}

}
