package main

import (
	"os"

	preview "github.com/troglobit/preview/src"
)

func main() {
	if err := preview.Execute(); err != nil {
		os.Exit(1)
	}
}
