package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Rizki Mahfuddin Harahap", "Rizki"))
	fmt.Println(strings.Split("Rizki Mahfuddin Harahap", " "))
	fmt.Println(strings.ToLower("Rizki Mahfuddin Harahap"))
	fmt.Println(strings.ToUpper("Rizki Mahfuddin Harahap"))
	fmt.Println(strings.Trim(" Rizki Mahfuddin Harahap ", " "))
	fmt.Println(strings.ReplaceAll("Rizki Mahfuddin Harahap", "Mahfuddin ", ""))
}
