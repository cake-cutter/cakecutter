package main

import (
	"github.com/joho/godotenv"
	"will-change.later/cmd"
)

func main() {
	godotenv.Load()
	cmd.Execute()
}
