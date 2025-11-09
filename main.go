package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "read error:", err)
		return
	}
	// Normalize windows line endings
	b = bytes.ReplaceAll(b, []byte{'\r', '\n'}, []byte{'\n'})
	input := string(b)

	// Drop final trailing newline if it created an empty line at the end
	lines := strings.Split(input, "\n")
	// If last element is empty (trailing newline) drop it.
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	// If nothing at all
	if len(lines) == 0 {
		fmt.Print("Not a quad function\n")
		return
	}

	// Convert lines to rune slices and ensure all same width
	h := len(lines)
	width := -1
	runes := make([][]rune, h)
	for i, line := range lines {
		// handle possible \r leftover just in case
		line = strings.TrimRight(line, "\r")
		// decode as runes
		rs := []rune(line)
		runes[i] = rs
		if width == -1 {
			width = len(rs)
		} else if width != len(rs) {
			fmt.Print("Not a quad function\n")
			return
		}
	}

	// width might be zero (empty lines) -> not rectangle
	if width <= 0 || h <= 0 {
		fmt.Print("Not a quad function\n")
		return
	}

	// helper to get rune at (r,c) where r and c are 1-based
	getAt := func(r, c int) rune {
		return runes[r-1][c-1]
	}

	matches := []string{}
	if matchesQuadA(getAt, width, h) {
		matches = append(matches, fmt.Sprintf("[quadA] [%d] [%d]", width, h))
	}
	if matchesQuadB(getAt, width, h) {
		matches = append(matches, fmt.Sprintf("[quadB] [%d] [%d]", width, h))
	}
	if matchesQuadC(getAt, width, h) {
		matches = append(matches, fmt.Sprintf("[quadC] [%d] [%d]", width, h))
	}
	if matchesQuadD(getAt, width, h) {
		matches = append(matches, fmt.Sprintf("[quadD] [%d] [%d]", width, h))
	}
	if matchesQuadE(getAt, width, h) {
		matches = append(matches, fmt.Sprintf("[quadE] [%d] [%d]", width, h))
	}

	if len(matches) == 0 {
		fmt.Print("Not a quad function\n")
		return
	}

	// sort alphabetical
	sort.Strings(matches)
	fmt.Print(strings.Join(matches, " || ") + "\n")
}

// expected rune functions for each quad type:

func matchesQuadA(getAt func(int, int) rune, w, h int) bool {
	for r := 1; r <= h; r++ {
		for c := 1; c <= w; c++ {
			var exp rune
			if (r == 1 && c == 1) || (r == 1 && c == w) || (r == h && c == 1) || (r == h && c == w) {
				exp = 'o'
			} else if r == 1 || r == h {
				exp = '-'
			} else if c == 1 || c == w {
				exp = '|'
			} else {
				exp = ' '
			}
			if getAt(r, c) != exp {
				return false
			}
		}
	}
	return true
}

func matchesQuadB(getAt func(int, int) rune, w, h int) bool {
	for r := 1; r <= h; r++ {
		for c := 1; c <= w; c++ {
			var exp rune
			if r == 1 && c == 1 {
				exp = '/'
			} else if r == 1 && c == w {
				exp = '\\'
			} else if r == h && c == 1 {
				exp = '\\'
			} else if r == h && c == w {
				exp = '/'
			} else if r == 1 || r == h || c == 1 || c == w {
				exp = '*'
			} else {
				exp = ' '
			}
			if getAt(r, c) != exp {
				return false
			}
		}
	}
	return true
}

func matchesQuadC(getAt func(int, int) rune, w, h int) bool {
	for r := 1; r <= h; r++ {
		for c := 1; c <= w; c++ {
			var exp rune
			if r == 1 && c == 1 {
				exp = 'A'
			} else if r == 1 && c == w {
				exp = 'A'
			} else if r == h && c == 1 {
				exp = 'C'
			} else if r == h && c == w {
				exp = 'C'
			} else if r == 1 || r == h || c == 1 || c == w {
				exp = 'B'
			} else {
				exp = ' '
			}
			if getAt(r, c) != exp {
				return false
			}
		}
	}
	return true
}

func matchesQuadD(getAt func(int, int) rune, w, h int) bool {
	for r := 1; r <= h; r++ {
		for c := 1; c <= w; c++ {
			var exp rune
			if r == 1 && c == 1 {
				exp = 'A'
			} else if r == 1 && c == w {
				exp = 'C'
			} else if r == h && c == 1 {
				exp = 'A'
			} else if r == h && c == w {
				exp = 'C'
			} else if r == 1 || r == h || c == 1 || c == w {
				exp = 'B'
			} else {
				exp = ' '
			}
			if getAt(r, c) != exp {
				return false
			}
		}
	}
	return true
}

func matchesQuadE(getAt func(int, int) rune, w, h int) bool {
	for r := 1; r <= h; r++ {
		for c := 1; c <= w; c++ {
			var exp rune
			if r == 1 && c == 1 {
				exp = 'A'
			} else if r == 1 && c == w {
				exp = 'C'
			} else if r == h && c == 1 {
				exp = 'C'
			} else if r == h && c == w {
				exp = 'A'
			} else if r == 1 || r == h || c == 1 || c == w {
				exp = 'B'
			} else {
				exp = ' '
			}
			if getAt(r, c) != exp {
				return false
			}
		}
	}
	return true
}
