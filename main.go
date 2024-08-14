package main

import "github.com/shawnkost/go-api/api"

func main() {
	rate, err := api.GetRate("BTC")
	print(rate, err)
}
