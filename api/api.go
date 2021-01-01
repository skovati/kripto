package api

import (
	"net/http"
    "os"
    "fmt"
    "encoding/json"
	"io/ioutil"
    "github.com/skovati/kripto/file"
)

var CacheDir string = getCacheDir() + "/kripto"
var CachePath string = CacheDir + "/cached_coins.json"

type CacheCoin struct{
  ID string `json:"id"`
  Symbol string `json:"symbol"`
  Name string `json:"name"`
}

func getCacheDir() string {
    ret, err := os.UserCacheDir()
    if err != nil {
        fmt.Println(err)
    }
    return ret
}

func GetPrice(currency string) [4]float64 {
    // convert currency to id
    id := GetCoinInfo(currency)[0]

    // construct url with id
	url := "https://api.coingecko.com/api/v3/coins/" + id + "?localization=false&tickers=false&market_data=true&community_data=false&developer_data=false&sparkline=false"

    // make request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("err")
	}

    // set header and do req
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err")
	}

    // defer closing and set to var
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

    // create nasty nested map to hold json
	var result map[string]map[string]map[string]float64
	json.Unmarshal([]byte(body), &result)

    // grab important info from json
    return [4]float64{
	    result["market_data"]["current_price"]["usd"],
	    result["market_data"]["price_change_percentage_1h_in_currency"]["usd"],
	    result["market_data"]["price_change_percentage_24h_in_currency"]["usd"],
	    result["market_data"]["price_change_percentage_7d_in_currency"]["usd"]}
}

func GetCoinInfo(currency string) [3]string {
    // get slice of supported coins
    supportedCoins := OpenCache()

    // loop through and check if user inputed string is supported
    for _, coin := range supportedCoins {
      if coin.Symbol == currency || coin.Name == currency || coin.ID == currency {
        return [3]string{coin.ID, coin.Symbol, coin.Name}
      }
    }

    // else, the currency is not supported
    // I need to do something important here
    return [3]string{"", "", ""}
}

func OpenCache() []CacheCoin {
    // if coin cache doesn't exist, run api cache function
    if !file.Exists(CachePath) {
        if !GetSupportedCoins() {
            fmt.Println("Error fetching results, check internet connection")
        }
    }

    // read cache file
    portfolioJson, err := ioutil.ReadFile(CachePath)
    if err != nil {
        fmt.Println(err)
    }

    // create slice and unmarshal into it
    supportedCoins := []CacheCoin{}
    json.Unmarshal(portfolioJson, &supportedCoins)

    // return slice
    return supportedCoins
}

func GetSupportedCoins() bool {
    // make dir if it doesn't exist
    err := os.MkdirAll(CacheDir, os.ModePerm)
    if err != nil {
        return false
        fmt.Println(err)
    }

    // set up http request
	url := "https://api.coingecko.com/api/v3/coins/list"
	req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return false
        fmt.Println(err)
    }

    // set header and make request
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return false
        fmt.Println(err)
    }

    // defer closing body and set to body var for reading
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)


    // write json to cachepath
	err = ioutil.WriteFile(CachePath, body, os.ModePerm)
	return err == nil
}
