package portfolio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
    "github.com/skovati/kripto/file"
    "github.com/skovati/kripto/coin"
    "github.com/skovati/kripto/api"
)

// location of portfolio
var PortfolioDir string = FindPortfolioPath()
var PortfolioPath string = PortfolioDir + "/portfolio.json"

func FindPortfolioPath() string {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        fmt.Println(err)
    }
    return homeDir + "/.local/share/kripto"
}

func OpenPortfolio() *[]coin.Coin {
    // portfolio is a slice of Coins
    portfolio := []coin.Coin{}
    // check if portfolio exists
    if !file.Exists(PortfolioPath) {
        CreatePortfolio()
    }
    // open portfolio location
	portfolioJson, err := ioutil.ReadFile(PortfolioPath)
    // if portfolio could not be found
	if err != nil {
		fmt.Println("File could not be read, try 'kripto init'")
		return nil
	}
    // unmarshal into portfolio var
	json.Unmarshal(portfolioJson, &portfolio)
    // return address of portfolio array
	return &portfolio
}

func CreatePortfolio() {
    // make directory
    os.MkdirAll(PortfolioDir, os.ModePerm)
    // create portfolio at PortfolioPath
    _, err := os.Create(PortfolioPath)
    if err != nil {
        fmt.Println("Cannot create new portfolio.")
        return
    }
    return
}

func SavePortfolio(portfolio *[]coin.Coin) {
	portfolioJson, err := json.MarshalIndent(*portfolio, "", "  ")
	if err != nil {
		fmt.Println("Error encoding as json.")
	}
	err = ioutil.WriteFile(PortfolioPath, portfolioJson, 0770)
	if err != nil {
		fmt.Println("Error writing to file.")
	}
	return
}

func AddCoin(portfolio *[]coin.Coin, currency string, amount float64) bool {
    // check to make sure coin is supported
    info := api.GetCoinInfo(currency)
    if info[0] == "" || info[1] == "" || info[2] == "" {
        return false
    }

    // create coin struct to add based on supported info
    toAdd := coin.Coin{
        Id: info[0],
        Symbol: info[1],
        Name: info[2],
        Amount: amount}

    *portfolio = append(*portfolio, toAdd)
    return true
}

func RemoveCoin(portfolio *[]coin.Coin, idToRemove string) bool {
    index := -1
    for i, c := range *portfolio {
        if c.Id == idToRemove {
            index = i
            break
        }
    }
    if index > -1 && index < len(*portfolio) {
        // set portfolio equal to 
        *portfolio = append((*portfolio)[:index], (*portfolio)[index+1:]...)
        return true
    }
    return false
}

func EditCoin(portfolio *[]coin.Coin, idToEdit string, amount float64) bool {
    for _, c := range *portfolio {
        if c.Id == idToEdit {
            c.Amount = amount
            return true
        }
    }
    return false
}
