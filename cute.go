package cute

import (
	"fmt"
	"os"
	"strings"
)

// CuteColor defines a color
type CuteColor string

// Special defines a color
type Special string

// RenderColor defines a color struct,contains background color and text color
type RenderColor struct {
	Background CuteColor
	TextColor  CuteColor
	Special    Special
}

func (r *RenderColor) String() string {
	return fmt.Sprintf("\033[%s;%s;%sm", r.Background, r.Special, r.TextColor)
}

// defines text color, background color, special characters control
const (
	NoColor CuteColor = "0" // no color, background or text will be set default color

	TextBlack  CuteColor = "30" // text black
	TextRed    CuteColor = "31" // text red
	TextGreen  CuteColor = "32" // text green
	TextYellow CuteColor = "33" // text yellow
	TextBlue   CuteColor = "34" // text blue
	TextPurple CuteColor = "35" // text purple
	TextCyan   CuteColor = "36" // text cyan
	TextWhite  CuteColor = "37" // text white

	BackgroundBlack  CuteColor = "40" // background black
	BackgroundRed    CuteColor = "41" // background red
	BackgroundGreen  CuteColor = "42" // background green
	BackgroundYellow CuteColor = "43" // background yellow
	BackgroundBlue   CuteColor = "44" // background blue
	BackgroundPurple CuteColor = "45" // background purple
	BackgroundCyan   CuteColor = "46" // background cyan
	BackgroundWhite  CuteColor = "47" // background white

	NoSpecial             = "0" // no special
	SpecialHighlight      = "1" // special highlight
	SpecialFade           = "2" // special fade
	SpecialItalic         = "3" // special italic
	SpecialUnderline      = "4" // special underline
	SpecialFlicker        = "5" // special flicker
	SpecialFlickerQuickly = "6" // special flicker quality
	SpecialReverse        = "7" // special reverse
	SpecialBlank          = "8" // special blank
	SpecialDelete         = "9" // special delete
)

// default color
var (
	titleColor *RenderColor = &RenderColor{
		Background: NoColor,
		TextColor:  TextYellow,
		Special:    NoSpecial,
	}

	messageColor *RenderColor = &RenderColor{
		Background: NoColor,
		TextColor:  TextPurple,
		Special:    NoSpecial,
	}

	noColor *RenderColor = &RenderColor{}
)

// SetTitleColor set the title color
func SetTitleColor(c *RenderColor) {
	titleColor = c
}

// SetMessageColor set the message color
func SetMessageColor(c *RenderColor) {
	messageColor = c
}

// titleDraw local drawing title
func titleDraw(title string) (box string) {
	// box elements
	topright := "╮"
	topleft := "╭"
	w := "─"
	h := "│"
	bottomright := "╯"
	bottomleft := "╰"

	// draw line
	line := strings.Repeat(w, len(title)+2)

	// print title box
	box = fmt.Sprintf("%v%v%v%v%v\n", titleColor, topleft, line, topright, noColor)
	box += fmt.Sprintf("%v%v %v %v%v\n", titleColor, h, title, h, noColor)
	box += fmt.Sprintf("%v%v%v%v%v", titleColor, bottomleft, line, bottomright, noColor)
	return
}

// messageDraw local draw message
func messageDraw(message string) (msg string) {
	msg = fmt.Sprintf("%v ▶ %s%v", messageColor, message, noColor)
	return
}

// Println println
func Println(title string, messages ...string) {
	fmt.Println(titleDraw(title))

	for _, m := range messages {
		fmt.Println(messageDraw(m))
	}
}

// Printf print format
func Printf(title string, message string, params ...interface{}) {
	fmt.Println(titleDraw(title))

	fmt.Printf(messageDraw(message), params...)
}

// Check a cute panic like
func Check(title string, err error) {
	if err != nil {
		// print title
		fmt.Printf("%v", TextYellow)
		fmt.Println(titleDraw(title))

		// print error message
		fmt.Printf("%v", TextRed)
		fmt.Println("▶", err.Error())
		fmt.Printf("%v", noColor)
		os.Exit(1)
	}
}
