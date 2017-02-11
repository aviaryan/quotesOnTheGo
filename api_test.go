package main

import "testing"

func TestApi(t *testing.T) {
	fRes := new(forismaticResp)
	getQuote(fRes)
	if fRes == nil || fRes.QuoteText == "" {
		t.Error("Unable to fetch quote")
	}
}
