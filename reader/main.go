package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Error: It is necessary to pass the input file name")
		os.Exit(1)
	}

	filename := args[1]
	fmt.Println("Opening file: ", filename)
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
