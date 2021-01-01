package main

import (
    "fmt"
    "github.com/skovati/kripto/portfolio"
    "github.com/skovati/kripto/api"
)

func main() {
    folio := *portfolio.OpenPortfolio()

    for _, c := range(folio) {
        fmt.Println(c.Name + ":")
        fmt.Println(api.GetPrice(c.ID))
    }

    portfolio.AddCoin(&folio, "cardano", 275.0)

    for _, c := range(folio) {
        fmt.Println(c.Name + ":")
        fmt.Println(api.GetPrice(c.ID))
    }

    portfolio.RemoveCoin(&folio, "cardano")

    for _, c := range(folio) {
        fmt.Println(c.Name + ":")
        fmt.Println(api.GetPrice(c.ID))
    }
}
