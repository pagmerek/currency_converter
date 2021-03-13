package main

import (
	"testing"
)

func getExchangeRateTest(t *testing.T) {
	got := getExchangeRate("gbp")
	if got < -1 {
		t.Errorf("got %f, below 0", got)
	}
}
