package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/briandowns/spinner"
)

const (
	forismatcURL = "http://api.forismatic.com/api/1.0/?method=getQuote&key=457653&format=json&lang=en"
	verInfo      = "quotesOnTheGo v0.0.3"
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
	timeout := time.Duration(4 * time.Second)
	client := http.Client{Timeout: timeout}
	resp, err := client.Get(forismatcURL)
	// handle error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	// good case
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func getQuoteRobust() *forismaticResp {
	// 15% times a request fails if requests are being done successively
	// spinner
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()
	// fetch data
	var err error
	fRes := new(forismaticResp)
	for i := 0; i < 5; i++ { // 5 times seems ok
		err = getQuote(fRes)
		// fmt.Print("Iteration ", i)
		if err == nil && fRes != nil && fRes.QuoteText != "" {
			break
		}
	}
	s.Stop() // stop spinner
	return fRes
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
		fRes := getQuoteRobust()
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
