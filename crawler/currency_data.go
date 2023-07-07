package crawler

import "fmt"

type CurrencyData struct {
	Time           string
	CurrencyName   string
	BuyingSpotRate float64
	isGreaterThan  bool
}

func (currencyData *CurrencyData) GreaterThan(num float64) {
	if currencyData.BuyingSpotRate > num {
		currencyData.isGreaterThan = true
	}
}

func (currencyData *CurrencyData) PrintCurrency(threshold float64) {

	fmt.Println("------------------------")
	fmt.Println(currencyData.Time)
	fmt.Println(currencyData.CurrencyName)
	fmt.Println(currencyData.BuyingSpotRate)
	fmt.Println("------------------------")
	if currencyData.BuyingSpotRate >= threshold {
		fmt.Println("is good now")
	}
}
