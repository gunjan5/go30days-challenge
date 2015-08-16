package main


import (
  "fmt"
  "io/ioutil"
  "encoding/xml"
  "net/http"
  "time"
)



func main() {

	start := time.Now()
	

	stockSymbols := []string{
		"googl",
		"csco",
		"cien",
		"msft",
		"aapl",
		"t",
		"vz",
		"tmus",
		"s",
	}
	for _,symbol := range stockSymbols {
		resp,_ := http.Get("http://dev.markitondemand.com/api/v2/quote?symbol="+symbol)
		defer resp.Body.Close()

		body, _ :=ioutil.ReadAll(resp.Body)

		quote := new(QuoteResponse)
		xml.Unmarshal(body, &quote)

		fmt.Printf("%s: %.2f \n", quote.Name, quote.LastPrice)

	}
	

	elapsed:= time.Since(start)

	fmt.Printf("Executed in time %s \n ", elapsed)

}


type QuoteResponse struct {
	Status string
	Name string
	LastPrice float32
	Change float32
	ChangePercent float32
	TimeStamp string
	MSDate float32
	MarketCap int
	Volume int
	ChangeYTD float32
	ChangePercentYTD float32
	High float32
	Low float32
	Open float32

}
