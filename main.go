package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
)

func crawlCurrency(url string) ([]*CurrencyData, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http.Get failed: %w", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("goquery failed: %w", err)
	}

	var currencyDataList []*CurrencyData

	timeParsed := make(map[string]struct{})

	doc.Find(".rate-content-sight").Each(func(i int, s *goquery.Selection) {
		if s.Parent().Parent().Get(0).Data == "tbody" {

			time := s.Parent().Find("td").Eq(0).Text()
			currency := s.Parent().Find("td").Eq(1).Text()
			buyingSpotRate, err := strconv.ParseFloat(s.Parent().Find("td").Eq(4).Text(), 64)
			if err != nil {
				log.Printf("Error parsing buyingSpotRate: %v", err)
				return
			}

			if _, ok := timeParsed[time]; ok {
				return
			}

			currencyDataList = append(currencyDataList, &CurrencyData{
				time:           time,
				currencyName:   currency,
				buyingSpotRate: buyingSpotRate,
			})
			timeParsed[time] = struct{}{}
		}
	})

	return currencyDataList, nil
}

func main() {

	list, err := crawlCurrency("https://rate.bot.com.tw/xrt/quote/day/EUR?Lang=en-US")

	if err != nil {
		log.Println(err)
		return
	}
	for _, v := range list {
		fmt.Println(v.time)
		fmt.Println(v.currencyName)
		fmt.Println(v.buyingSpotRate)
		fmt.Println(v.GreaterThan(31))
	}
}
