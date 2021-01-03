package main

import (
    "fmt"
    "strconv"

    "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
    "github.com/skovati/kripto/tui"
    // "github.com/skovati/kripto/portfolio"
)

// Slide is a function that returns each slides Primative and has a parameter
// called nextSlide that can be called to advance
type Slide func(nextSlide func()) (title string, content tview.Primitive)

// main tview application
var app = tview.NewApplication()

func main() {
    // create slice of Slides
    slides := []Slide{
		tui.Cover,
        tui.TopCoinsView,
		tui.Portfolio}

    // create pages view
	pages := tview.NewPages()

    // The bottom row has some info on where we are.
	info := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
        SetTextAlign(tview.AlignCenter).
		SetHighlightedFunc(func(added, removed, remaining []string) {
			pages.SwitchToPage(added[0])
		})

    // Create the pages for all slides.
	previousSlide := func() {
		slide, _ := strconv.Atoi(info.GetHighlights()[0])
		slide = (slide - 1 + len(slides)) % len(slides)
		info.Highlight(strconv.Itoa(slide)).
			ScrollToHighlight()
	}

    // define nextSlide function
	nextSlide := func() {
		slide, _ := strconv.Atoi(info.GetHighlights()[0])
		slide = (slide + 1) % len(slides)
		info.Highlight(strconv.Itoa(slide)).
			ScrollToHighlight()
	}

    // populate slides slice
	for index, slide := range slides {
		title, primitive := slide(nextSlide)
		pages.AddPage(strconv.Itoa(index), primitive, true, index == 0)
		fmt.Fprintf(info, `%d ["%d"][teal]%s[white][""]  `, index+1, index, title)
	}

    // highlight curr selection
	info.Highlight("0")

    // Create the main layout.
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(pages, 0, 1, true).
		AddItem(info, 1, 1, false)

	// Shortcuts to navigate the slides.
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlN {
			nextSlide()
			return nil
		} else if event.Key() == tcell.KeyCtrlP {
			previousSlide()
			return nil
		}
		return event
	})

    // Start the application.
	if err := app.SetRoot(layout, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
