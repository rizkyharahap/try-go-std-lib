package main

import "fmt"

func main() {
	firstName := "Rizki"
	lastName := "Harahap"

	fmt.Println("Hello '", firstName, lastName, "'")

	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}
