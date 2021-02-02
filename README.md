# kripto
Simple cryptocurrency tracker written in Go.
Uses the [CoinGecko](https://www.coingecko.com/en/api) api to fetch real time prices and persistently stores the portfolio in json format.

![](https://i.imgur.com/XfD2GO0.png)
![](https://i.imgur.com/nJxhmgI.png)

### installation
Unix
```
git clone https://github.com/skovati/kripto
cd kripto
sudo make install
```
Gophers can build from source with 
```
make && sudo make install
```

### usage
```
kripto
```
Ctrl-n and Ctrl-p for movement, Ctrl-c to escape

### uninstall
```
sudo make uninstall
```
