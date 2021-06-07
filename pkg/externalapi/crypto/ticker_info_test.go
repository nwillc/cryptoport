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
