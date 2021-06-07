package crypto

const (
	tickerPath = "currencies/ticker"
)

type CurrencyService interface {
	Ticker(currency []Currency, period []Period) ([]TickerInfo, error)
}

type nomics struct {
	client *Client
}

var _ CurrencyService = (*nomics)(nil)

func NewNomicsCurrencyService(client *Client) *nomics {
	return &nomics{
		client: client,
	}
}

func (n nomics) Ticker(currencies []Currency, periods []Period) ([]TickerInfo, error) {
	params := map[string]string{
		"ids":      CurrencyList(currencies),
		"interval": PeriodList(periods),
	}
	var ti = make([]TickerInfo,0)
	err := n.client.get(tickerPath, params, &ti)
	if err != nil {
		return nil, err
	}
	return ti, nil
}
