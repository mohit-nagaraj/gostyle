package cowsay

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var Cow = ` 
         \  ^__^
          \ (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
`

// Builds a text balloon for the fortune, ensuring proper alignment and padding.
//
// fortune: The text of the fortune to be placed in the balloon.
func BuildBalloon(fortune string) string {
	lines := strings.Split(fortune, "\n") // Convert the fortune into a slice of lines based on the newline character.
	lines = tabsToSpaces(lines)           // Convert tabs to spaces.
	maxwidth := calculateMaxWidth(lines)  // Calculate the maximum width of the lines.
	messages := normalizeStringsLength(lines, maxwidth)
	balloon := buildBalloon(messages, maxwidth)
	return balloon
}

func buildBalloon(lines []string, maxwidth int) string {
	var borders []string
	count := len(lines)
	var ret []string

	borders = []string{"/", "\\", "\\", "/", "|", "<", ">"}

	top := " " + strings.Repeat("_", maxwidth+2)
	bottom := " " + strings.Repeat("-", maxwidth+2)

	ret = append(ret, top)
	if count == 1 {
		s := fmt.Sprintf("%s %s %s", borders[5], lines[0], borders[6])
		ret = append(ret, s)
	} else {
		s := fmt.Sprintf(`%s %s %s`, borders[0], lines[0], borders[1])
		ret = append(ret, s)
		i := 1
		for ; i < count-1; i++ {
			s = fmt.Sprintf(`%s %s %s`, borders[4], lines[i], borders[4])
			ret = append(ret, s)
		}
		s = fmt.Sprintf(`%s %s %s`, borders[2], lines[i], borders[3])
		ret = append(ret, s)
	}

	ret = append(ret, bottom)
	return strings.Join(ret, "\n")
}

func tabsToSpaces(lines []string) []string {
	var ret []string
	for _, l := range lines {
		l = strings.Replace(l, "\t", "    ", -1)
		ret = append(ret, l)
	}
	return ret
}

// Calculates the maximum width of the lines in the text.
//
// lines: The lines of text to measure.
func calculateMaxWidth(lines []string) int {
	w := 0
	for _, l := range lines {
		//check the maximum width of the line
		len := utf8.RuneCountInString(l)
		if len > w {
			w = len
		}
	}

	return w
}

// Normalizes the length of all lines by appending spaces so that
// all lines have the same width.
//
// lines: The lines of text to normalize.
// maxwidth: The maximum width to which the lines should be padded.
func normalizeStringsLength(lines []string, maxwidth int) []string {
	var ret []string
	for _, l := range lines {
		s := l + strings.Repeat(" ", maxwidth-utf8.RuneCountInString(l))
		ret = append(ret, s)
	}
	return ret
}
