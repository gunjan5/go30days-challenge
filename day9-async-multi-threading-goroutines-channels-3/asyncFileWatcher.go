package main

import (
  "fmt"
  "os"
  _"time"
  "io/ioutil"
  "strings"
  "encoding/csv"
  "strconv"
  "runtime"
)



const watchPath = "./watchme"

func main() {

	runtime.GOMAXPROCS(4)

	for{
		d,_ :=os.Open(watchPath)
		files,_ := d.Readdir(-1)
		for _,fi :=range files{
			filePath := watchPath+"/"+fi.Name()
			f,_:=os.Open(filePath)
			data,_:=ioutil.ReadAll(f)
			f.Close() //can't use defer here cuz our app wont quit until main ends, and this is a forever watcher doh
			os.Remove(filePath)

			go func (data string){
				 reader :=csv.NewReader(strings.NewReader(data))
				 records,_ :=reader.ReadAll()
				 for _,r:= range records{
				 	invoice := new(Invoice)
				 	invoice.Number = r[0]
				 	invoice.Amount,_ = strconv.ParseFloat(r[1],64)
				 	invoice.PONumber,_=strconv.Atoi(r[2])
				 	invoice.InvoiceDate,_=strconv.Atoi(r[3])

				 	fmt.Printf("Received invoice '%v' for $%.2f and submitted for processing, PONumber is %v, Date is %v\n",
				 	 invoice.Number, invoice.Amount, invoice.PONumber, invoice.InvoiceDate)
				 }

				}(string(data))



		}
	}
	
}


type Invoice struct{

	Number string
	Amount float64
	PONumber int
	InvoiceDate int

}
