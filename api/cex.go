package api

import (
	"encoding/json"
	"fmt"
	"io"
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
	var response CEXResponse
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body) // Kinda like JS version of await.
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, fmt.Errorf("status code received: %v", res.StatusCode)
	}
	rate := currencies.Rate{Currency: currency, Price: response.Bid}
	return &rate, nil
}
