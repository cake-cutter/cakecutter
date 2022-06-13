package main

import (
	"github.com/cake-cutter/cli/cmd"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cmd.Execute()
}
