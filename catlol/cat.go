package catlol

import (
	"fmt"
	"math"
)

func rgb(i int) (int, int, int) {
	var f = 0.1
	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}

func MakeRainbowText(input string) string {
	var result string
	j := 0
	for _, r := range input {
		if r == '\n' {
			result += "\n"
			continue
		}
		red, green, blue := rgb(j)
		result += fmt.Sprintf("\033[38;2;%d;%d;%dm%c\033[0m", red, green, blue, r)
		j++
	}
	return result
}
