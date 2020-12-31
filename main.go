package main

import (
    "fmt"
    // "github.com/skovati/kripto/coin"
    "github.com/skovati/kripto/portfolio"
    "github.com/skovati/kripto/api"
)

func main() {
    folio := *portfolio.OpenPortfolio()
    for _, c := range(folio) {
        fmt.Println(c)
    }
    fmt.Println(api.GetPrice(folio[0].ID))
}
