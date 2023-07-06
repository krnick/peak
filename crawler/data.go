package crawler

type CurrencyData struct {
	Time           string
	CurrencyName   string
	BuyingSpotRate float64
}

func (currencyData *CurrencyData) GreaterThan(num float64) bool {

	return currencyData.BuyingSpotRate > num
}
