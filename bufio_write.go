package main

import (
	"bufio"
	"os"
)

func main() {
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	_, _ = writer.WriteString("hello world\n")
	_, _ = writer.WriteString("Happy Code\n")

	writer.Flush()
}
