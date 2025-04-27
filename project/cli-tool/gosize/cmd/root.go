package cmd

import (
	"fmt"
	"gosize/internal/size"
	"os"
)

func Execute() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gosize <path>")
		return
	}

	path := os.Args[1]
	err := size.ListSizes(path)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
