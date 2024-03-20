package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "Rizki,Mahfuddin,Harahap\n" +
		"Roy,Martin,Ramones\n" +
		"Gading,Martin,Lupa"

	var reader *csv.Reader = csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}
