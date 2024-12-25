package market

import "github.com/liq4u/huobi_golang/pkg/model/base"

type SubscribeTickerResponse struct {
	base.WebSocketResponseBase
	Data *Ticker
	Tick *Ticker
}
