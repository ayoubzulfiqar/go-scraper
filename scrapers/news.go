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
	// collector.SetRequestTimeout(120 * time.Second)

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

// func main() {
// 	NewsCrawlerServer()
// }
