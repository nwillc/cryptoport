package model

import (
	"github.com/nwillc/cryptoport/pkg/externalapi/crypto"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestPosition_String(t *testing.T) {
	type fields struct {
		Currency crypto.Currency
		Holding  float64
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
				Holding:  1,
			},
			want: "1.000000 BTC",
		},
		{
			name: "ETH_Fractional",
			fields: fields{
				Currency: "ETH",
				Holding:  .001,
			},
			want: "0.001000 ETH",
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
					Holding:  1,
				}},
			},
			want: "1.000000 BTC\n",
		},
		{
			name: "BTCETC",
			fields: fields{
				Positions: []Position{
					{
						Currency: "ETH",
						Holding:  .5,
					},
					{
						Currency: "BTC",
						Holding:  1,
					},
				},
			},
			want: "0.500000 ETH\n1.000000 BTC\n",
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
		prices map[crypto.Currency]crypto.TickerInfo
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[Position]float64
	}{
		{
			name: "BTC_Two",
			fields: fields{
				Positions: []Position{
					{
						Currency: "BTC",
						Holding:  2,
					},
				},
			},
			args: args{
				prices: map[crypto.Currency]crypto.TickerInfo{
					"BTC": {
						Currency:  "BTC",
						Price:     1,
						Timestamp: "",
					},
				},
			},
			want: map[Position]float64{
				{
					Currency: "BTC",
					Holding:  2,
				}: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Portfolio{
				Positions: tt.fields.Positions,
			}
			if got := p.Values(tt.args.prices); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}
