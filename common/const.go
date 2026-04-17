package common

import "strings"

const lineLength int = 40

// access within this Package
func PrintLine(line string) {
	println(strings.Repeat(line, lineLength))
}
