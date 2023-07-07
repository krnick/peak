package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/krnick/peak/crawler"
)

func main() {

	currencyList, err := crawler.CrawlCurrencyInBus("https://rate.bot.com.tw/xrt/quote/day/EUR?Lang=en-US")

	if err != nil {
		log.Println(err)
		return
	}

	if len(os.Args) != 2 {
		fmt.Println("Usage: ./peak <number-for-greaterthan>")
		fmt.Println("E.g.: ./peak 33")
		return
	}

	numGreaterThan, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		return
	}

	// latest one
	currencyList[len(currencyList)-1].PrintCurrency(numGreaterThan)

}
