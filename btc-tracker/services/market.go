package services

import (
	"encoding/json"
	"strconv"

	"github.com/pho-coder/ants-plan/btc-tracker/config"
	"github.com/pho-coder/ants-plan/btc-tracker/models"
	"github.com/pho-coder/ants-plan/btc-tracker/utils"
)

// GetMarketDepth  : 获取交易深度信息
// strSymbol: 交易对, btcusdt, bccbtc......
// strType: Depth类型, step0、step1......stpe5 (合并深度0-5, 0时不合并)
// return: MarketDepthReturn对象
func GetMarketDepth(strSymbol, strType string) models.MarketDepthReturn {
	marketDepthReturn := models.MarketDepthReturn{}

	mapParams := make(map[string]string)
	mapParams["symbol"] = strSymbol
	mapParams["type"] = strType

	strRequestURL := "/market/depth"
	strURL := config.MARKET_URL + strRequestURL

	jsonMarketDepthReturn := utils.HttpGetRequest(strURL, mapParams)
	json.Unmarshal([]byte(jsonMarketDepthReturn), &marketDepthReturn)

	return marketDepthReturn
}

// GetMarketDetailTrade : 获取交易细节信息
// strSymbol: 交易对, btcusdt, bccbtc......
// return: TradeDetailReturn对象
func GetMarketDetailTrade(strSymbol string) models.TradeDetailReturn {
	tradeDetailReturn := models.TradeDetailReturn{}

	mapParams := make(map[string]string)
	mapParams["symbol"] = strSymbol

	strRequestURL := "/market/trade"
	strURL := config.MARKET_URL + strRequestURL

	jsonTradeDetailReturn := utils.HttpGetRequest(strURL, mapParams)
	json.Unmarshal([]byte(jsonTradeDetailReturn), &tradeDetailReturn)

	return tradeDetailReturn
}

//GetMarketTrade : 批量获取最近的交易记录
// strSymbol: 交易对, btcusdt, bccbtc......
// nSize: 获取交易记录的数量, 范围1-2000
// return: TradeReturn对象
func GetMarketTrade(strSymbol string, nSize int) models.TradeReturn {
	tradeReturn := models.TradeReturn{}

	mapParams := make(map[string]string)
	mapParams["symbol"] = strSymbol
	mapParams["size"] = strconv.Itoa(nSize)

	strRequestURL := "/market/history/trade"
	strURL := config.MARKET_URL + strRequestURL

	jsonTradeReturn := utils.HttpGetRequest(strURL, mapParams)
	json.Unmarshal([]byte(jsonTradeReturn), &tradeReturn)

	return tradeReturn
}
