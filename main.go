package main

import (
    "fmt"
    // "github.com/skovati/kripto/portfolio"
    "github.com/skovati/kripto/api"
)

func main() {
    // folio := *portfolio.OpenPortfolio()

    topCoins := api.GetTopCoins(50)

    for i, tc := range topCoins {
        fmt.Printf("---------------\n")
        fmt.Printf("%d. " + tc.Name + ":\n", i+1)
        fmt.Printf("Price: $%.2f\n", tc.Price)
        fmt.Printf("Percent Change 1 Hour: %.2f%%\n", tc.Percent1H)
        fmt.Printf("Market Cap: $%d\n", tc.MarketCap)
    }
}
