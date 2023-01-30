# Scraping with Go

## Basic of HTML Elements

##### 1. Search for tags

In case of tags we just have to write like this for any html tag because in golang single quote represent `runes` so brackets under double quotes " "

```html
for anchor tag 
(".a")

for paragraph Tag
(".p")

same for other tags....
```

##### 2. Search for all div attributes

id considered as attribute associated with div tag

```html
- example
("div[id]")

general
("html-tag[html-attribute]")
```

##### 3. Search for all name attributes

For Example `("div[id=comment]")`
comment is name of id attribute so fo search all related name attributes we have to write like this

```html
("#comment")
```

##### 4. Search for all elements based on class

For Example `("div class=writer")`

```html
(".writer")
```

##### 5. Search for all elements have set same attribute set

```html
("*[html-attribute]")
```

## Scraping in Golang Using Colly and exporting into CSV File

### Example - 1

```go
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
```

### Example - 2

```go
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
```

## Scraping in Golang Using Colly and exporting into JSON File

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)



type NEWS struct {
	TITLE string `json:"title"`
	LINKS string `json:"links"`
	DATE  string `json:"date"`
}

func NewsCrawlerServer() {
	var url string = "https://www.thenews.com.pk/latest-stories"
	fmt.Println("Starting Scraping....")
	collector := colly.NewCollector()
	var data []NEWS
	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})
	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	collector.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	collector.OnHTML(".writter-list-item-story", func(element *colly.HTMLElement) {
		news := &NEWS{}
		element.ForEach(".latest-right", func(_ int, h *colly.HTMLElement) {
			news.TITLE = h.ChildText(".open-section")
			news.LINKS = h.ChildAttr(".open-section", "href")
			news.DATE = h.ChildText(".latestDate")
			data = append(data, *news)
		})
	})
	collector.Visit(url)
	content, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
	}

	os.WriteFile("news.json", content, 0644)
	fmt.Println("NEWS ", len(data))

}
```

## Scraping Table Data in Golang Using colly and Exporting into CSV

```go
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
```

## Scraping Data of Multiple Pages in Golang Using colly and Exporting into CSV

```go
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Book struct {
	Title string
	Price string
}

func Crawling() {
	Request()
	Response()
	HTML()
	NextPageHTML()
	Visiting()
}

func Data(data []string) {
	file, err := os.Create("export.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	headers := []string{"TITLE", "PRICE"}
	writer.Write(headers)
	writer.Write(data)
}

var collector *colly.Collector = colly.NewCollector(
	colly.AllowedDomains("books.toscrape.com"),
)

func requesting(r *colly.Request) {
	fmt.Println("Visiting: ", r.URL)
}

func Request() {
	collector.OnRequest(requesting)
}

// responding

func responding(r *colly.Response) {
	fmt.Println("Response: ", r.StatusCode)
}

func Response() {
	collector.OnResponse(responding)
}

func htmlElement(e *colly.HTMLElement) {

	book := &Book{}
	book.Title = e.ChildAttr(".image_container img", "alt")
	book.Price = e.ChildText(".price_color")

	row := []string{book.Title, book.Price}
	Data(row)
}

func HTML() {
	collector.OnHTML(".product_pod", htmlElement)
	// collector.OnHTML(".next > a", pagination)
}

func pagination(e *colly.HTMLElement) {
	nextPage := e.Request.AbsoluteURL(e.Attr("href"))
	collector.Visit(nextPage)
}

func NextPageHTML() {
	collector.OnHTML(".next > a", pagination)
}

func Visiting() {
	collector.Visit("https://books.toscrape.com/")

}
```
