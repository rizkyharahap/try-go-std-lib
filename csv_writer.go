package main

import (
	"encoding/csv"
	"os"
)

func main() {
	var writer *csv.Writer = csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"Rizki", "Mahdfuddin", "Harahap"})
	_ = writer.Write([]string{"Roy", "Martin", "Ramones"})
	_ = writer.Write([]string{"Gading", "Martin", "Lupa"})

	writer.Flush()
}
