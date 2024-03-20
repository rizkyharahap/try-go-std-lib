package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Rizki", "Mahfuddin", "Harahap", "Roy"}
	values := []int{100, 95, 80, 90}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Harahap"))
	fmt.Println(slices.Index(names, "Roy"))
}
