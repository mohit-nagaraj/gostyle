package catlol

import (
	"fmt"
	"math"
)

// Calculates the RGB color values based on the input index.
//
// i: The input index for the color calculation.
func rgb(i int) (int, int, int) {
	var f = 0.1
	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}

// Converts the input text to rainbow-colored text using RGB values.
//
// input: The input text to be colored.
func MakeRainbowText(input string) string {
	var result string
	j := 0
	for _, r := range input {
		if r == '\n' {
			result += "\n"
			continue
		}
		red, green, blue := rgb(j)
		//fmt.Sprintf("\033[38;2;%d;%d;%dm%c\033[0m", red, green, blue, r):
		//This line formats the character r with RGB color values using ANSI escape codes.
		//The escape code \033[38;2;R;G;Bm sets the text color to the RGB values (R, G, B).
		//The escape code \033[0m resets the text color to default.
		result += fmt.Sprintf("\033[38;2;%d;%d;%dm%c\033[0m", red, green, blue, r)
		j++
	}
	return result
}
