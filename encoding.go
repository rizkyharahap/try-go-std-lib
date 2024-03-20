package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	value := "Rizki Mahfuddin Harahap"

	encode := base64.StdEncoding.EncodeToString([]byte(value))

	fmt.Println(encode)

	decoded, err := base64.StdEncoding.DecodeString(encode)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(string(decoded))
	}

}
