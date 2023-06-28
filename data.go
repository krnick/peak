package main

type CurrencyData struct {
	time           string
	currencyName   string
	buyingSpotRate float64
}

func (currencyData *CurrencyData) GreaterThan(num float64) bool {
	if currencyData.buyingSpotRate > num {
		return true
	}
	return false
}
