package main

import (
	"github.com/cake-cutter/cc/cmd"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cmd.Execute()
}
