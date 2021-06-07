package crypto

const (
	tickerPath = "currencies/ticker"
)

// CurrencyService provides services to do with crypto currencies
type CurrencyService interface {
	Tickers(currency []Currency, period []Period) (map[Currency]TickerInfo, error)
}

// NomicsCurrencyService is a CurrencyService implemented using Nomics API.
type NomicsCurrencyService struct {
	client *Client
}

var _ CurrencyService = (*NomicsCurrencyService)(nil)

// NewNomicsCurrencyService creates a Nomics based CurrencyService.
func NewNomicsCurrencyService(client *Client) *NomicsCurrencyService {
	return &NomicsCurrencyService{
		client: client,
	}
}

// Tickers returns TickerInfo for the Currency list and Period list provides.
func (n NomicsCurrencyService) Tickers(currencies []Currency, periods []Period) (map[Currency]TickerInfo, error) {
	params := map[string]string{
		"ids":      CurrencyList(currencies),
		"interval": PeriodList(periods),
	}
	var tiList = make([]TickerInfo, 0)
	err := n.client.get(tickerPath, params, &tiList)
	if err != nil {
		return nil, err
	}
	var mapped = make(map[Currency]TickerInfo)
	for _, ti := range tiList {
		mapped[ti.Currency] = ti
	}
	return mapped, nil
}
