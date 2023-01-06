package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
)

// Its our data model based on this we will specify the elements we will scrap
type Quotes struct {
	AUTHOR string
	QUOTE  string
}

func QuoteScrapper() {
	// url of the website that we want to scrap
	var url string = "https://www.brainyquote.com/top_100_quotes"
	// file name of our csv file - yu can give it anything you want
	var fileName string = "quote.csv"
	fmt.Println("Starting Scraping....")
	// using os library we will create a csv file in our directory
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Panic: Could not be able to Create file", fileName, err)
		return
	}
	defer file.Close()
	// writer will write the context of the file
	writer := csv.NewWriter(file)
	defer writer.Flush()
	// first two heading of the CSV file
	writer.Write([]string{"Author", "Quote"})

	// Colly - Initializing our collector
	c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})
	c.OnHTML(".quoteContent", func(h *colly.HTMLElement) {
		quote := &Quotes{}

		quote.AUTHOR = h.ChildText(".bq_fq_a")
		quote.QUOTE = h.ChildText(".b-qt-qt")
		writer.Write([]string{quote.AUTHOR, quote.QUOTE})
	})
	c.Visit(url)
	fmt.Println("End of Era: ", url)
}

// func main() {
// 	QuoteScrapper()
// }
