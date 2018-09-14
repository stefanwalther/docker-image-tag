package main

import (
	"fmt"
	"os"
)

func failIfErrorIsNotNil(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
