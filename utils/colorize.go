package utils

import "fmt"

var (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[96m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
	colorGray   = "\033[90m"
)

var colors = map[string]string{
	"red":    colorRed,
	"green":  colorGreen,
	"yellow": colorYellow,
	"blue":   colorBlue,
	"purple": colorPurple,
	"cyan":   colorCyan,
	"white":  colorWhite,
	"gray":   colorGray,
}

func Colorize(color, text string) string {
	return colors[color] + text + colorReset
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
