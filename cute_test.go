package cute

import (
	"testing"
)

// Test equality of CuteColor type.
func TestCuteColorType(t *testing.T) {
	redColor := CuteColor("31")
	if TextRed != redColor {
		t.Fail()
	}
}

func TestPrintln(t *testing.T) {
	/*SetTitleColor(&RenderColor{
		Background: NoColor,
		TextColor:  TextPurple,
		Special:    SpecialItalic,
	})
	SetMessageColor(&RenderColor{
		Background: NoColor,
		TextColor:  TextYellow,
		Special:    SpecialItalic,
	})*/

	Println("Foo", "Bar", "hello")
	Printf("Foo", "Bar %d", 1)
}
