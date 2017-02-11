package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const (
	forismatcURL = "http://api.forismatic.com/api/1.0/?method=getQuote&key=457653&format=json&lang=en"
	verInfo      = "quotesOnTheGo v0.0.1"
	helpStr      = `
Quotes On The Go
Running quotesOnTheGo is what it takes to make up your day.
Try it now

--help:       Show this help
--version:    Show version information
`
)

type forismaticResp struct {
	QuoteAuthor string
	QuoteLink   string
	QuoteText   string
	SenderLink  string
	SenderName  string
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

func showHelp() {
	fmt.Println(helpStr)
}

func showVersion() {
	fmt.Println(verInfo)
}

func main() {
	// http://thenewstack.io/cli-command-line-programming-with-go/
	argCount := len(os.Args[1:])
	if argCount == 0 {
		fRes := new(forismaticResp)
		getQuote(fRes)
		fmt.Println(fRes.QuoteText)
		fmt.Println("\n-- " + fRes.QuoteAuthor)
	} else if argCount == 1 {
		switch os.Args[1] {
		case "--help":
			showHelp()
		case "--version":
			showVersion()
		default:
			showHelp()
		}
	}
}
