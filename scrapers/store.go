// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/gocolly/colly"
// )

// type ITEMS struct {
// 	Name     string
// 	Image    string
// 	Price    string
// 	Url      string
// 	Discount string
// }

// func main() {
// 	Store()
// }

// func Store() {

// 	collector := colly.NewCollector()
// 	collector.SetRequestTimeout(120 * time.Second)
// 	items := make([]ITEMS, 0)
// 	collector.OnRequest(on_request)
// 	collector.OnResponse(on_response)
// 	collector.OnError(on_error)

// 	collector.OnHTML("a.core", func(e *colly.HTMLElement) {
// 		// a.core is belong to single item so we are looping on single item that is why we use foreach loop
// 		e.ForEach("div.name", func(i int, h *colly.HTMLElement) {
// 			item := ITEMS{}
// 			item.Name = h.Text
// 			item.Image = e.ChildAttr("img", "data-src")
// 			item.Price = e.Attr("data-price")
// 			item.Url = "https://jumia.com.ng" + e.Attr("href")
// 			item.Discount = e.ChildText("div.tag._dsct")
// 			items = append(items, item)
// 		})

// 	})
// 	collector.OnScraped(on_scraped)

// 	collector.Visit("https://jumia.com.ng")
// }

// func on_request(r *colly.Request) {
// 	fmt.Println("URL: ", r.URL)
// }

// func on_response(r *colly.Response) {
// 	fmt.Println("Getting Response from: ", r.Request.URL)
// }

// func on_error(r *colly.Response, e error) {
// 	fmt.Println("Error: ", e)
// }

// func on_scraped(r *colly.Response) {
// 	items := &ITEMS{}
// 	fmt.Println("Finished", r.Request.URL)
// 	js, err := json.MarshalIndent(items, "", "    ")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Writing data to file")
// 	if err := os.WriteFile("products.json", js, 0664); err == nil {
// 		fmt.Println("Data written to file successfully")
// 	}
// }

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
)

type PRODUCTS struct {
	Name     string
	Image    string
	Price    string
	Url      string
	Discount string
}

func StoreScrapper() {
	c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)

	var fileName string = "products.csv"
	fmt.Println("Starting Scraping....")

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Panic: Could not be able to Create file", fileName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{"Name", "Quote"})

	// Callbacks
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e, r.StatusCode)
	})
	c.OnHTML(".core", func(e *colly.HTMLElement) {
		e.ForEach(".name", func(_ int, h *colly.HTMLElement) {
			item := &PRODUCTS{}
			item.Name = h.Text
			// item.Image = e.ChildAttr(".img-C", ".data-src")
			item.Price = e.ChildText(".data-price")
			item.Url = "https://jumia.com.ng" + e.Attr(".href")
			// item.Discount = e.ChildText(".div.tag._dsct")

			writer.Write([]string{item.Name, item.Price, item.Url})
		})

	})

	c.Visit("https://www.jumia.com.ng/flash-sales/")
}

// func main()  {
// 	StoreScrapper()
// }
