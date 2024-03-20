package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`r(.*[a-z])i`)

	fmt.Println(regex.MatchString("rizki"))
	fmt.Println(regex.MatchString("rIZKi"))
}
