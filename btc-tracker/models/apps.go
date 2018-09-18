package models

// AppCurrentPrice ...
type AppCurrentPrice struct {
	Price     float64 `json:"price"`     // 成交价
	Amount    float64 `json:"amount"`    // 成交量
	Direction string  `json:"direction"` // 主动成交方向
	Time      string  `json:"time"`      // 成交时间
}
