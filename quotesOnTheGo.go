package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	forismatcURL = "http://api.forismatic.com/api/1.0/?method=getQuote&key=457653&format=json&lang=en"
)

type forismaticResp struct {
	QuoteAuthor, QuoteLink, QuoteText, SenderLink, SenderName string
	// ^^ uppercase fields even when json is small
	// as lowercase are not exported so json package can't access them
}

func getQuote(target interface{}) error {
	resp, err := http.Get(forismatcURL)
	// handle error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	// good case
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func main() {
	fmt.Println("Hello World")
	fRes := new(forismaticResp)
	getQuote(fRes)
	fmt.Println(fRes)
}
