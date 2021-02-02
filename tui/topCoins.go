package tui

import (
    "strconv"

    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
    "github.com/skovati/kripto/api"
)

func TopCoinsView(nextSlide func()) (title string, content tview.Primitive) {
    table := tview.NewTable().
        SetBorders(true).
        SetFixed(1,0)
    // set num of top coins to get
    num := 100
    cols, rows := 7, num+1

    // first row
    headers := []string{"  Rank  ","    Coin    ","  Price  ","  % Change 1h  ","  % Change 24h  ","  % Change 7d  ","  Market Cap  "}

    // create channel for TopCoins
    ch1 := make(chan []api.TopCoin)
    // get slice of TopCoins
    go api.GetTopCoins(ch1, rows)

    coins := <- ch1

    // make 2d slice to hold info
    data := make([][]string, rows)

    // each row is a coins: rank, name, price, % 1h, 24h, 7d, market cap
    for i, tc := range coins {
        data[i] = append(data[i], strconv.Itoa(i+1))
        data[i] = append(data[i], tc.Name)
        data[i] = append(data[i], FormatPrice(tc.Price))
        data[i] = append(data[i], strconv.FormatFloat(tc.Percent1H, 'f', 2, 64) + "%")
        data[i] = append(data[i], strconv.FormatFloat(tc.Percent1D, 'f', 2, 64) + "%")
        data[i] = append(data[i], strconv.FormatFloat(tc.Percent1W, 'f', 2, 64) + "%")
        data[i] = append(data[i], FormatMarketCap(tc.MarketCap))
    }

    // set default color
    color := tcell.ColorWhite
    align := tview.AlignCenter
    // now, loop through rows and cols and set each cell
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            // if first row, color yellow
            if r == 0 {
                color = tcell.ColorYellow
                table.SetCell(r, c,
                tview.NewTableCell(headers[c]).
                    SetTextColor(tcell.ColorYellow).
                    SetAlign(tview.AlignCenter))
                    continue
            // if negative, color red
            } else if data[r-1][c][0] == '-' {
                color = tcell.ColorRed
            } else if c == 3 || c == 4 || c == 5 {
                color = tcell.ColorGreen
            } else {
                color = tcell.ColorWhite
            }
            if len(data[r-1][c]) > 14 {
                align = tview.AlignLeft
            } else {
                align = tview.AlignCenter
            }
            switch {
            case r == 1 && c == 0:
                color = tcell.ColorRed
            case r == 2 && c == 0:
                color = tcell.ColorGreen
            case r == 3 && c == 0:
                color = tcell.ColorBlue
            }
            // if a percent and not negative, color green
            table.SetCell(r, c,
            tview.NewTableCell(data[r-1][c]).
                SetTextColor(color).
                SetAlign(align).
                SetMaxWidth(14))
        }
    }

    flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(table, 0, 3, true).
			AddItem(tview.NewBox(), 0, 1, false), 0, 1, true)

    return "Top Coins", flex
}

func FormatMarketCap(marketCap int) string {
    numDigits := len(strconv.Itoa(marketCap))
    switch {
    case numDigits > 12:
        return "$" + strconv.FormatFloat(float64(marketCap) / 1000000000000, 'f', 2, 64) + " T"
    case numDigits > 9:
        return "$" + strconv.FormatFloat(float64(marketCap) / 1000000000, 'f', 2, 64) + " B"
    case numDigits > 6:
        return "$" + strconv.FormatFloat(float64(marketCap) / 1000000, 'f', 2, 64) + " M"
    case numDigits > 3:
        return "$" + strconv.Itoa(marketCap)
    default:
        return "$" + strconv.Itoa(marketCap)
    }
}

func FormatPrice(price float64) string {
    strNum := strconv.FormatFloat(price, 'f', 0, 64)
    numDigits := len(strNum)
    if numDigits > 3 {
        strNum = strNum[:(numDigits-3)] + "," + strNum[:3]
    } else {
        strNum = strconv.FormatFloat(price, 'f', 2, 64)
    }
    return "$" + strNum
}
