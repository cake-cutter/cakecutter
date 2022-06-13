package utils

import (
	"fmt"
	"os"
)

func Check(err error) {
	if err != nil {
		fmt.Println(Colorize("red", "Error: "+err.Error()))
		os.Exit(1)
	}
}
