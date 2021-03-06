package model

import (
	crypto2 "github.com/nwillc/cryptoport/externalapi/crypto"
	"github.com/nwillc/genfuncs/container"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPosition_String(t *testing.T) {
	type fields struct {
		Currency crypto2.Currency
		Holding  decimal.Decimal
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "BTC_1",
			fields: fields{
				Currency: "BTC",
				Holding:  decimal.NewFromFloat(1.0),
			},
			want: "1 BTC",
		},
		{
			name: "ETH_Fractional",
			fields: fields{
				Currency: "ETH",
				Holding:  decimal.NewFromFloat(.001),
			},
			want: "0.001 ETH",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Position{
				Currency: tt.fields.Currency,
				Holding:  tt.fields.Holding,
			}
			assert.Equal(t, tt.want, p.String())
		})
	}
}

func TestPortfolio_String(t *testing.T) {
	type fields struct {
		Positions []Position
	}
	var tests = []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "BTCSingle",
			fields: fields{
				Positions: []Position{{
					Currency: "BTC",
					Holding:  decimal.NewFromFloat(1),
				}},
			},
			want: "1 BTC\n",
		},
		{
			name: "BTCETC",
			fields: fields{
				Positions: []Position{
					{
						Currency: "ETH",
						Holding:  decimal.NewFromFloat(.5),
					},
					{
						Currency: "BTC",
						Holding:  decimal.NewFromFloat(1),
					},
				},
			},
			want: "0.5 ETH\n1 BTC\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Portfolio{
				Positions: tt.fields.Positions,
			}
			assert.Equal(t, tt.want, p.String())
		})
	}
}

func TestPortfolio_Values(t *testing.T) {
	type fields struct {
		Positions []Position
	}
	type args struct {
		prices container.GMap[crypto2.Currency, *crypto2.TickerInfo]
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   container.GMap[crypto2.Currency, decimal.Decimal]
	}{
		{
			name: "BTC_Two",
			fields: fields{
				Positions: []Position{
					{
						Currency: "BTC",
						Holding:  decimal.NewFromFloat(2),
					},
				},
			},
			args: args{
				prices: container.GMap[crypto2.Currency, *crypto2.TickerInfo]{
					"BTC": {
						Currency:  "BTC",
						Price:     decimal.NewFromFloat(1),
						Timestamp: "",
					},
				},
			},
			want: container.GMap[crypto2.Currency, decimal.Decimal]{
				"BTC": decimal.NewFromFloat(2),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Portfolio{
				Positions: tt.fields.Positions,
			}
			got := p.Values(tt.args.prices)
			require.Equal(t, len(got), len(tt.want))
			for k, v := range got {
				assert.True(t, v.Equal(tt.want[k.Currency]), "key %s:%s != %s:%s", k, v, k, tt.want[k.Currency])
			}
		})
	}
}
