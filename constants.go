package color

const (
	// Reset all colors
	Reset       Color = 0
	resetEscape       = "\x1B[0m"
)

// Foreground colors 30-37
const (
	FgBlack Color = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
	FgDefault Color = 39
)

// Additional foreground colors 90-97
const (
	FgGrey Color = iota + 90
	FgBrightRed
	FgBrightGreen
	FgBrightYellow
	FgBrightBlue
	FgBrightMagenta
	FgBrightCyan
	FgBrightWhite
)

// Background colors 40-47
const (
	BgBlack Color = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
	BgDefault Color = 49
)

// Additional background colors 100-107
const (
	BgGrey Color = iota + 100
	BgBrightRed
	BgBrightGreen
	BgBrightYellow
	BgBrightBlue
	BgBrightMagenta
	BgBrightCyan
	BgBrightWhite
)

// Foreground color set alias
var Foreground = colorSet{
	Black:   FgBlack,
	Red:     FgRed,
	Green:   FgGreen,
	Yellow:  FgYellow,
	Blue:    FgBlue,
	Magenta: FgMagenta,
	Cyan:    FgCyan,
	White:   FgWhite,
	Default: FgDefault,

	Bright: &colorSet{
		Black:   FgGrey,
		Red:     FgBrightRed,
		Green:   FgBrightGreen,
		Yellow:  FgBrightYellow,
		Blue:    FgBrightBlue,
		Magenta: FgBrightMagenta,
		Cyan:    FgBrightCyan,
		White:   FgBrightWhite,
		Default: FgDefault,
	},
}

// Background color set alias
var Background = colorSet{
	Black:   BgBlack,
	Red:     BgRed,
	Green:   BgGreen,
	Yellow:  BgYellow,
	Blue:    BgBlue,
	Magenta: BgMagenta,
	Cyan:    BgCyan,
	White:   BgWhite,
	Default: BgDefault,

	Bright: &colorSet{
		Black:   BgGrey,
		Red:     BgBrightRed,
		Green:   BgBrightGreen,
		Yellow:  BgBrightYellow,
		Blue:    BgBrightBlue,
		Magenta: BgBrightMagenta,
		Cyan:    BgBrightCyan,
		White:   BgBrightWhite,
		Default: BgDefault,
	},
}
