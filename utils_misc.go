package main

import (
	"fmt"
	"os"
)

func exitIfError(err error) {
	if err != nil {
		fmt.Printf("Fatal error: %v\n", err)
		os.Exit(1)
	}
}
