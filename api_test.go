package main

import "testing"

func TestApi(t *testing.T) {
	fRes := getQuoteRobust()
	if fRes == nil || fRes.QuoteText == "" {
		t.Error("Unable to fetch quote")
	}
}
