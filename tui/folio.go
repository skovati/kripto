package tui

import (
    "fmt"
    "strconv"

    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
    "github.com/skovati/kripto/api"
    "github.com/skovati/kripto/portfolio"
)

type FolioCoin struct {
    Name string
    Price float64
    Amount float64
    Value float64
    Percent1H float64
    Percent24H float64
    Percent7D float64
}

// each coin: name, price, amount, value, %1h, %24h, %7d, Price 24, Price 7d
func Portfolio(nextSlide func()) (title string, content tview.Primitive) {
    table := tview.NewTable().
        SetBorders(true).
        SetFixed(1,0)
    rawFolio := portfolio.OpenPortfolio()
    headers := []string{"  Name  ", "  Price  ", "  Amount  ", "  Value  ", " % Change 1h ", " % Change 24h ", " % Change 7d ", "  Price 24h Ago  ", "  Price 7d Ago  "}
    cols := len(headers)

    folio := make([][]string, 0)
    apiData := [4]float64{}

    for i, c := range *rawFolio {
        apiData = api.GetPrice(c.Id)
        folio = append(folio, []string{})
        folio[i] = append(folio[i], c.Name)
        folio[i] = append(folio[i], FormatPrice(apiData[0]))
        folio[i] = append(folio[i], strconv.FormatFloat(c.Amount, 'f', 2, 64))
        folio[i] = append(folio[i], FormatPrice(apiData[0] * c.Amount))
        folio[i] = append(folio[i], strconv.FormatFloat(apiData[1], 'f', 1, 64) + "%")
        folio[i] = append(folio[i], strconv.FormatFloat(apiData[2], 'f', 1, 64) + "%")
        folio[i] = append(folio[i], strconv.FormatFloat(apiData[3], 'f', 1, 64) + "%")
        // calc previous prices based on percentages
        folio[i] = append(folio[i], FormatPrice(apiData[0] / (1+(apiData[2]/100))))
        folio[i] = append(folio[i], FormatPrice(apiData[0] / (1+(apiData[3]/100))))
    }

    fmt.Println(folio[0][0][0])

    for i, s := range headers {
        table.SetCell(0, i,
        tview.NewTableCell(s).
            SetTextColor(tcell.ColorYellow).
            SetAlign(tview.AlignCenter))
    }

    color := tcell.ColorWhite

    for row:=0;row<len(folio);row++ {
        for col:=0;col<cols;col++ {
            if folio[row][col][0] == '-' {
                color = tcell.ColorRed
            } else if col > 3 && col < 7 {
                color = tcell.ColorGreen
            } else {
                color = tcell.ColorWhite
            }
            table.SetCell(row+1, col,
            tview.NewTableCell(folio[row][col]).
                SetTextColor(color).
                SetAlign(tview.AlignCenter))
        }
    }

    // totals := tview.NewTable().
    //     SetBorders(true)

    // totalHeaders

    flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(table, 0, 5, true).
			AddItem(tview.NewBox(), 0, 1, false), 0, 1, true)

    return "Portfolio", flex


}


