package tui

import (
    "fmt"

    // "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
    "github.com/skovati/kripto/api"
)

func AddCoin(nextSlide func()) (title string, content tview.Primitive) {
    checkSupported := func(currency string, lastChar rune) bool {
        supported, _ := api.GetCoinInfo(currency)
        return supported
    }

    saveData := func() {
        fmt.Println(content)
        return

    }

    form := tview.NewForm().
        AddInputField("Enter coin name or ticker:", "", 20, checkSupported, nil).
        AddInputField("Enter amount:", "", 10, nil, nil).
        AddButton("Save", saveData).
		AddButton("Cancel", nextSlide)
    return "Add New Coin", form
}

