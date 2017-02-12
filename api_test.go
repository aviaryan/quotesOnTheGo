package main

import (
	"fmt"
	"testing"
)

func TestApi(t *testing.T) {
	fRes := getQuoteRobust()
	if fRes == nil || fRes.QuoteText == "" {
		fmt.Println(fRes)
		t.Error("Unable to fetch quote")
	}
}
