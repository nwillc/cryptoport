package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurrencyList(t *testing.T) {
	cl := []Currency{"A", "B", "C"}
	list := CurrencyList(cl)
	assert.Equal(t, "A,B,C", list)
}

func TestPeriodsList(t *testing.T) {
	pl := []Period{"A", "B", "C"}
	list := PeriodList(pl)
	assert.Equal(t, "A,B,C", list)
}
