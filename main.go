package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/krnick/peak/crawler"
)

func startCrawling(url string, threshold float64) error {

	ticker := time.NewTicker(time.Minute)

	for ; ; <-ticker.C {
		// latest one
		currencyList, err := crawler.CrawlCurrencyInBus("https://rate.bot.com.tw/xrt/quote/day/EUR?Lang=en-US")
		if err != nil {
			log.Println(err)
			return errors.New("crawl currency failed")
		}

		currencyList[len(currencyList)-1].PrintCurrency(threshold)
	}

}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: ./peak <number-for-greaterthan>")
		fmt.Println("E.g.: ./peak 33")
		return
	}

	numGreaterThan, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		return
	}

	go startCrawling("https://rate.bot.com.tw/xrt/quote/day/EUR?Lang=en-US", numGreaterThan)

	select {}

}
