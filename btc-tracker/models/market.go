package models

// MarketDepth ...
type MarketDepth struct {
	ID   int64       `json:"id"`   // 消息ID
	Ts   int64       `json:"ts"`   // 消息声称事件, 单位: 毫秒
	Bids [][]float64 `json:"bids"` // 买盘, [price(成交价), amount(成交量)], 按price降序排列
	Asks [][]float64 `json:"asks"` // 卖盘, [price(成交价), amount(成交量)], 按price升序排列
}

// MarketDepthReturn ...
type MarketDepthReturn struct {
	Status  string      `json:"status"` // 请求状态, ok或者error
	Ts      int64       `json:"ts"`     // 响应生成时间点, 单位: 毫秒
	Tick    MarketDepth `json:"tick"`   // Depth数据
	Ch      string      `json:"ch"`     //  数据所属的Channel, 格式: market.$symbol.depth.$type
	ErrCode string      `json:"err-code"`
	ErrMsg  string      `json:"err-msg"`
}

// TradeDetailData ...
type TradeDetailData struct {
	ID        int64   `json:"id"`        // 成交ID
	Price     float64 `json:"price"`     // 成交价
	Amount    float64 `json:"amount"`    // 成交量
	Direction string  `json:"direction"` // 主动成交方向
	Ts        int64   `json:"ts"`        // 成交时间
}

// TradeDetail ...
type TradeDetail struct {
	ID   int64             `json:"id"`   // 消息ID
	Ts   int64             `json:"ts"`   // 最新成交时间
	Data []TradeDetailData `json:"data"` // 交易细节数据
}

// TradeDetailReturn ...
type TradeDetailReturn struct {
	Status  string      `json:"status"`   // 请求处理结果, "ok"、"error"
	Ts      int64       `json:"ts"`       // 响应生成时间点, 单位毫秒
	Tick    TradeDetail `json:"tick"`     // TradeDetail数据
	Ch      string      `json:"ch"`       // 数据所属的Channel, 格式: market.$symbol.trade.detail
	ErrCode string      `json:"err-code"` // 错误代码
	ErrMsg  string      `json:"err-msg"`  // 错误提示
}

// TradeData ...
type TradeData struct {
	ID        int64   `json:"id"`        //成交ID
	Price     float64 `json:"price"`     // 成交价
	Amount    float64 `json:"amount"`    // 成交量
	Direction string  `json:"direction"` // 主动成交方向
	Ts        int64   `json:"ts"`        // 成交时间
}

// TradeTick ...
type TradeTick struct {
	ID   int64       `json:"id"`   // 消息ID
	Ts   int64       `json:"ts"`   // 最新成交时间
	Data []TradeData `json:"data"` // Trade数据
}

// TradeReturn ...
type TradeReturn struct {
	Status  string      `json:"status"` // 请求状态, ok或者error
	Ch      string      `json:"ch"`     // 数据所属的Channel, 格式: market.$symbol.trade.detail
	Ts      int64       `json:"ts"`     // 发送时间
	Data    []TradeTick `json:"data"`   // 成交记录
	ErrCode string      `json:"err-code"`
	ErrMsg  string      `json:"err-msg"`
}
