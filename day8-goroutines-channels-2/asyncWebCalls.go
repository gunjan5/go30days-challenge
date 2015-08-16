package main


import (
  "fmt"
  "io/ioutil"
  "encoding/xml"
  "net/http"
  "time"
  "runtime"
)



func main() {

	runtime.GOMAXPROCS(8)

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
		"ibm", 
	}

	numComplete :=0

	for _,symbol := range stockSymbols {

		go func(symbol string) {
			resp,_ := http.Get("http://dev.markitondemand.com/api/v2/quote?symbol="+symbol)
			defer resp.Body.Close()
			//fmt.Println(resp)
			body, _ :=ioutil.ReadAll(resp.Body)
			//fmt.Println(body)

			quote := new(QuoteResponse)
			xml.Unmarshal(body, &quote)
			//fmt.Println(quote)

			fmt.Printf("%s: %.2f \n", quote.Name, quote.LastPrice)
			numComplete++
		}(symbol)


	}
	
	for numComplete <len(stockSymbols) {
		time.Sleep(10* time.Millisecond)
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
