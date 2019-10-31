// Package color is simple colorization using ansi escape codes. It uses
// the standard Print/Fprint/Sprint(f|ln) interface. For windows support
// consider using go-colorable.
package color

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var (
	// Disable output of ansi color codes
	// if this is set to true.
	Disable = false
	// Writer is the default writer for Print statements.
	Writer io.Writer = os.Stdout
)

type (
	// Color represents a specific color
	Color int
	// Colors is a bunch of different colors combined
	Colors []Color

	colorSet struct {
		Black, Red, Green, Yellow, Blue, Magenta, Cyan, White, Default Color

		Bright *colorSet
	}
)

var (
	rgxCodes = regexp.MustCompile("\x1B" + `\[\d+(?:;\d+)?m`)
)

// Mix colors together
func Mix(colors ...Color) Colors {
	c := make([]Color, len(colors))
	copy(c, colors)
	return Colors(c)
}

// Println the arguments and append a newline
func (c Color) Println(args ...interface{}) error {
	return c.Fprintln(Writer, args...)
}

// Print the arguments
func (c Color) Print(args ...interface{}) error {
	return c.Fprint(Writer, args...)
}

// Printf prints a formatted string
func (c Color) Printf(format string, args ...interface{}) error {
	return c.Fprintf(Writer, format, args...)
}

// Fprintln the arguments to the writer and append a newline
func (c Color) Fprintln(w io.Writer, args ...interface{}) error {
	_, err := io.WriteString(w, c.Sprintln(args...))
	return err
}

// Fprint the arguments to the writer
func (c Color) Fprint(w io.Writer, args ...interface{}) error {
	_, err := io.WriteString(w, c.Sprint(args...))
	return err
}

// Fprintf formats string and writes it to w
func (c Color) Fprintf(w io.Writer, format string, args ...interface{}) error {
	_, err := io.WriteString(w, c.Sprintf(format, args...))
	return err
}

// Sprintln the arguments and append a newline
func (c Color) Sprintln(args ...interface{}) string {
	return surround(fmt.Sprintln(args...), c)
}

// Sprint the arguments
func (c Color) Sprint(args ...interface{}) string {
	return surround(fmt.Sprint(args...), c)
}

// Sprintf prints a formatted string
func (c Color) Sprintf(format string, args ...interface{}) string {
	return surround(fmt.Sprintf(format, args...), c)
}

// Println the arguments and append a newline
func (c Colors) Println(args ...interface{}) error {
	return c.Fprintln(Writer, args...)
}

// Print the arguments
func (c Colors) Print(args ...interface{}) error {
	return c.Fprint(Writer, args...)
}

// Printf prints a formatted string
func (c Colors) Printf(format string, args ...interface{}) error {
	return c.Fprintf(Writer, format, args...)
}

// Fprintln the arguments to the writer and append a newline
func (c Colors) Fprintln(w io.Writer, args ...interface{}) error {
	_, err := io.WriteString(w, c.Sprintln(args...))
	return err
}

// Fprint the arguments to the writer
func (c Colors) Fprint(w io.Writer, args ...interface{}) error {
	_, err := io.WriteString(w, c.Sprint(args...))
	return err
}

// Fprintf formats string and writes it to w
func (c Colors) Fprintf(w io.Writer, format string, args ...interface{}) error {
	_, err := io.WriteString(w, c.Sprintf(format, args...))
	return err
}

// Sprintln the arguments and append a newline
func (c Colors) Sprintln(args ...interface{}) string {
	return surround(fmt.Sprintln(args...), c...)
}

// Sprint the arguments
func (c Colors) Sprint(args ...interface{}) string {
	return surround(fmt.Sprint(args...), c...)
}

// Sprintf prints a formatted string
func (c Colors) Sprintf(format string, args ...interface{}) string {
	return surround(fmt.Sprintf(format, args...), c...)
}

// Clean all color codes from a string
func Clean(s string) string {
	return rgxCodes.ReplaceAllString(s, "")
}

func surround(s string, colors ...Color) string {
	if Disable {
		return s
	}

	// Clear any resets that we have inside this string
	// as they will mangle the colors we're about to surround with
	newColor := ansiEscape(colors...)
	s = strings.ReplaceAll(s, resetEscape, resetEscape+newColor)
	return fmt.Sprintf("%s%s%s", newColor, s, resetEscape)
}

func ansiEscape(colors ...Color) string {
	var foreground Color
	var background Color

	for _, c := range colors {
		if c == 0 {
			return resetEscape
		}

		if (c >= FgBlack && c <= FgWhite) || (c >= FgGrey && c <= FgBrightWhite) {
			foreground = c
			continue
		}

		background = c
	}

	switch {
	case foreground != 0 && background != 0:
		return fmt.Sprintf("\x1B[%d;%dm", foreground, background)
	case foreground != 0:
		return fmt.Sprintf("\x1B[%dm", foreground)
	case background != 0:
		return fmt.Sprintf("\x1B[%dm", background)
	default:
		return ""
	}
}
