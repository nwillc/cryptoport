package model

import (
	"fmt"
	"github.com/nwillc/cryptoport/pkg/externalapi/crypto"
	"github.com/shopspring/decimal"
	"strings"
)

// Position represents the Holding of a given Currency.
type Position struct {
	Currency crypto.Currency
	Holding  decimal.Decimal
}

var _ fmt.Stringer = (*Position)(nil)

// Portfolio a set of Position.
type Portfolio struct {
	Positions []Position
}

var _ fmt.Stringer = (*Portfolio)(nil)

func (p Position) String() string {
	return fmt.Sprintf("%s %s", p.Holding.String(), p.Currency)
}

func (p Portfolio) String() string {
	var sb strings.Builder
	for _, pos := range p.Positions {
		sb.WriteString(pos.String())
		sb.WriteString("\n")
	}
	return sb.String()
}

// Values calculated for a Portfolio at given crypto.TickerInfo of the crypto.Currency.
func (p Portfolio) Values(prices map[crypto.Currency]crypto.TickerInfo) map[Position]decimal.Decimal {
	values := make(map[Position]decimal.Decimal)
	for _, position := range p.Positions {
		ti, ok := prices[position.Currency]
		if ok {
			values[position] = ti.Price.Mul(position.Holding)
		}
	}
	return values
}
