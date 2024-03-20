package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	// With temporari
	// temp := s[i]
	// s[i] = s[j]
	// s[j] = temp

	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{"Rizki", 30},
		{"Mahfuddin", 35},
		{"Harahap", 25},
		{"Hap", 20},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
