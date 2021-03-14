package main

import (
	"testing"
)

func getExchangeRateTest(t *testing.T) {
	implemented_currencies := []string{"gbp"}
	for _, currency := range implemented_currencies {
		got := getExchangeRate(currency)
		if got <= 0 {
			t.Errorf("got %f, below 0", got)
		}
	}
}
