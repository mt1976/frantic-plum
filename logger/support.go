package logger

import (
	"strings"

	"github.com/mt1976/frantic-plum/colours"
)

func fileName(in, name string) string {
	return in + name + ".log"
}

func nameIt(colour, name string) string {
	name = strings.ToUpper(name)
	return colour + sBracket(name) + Reset + " "
}

// sBracket adds square brackets to a string
func sBracket(s string) string {
	return "[" + s + "]"
}

func setColoursNormal() {
	Reset = colours.Reset
	Red = colours.Red
	Green = colours.Green
	Yellow = colours.Yellow
	Blue = colours.Blue
	Magenta = colours.Magenta
	Cyan = colours.Cyan
	Gray = colours.Gray
	White = colours.White
}

func setColoursWindows() {
	Reset = ""
	Red = ""
	Green = ""
	Yellow = ""
	Blue = ""
	Magenta = ""
	Cyan = ""
	Gray = ""
	White = ""
}
