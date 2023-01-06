package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type PSX struct {
	LDCP    string
	SCRIP   string
	OPEN    string
	HIGH    string
	LOW     string
	CURRENT string
	VOLUME  string
	CHANGE  string
}

func StockTableCrawler() {
	fName := "data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, err: %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	var _url string = "https://www.urdupoint.com/english/"
	// var _fileName string = "psx.json"
	fmt.Println("Service Started....")

	collector := colly.NewCollector()

	collector.OnRequest(onRequest)
	collector.OnResponse(onResponse)
	collector.OnError(onError)
	collector.OnHTML(".table-responsive", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, eh *colly.HTMLElement) {
			psxData := PSX{
				SCRIP:   eh.ChildText("td:nth-child(1)"),
				LDCP:    eh.ChildText("td:nth-child(2)"),
				OPEN:    eh.ChildText("td:nth-child(3)"),
				HIGH:    eh.ChildText("td:nth-child(4)"),
				LOW:     eh.ChildText("td:nth-child(5)"),
				CURRENT: eh.ChildText("td:nth-child(6)"),
				CHANGE:  eh.ChildText("td:nth-child(7)"),
				VOLUME:  eh.ChildText("td:nth-child(8)"),
			}
			writer.Write([]string{
				psxData.SCRIP,
				psxData.LDCP,
				psxData.OPEN,
				psxData.HIGH,
				psxData.LOW,
				psxData.CURRENT,
				psxData.CHANGE,
				psxData.VOLUME,
			})

		})
		fmt.Println("Scrapping Completed")
	})

	// collector.OnHTML(".table-responsive", onHTML)
	fmt.Println("Scrapping Completed")
	collector.Visit(_url)
}

// on Request
func onRequest(r *colly.Request) {
	fmt.Println("Scraping:", r.URL)
}

// on Response

func onResponse(r *colly.Response) {
	fmt.Println("Status:", r.StatusCode)
}

// on ERROR

func onError(r *colly.Response, err error) {
	fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
}

// func main() {
// 	stockTableCrawler()
// }
