package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/shawnkost/go-api/currencies"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*currencies.Rate, error) {
	upCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusOK {

	}
}
