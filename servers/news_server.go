package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	addr := ":7171"

	http.HandleFunc("/", handler)

	log.Println("listening on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

type NEWS struct {
	TITLE string
	LINKS string
	DATE  string
}

func handler(w http.ResponseWriter, r *http.Request) {
	var url string = "https://www.thenews.com.pk/latest-stories"
	colly.Async(true)

	collector := colly.NewCollector()
	p := &NEWS{}
	data := []NEWS{}
	// count links
	collector.SetRequestTimeout(120 * time.Second)

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
		// news := &NEWS{}

		element.ForEach(".latest-right", func(_ int, h *colly.HTMLElement) {
			// fmt.Println(h.Text)
			// news.TITLE = h.ChildText(".open-section")
			// news.LINKS = h.ChildAttr(".open-section", "href")
			// news.DATE = h.ChildText(".latestDate")
			p.TITLE = h.ChildText(".open-section")
			p.LINKS = h.ChildAttr(".open-section", "href")
			p.DATE = h.ChildText(".latestDate")

			data = append(data, *p)

			// writer.Write([]string{news.TITLE, news.LINKS, news.DATE})
		})
	})
	collector.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})
	collector.Visit(url)
	collector.Wait()

	// dump results
	b, err := json.Marshal(data)
	if err != nil {
		log.Println("Panic: Failed To Serialize Response:", err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}
