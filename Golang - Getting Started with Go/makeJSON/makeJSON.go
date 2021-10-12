package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var nameIN string
	var addrIN string
	var person map[string]string

	person = make(map[string]string)

	fmt.Println("Insert your name:")
	fmt.Scan(&nameIN)
	person["name"] = nameIN

	fmt.Println("Insert your address:")
	fmt.Scan(&addrIN)
	person["address"] = addrIN

	barr, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(barr))

}
