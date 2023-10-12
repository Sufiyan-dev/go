package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//
	name := ""
	address := ""

	// mapping
	info := make(map[string]string)

	// taking user inputs
	fmt.Print("Enter your name ")
	fmt.Scan(&name)

	fmt.Print("Enter your address ")
	fmt.Scan(&address)

	// saving user state
	info["name"] = name
	info["address"] = address

	// converting mapping to json object
	data, err := json.Marshal(info)

	// error handling
	if err != nil {
		fmt.Println("Error converting mapping pair to json: ", err)
		return
	}

	// Printing tha json data
	fmt.Println(string(data))

}
