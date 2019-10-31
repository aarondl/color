package color

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"testing"
)

const testFile = "testdata/expect.golden"

var (
	flagGolden = flag.Bool("golden", false, "write golden files")
)

func TestColors(t *testing.T) {
	t.Parallel()

	writer := new(bytes.Buffer)

	fmt.Fprintln(writer,
		FgBlack.Sprint("fg  "),
		FgRed.Sprint("fg  "),
		FgGreen.Sprint("fg  "),
		FgYellow.Sprint("fg  "),
		FgBlue.Sprint("fg  "),
		FgMagenta.Sprint("fg  "),
		FgCyan.Sprint("fg  "),
		FgWhite.Sprint("fg  "),
	)

	fmt.Fprintln(writer,
		BgBlack.Sprint("bg  "),
		BgRed.Sprint("bg  "),
		BgGreen.Sprint("bg  "),
		BgYellow.Sprint("bg  "),
		BgBlue.Sprint("bg  "),
		BgMagenta.Sprint("bg  "),
		BgCyan.Sprint("bg  "),
		BgWhite.Sprint("bg  "),
	)

	fmt.Fprintln(writer,
		FgGrey.Sprint("fgb "),
		FgBrightRed.Sprint("fgb "),
		FgBrightGreen.Sprint("fgb "),
		FgBrightYellow.Sprint("fgb "),
		FgBrightBlue.Sprint("fgb "),
		FgBrightMagenta.Sprint("fgb "),
		FgBrightCyan.Sprint("fgb "),
		FgBrightWhite.Sprint("fgb "),
	)

	fmt.Fprintln(writer,
		BgGrey.Sprint("bgb "),
		BgBrightRed.Sprint("bgb "),
		BgBrightGreen.Sprint("bgb "),
		BgBrightYellow.Sprint("bgb "),
		BgBrightBlue.Sprint("bgb "),
		BgBrightMagenta.Sprint("bgb "),
		BgBrightCyan.Sprint("bgb "),
		BgBrightWhite.Sprint("bgb "),
	)

	fmt.Fprintln(writer,
		Mix(FgBlack, BgBlack).Sprint("fgbg"),
		Mix(FgRed, BgRed).Sprint("fgbg"),
		Mix(FgGreen, BgGreen).Sprint("fgbg"),
		Mix(FgYellow, BgYellow).Sprint("fgbg"),
		Mix(FgBlue, BgBlue).Sprint("fgbg"),
		Mix(FgMagenta, BgMagenta).Sprint("fgbg"),
		Mix(FgCyan, BgCyan).Sprint("fgbg"),
		Mix(FgWhite, BgWhite).Sprint("fgbg"),
	)

	// Test defaults
	fmt.Fprintln(writer,
		FgRed.Sprintf("red %s red",
			FgDefault.Sprint("none"),
		),
	)

	// Test additive
	fmt.Fprintln(writer,
		FgBrightRed.Sprintf("ketchup %s ketchup",
			BgBrightYellow.Sprint("mustard"),
		),
	)

	// Test clean
	fmt.Fprintln(writer,
		Clean(
			FgBrightGreen.Sprintf("strip %s colors from %s should work",
				Mix(FgBlue, BgRed).Sprint("all"),
				BgDefault.Sprint("strings"),
			),
		),
	)

	got := writer.Bytes()
	t.Log("\n" + string(got))

	if *flagGolden {
		if err := ioutil.WriteFile(testFile, got, 0664); err != nil {
			t.Fatal(err)
		}
		return
	}

	want, err := ioutil.ReadFile(testFile)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(want, got) {
		t.Errorf("want:\n%s\ngot:\n%s", want, got)
	}
}
