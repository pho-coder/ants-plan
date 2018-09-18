package apps

import (
	"github.com/pho-coder/ants-plan/btc-tracker/services"
)

// GetCurrentPrice ...
func GetCurrentPrice() float64 {
	marketDetailTrade := services.GetMarketDetailTrade("btcusdt")
	if "ok" == marketDetailTrade.Status {
		return marketDetailTrade.Tick.Data[0].Price
	}
	return -1
}
