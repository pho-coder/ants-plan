package apps

import (
	"time"

	"github.com/pho-coder/ants-plan/btc-tracker/models"
	"github.com/pho-coder/ants-plan/btc-tracker/services"
)

// GetCurrentPrice ...
func GetCurrentPrice() models.AppCurrentPrice {
	appCurrentPrice := models.AppCurrentPrice{}
	marketDetailTrade := services.GetMarketDetailTrade("btcusdt")
	if "ok" == marketDetailTrade.Status {
		appCurrentPrice.Amount = marketDetailTrade.Tick.Data[0].Amount
		appCurrentPrice.Direction = marketDetailTrade.Tick.Data[0].Direction
		appCurrentPrice.Price = marketDetailTrade.Tick.Data[0].Price
		appCurrentPrice.Time = time.Unix(int64(marketDetailTrade.Tick.Data[0].Ts/1000), 0).Format("2006-01-02 15:04:05")
		return appCurrentPrice
	}
	return appCurrentPrice
}
