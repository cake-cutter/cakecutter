package utils

import "fmt"

var colorReset = "\033[0m"
var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorYellow = "\033[33m"
var colorBlue = "\033[96m"
var colorPurple = "\033[35m"
var colorCyan = "\033[36m"
var colorWhite = "\033[37m"
var colorGray = "\033[90m"

var colors map[string]string = map[string]string{
	"red":    colorRed,
	"green":  colorGreen,
	"yellow": colorYellow,
	"blue":   colorBlue,
	"purple": colorPurple,
	"cyan":   colorCyan,
	"white":  colorWhite,
	"gray":   colorGray,
}

func Colorize(color string, text string) string {
	return colors[color] + text + colorReset
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
