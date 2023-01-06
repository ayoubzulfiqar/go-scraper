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

// func main() {
// 	my_scheduler := gocron.NewScheduler(time.UTC)
// 	my_scheduler.Every(2).Minute().Do(Crawling)
// 	my_scheduler.StartBlocking()
// }

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
