package utils

import (
	"time"

	"github.com/briandowns/spinner"
)

var spin = []string{
	" [                   ] ",
	" [█                  ] ",
	" [██                 ] ",
	" [███                ] ",
	" [████               ] ",
	" [█████              ] ",
	" [██████             ] ",
	" [███████            ] ",
	" [████████           ] ",
	" [█████████          ] ",
	" [██████████         ] ",
	" [███████████        ] ",
	" [████████████       ] ",
	" [█████████████      ] ",
	" [██████████████     ] ",
	" [███████████████    ] ",
	" [████████████████   ] ",
	" [█████████████████  ] ",
	" [██████████████████ ] ",
	" [███████████████████] ",
}

func MakeItSpin(f func(), suffix string) {

	s := spinner.New(spin, 50*time.Millisecond)
	s.Suffix = Colorize("blue", suffix)
	s.Start()

	f()

	time.Sleep(time.Second)

	s.Stop()

}
