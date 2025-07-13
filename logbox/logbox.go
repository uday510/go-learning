package logbox

import (
	"fmt"
	"strings"
)

func PrintBox(lines []string) {
	withPadding := make([]string, len(lines))
	for i, line := range lines {
		withPadding[i] = " " + line
	}

	width := maxLineLength(withPadding) + 1
	fmt.Println("╔" + strings.Repeat("═", width) + "╗")
	for _, line := range withPadding {
		fmt.Printf("║%-*s║\n", width, line)
	}
	fmt.Println("╚" + strings.Repeat("═", width) + "╝")
}

func maxLineLength(lines []string) int {
	max := 0
	for _, line := range lines {
		if len(line) > max {
			max = len(line)
		}
	}
	return max
}
