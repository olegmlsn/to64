package main

import (
	"fmt"
	"os"

	"github.com/olegmlsn/to64/packages/encoder"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Oops! You need to provide a file path.")
		os.Exit(1)
	}

	result, err := encoder.EncodeFileToBase64(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(result)
}
