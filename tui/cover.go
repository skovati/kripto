package tui

import (
    "strings"
    "fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const logo = `
 _         _       _
| |       (_)     | |
| | ___ __ _ _ __ | |_ ___
| |/ / '__| | '_ \| __/ _ \
|   <| |  | | |_) | || (_) |
|_|\_\_|  |_| .__/ \__\___/
            | |
            |_|
`

const (
	subtitle   = `kripto - simple cryptocurrency tracker`
	navigation = `ctrl-n: Next screen    ctrl-p: Previous screen    ctrl-c: Exit`
)

// Cover returns the cover page.
func Cover(nextSlide func()) (title string, content tview.Primitive) {
	// What's the size of the logo?
	lines := strings.Split(logo, "\n")
	logoWidth := 0
	logoHeight := len(lines)
	for _, line := range lines {
		if len(line) > logoWidth {
			logoWidth = len(line)
		}
	}
	logoBox := tview.NewTextView().
		SetTextColor(tcell.ColorRed).
		SetDoneFunc(func(key tcell.Key) {
			nextSlide()
		})
	fmt.Fprint(logoBox, logo)

	// Create a frame for the subtitle and navigation infos.
	frame := tview.NewFrame(tview.NewBox()).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText(subtitle, true, tview.AlignCenter, tcell.ColorWhite).
		AddText("", true, tview.AlignCenter, tcell.ColorWhite).
		AddText(navigation, true, tview.AlignCenter, tcell.ColorGreen)

	// Create a Flex layout that centers the logo and subtitle.
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewBox(), 0, 7, false).
		AddItem(tview.NewFlex().
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(logoBox, logoWidth, 1, true).
			AddItem(tview.NewBox(), 0, 1, false), logoHeight, 1, true).
		AddItem(frame, 0, 10, false)

	return "Kripto", flex
}
