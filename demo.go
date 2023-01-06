// package main

// // //books.go

// // func main() {
// // 	crawl()
// // }

// // type Book struct {
// // 	Title string
// // 	Price string
// // }

// // func crawl() {
// // 	file, err := os.Create("export2.csv")
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// // 	defer file.Close()
// // 	writer := csv.NewWriter(file)
// // 	defer writer.Flush()
// // 	headers := []string{"Title", "Price"}
// // 	writer.Write(headers)

// // 	c := colly.NewCollector(
// // 		colly.AllowedDomains("books.toscrape.com"),
// // 	)

// // 	c.OnRequest(func(r *colly.Request) {
// // 		fmt.Println("Visiting: ", r.URL.String())
// // 	})

// // 	c.OnHTML(".next > a", func(e *colly.HTMLElement) {
// // 		nextPage := e.Request.AbsoluteURL(e.Attr("href"))
// // 		c.Visit(nextPage)
// // 	})

// // 	c.OnHTML(".product_pod", func(e *colly.HTMLElement) {
// // 		book := Book{}
// // 		book.Title = e.ChildAttr(".image_container img", "alt")
// // 		book.Price = e.ChildText(".price_color")
// // 		row := []string{book.Title, book.Price}
// // 		err = writer.Write(row)
// // 		if err != nil {
// // 			log.Fatal(err)
// // 		}
// // 	})

// // 	// startUrl := fmt.Sprintf("https://books.toscrape.com/")
// // 	c.Visit("https://books.toscrape.com/")
// // }

// // type Road struct {
// // 	WHITE string
// // 	BLACK string
// // }

// // func main() {
// // 	records := []Road{}
// // 	file, err := os.Create("demo.cvs")
// // 	// defer file.Close()
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}

// // 	writer := csv.NewWriter(file)
// // 	defer writer.Flush()
// // 	for _, record := range records {
// // 		row := []string{record.WHITE, record.BLACK}
// // 		writer.Write(row)
// // 	}
// // }

// import (
// 	"encoding/csv"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/gocolly/colly"
// )

// type Movies struct {
// 	Number      string
// 	Name        string
// 	Year        string
// 	Certificate string
// 	Runtime     string
// 	Genre       string
// 	Rating      string
// 	Vote        string
// 	Gross       string
// }

// func main() {
// 	fetchURL := "https://www.imdb.com/list/ls033609554/"
// 	fileName := "disney-movies.csv"
// 	file, err := os.Create(fileName)
// 	if err != nil {
// 		log.Fatal("ERROR: Could not create file : \n", fileName, err)
// 		return
// 	}
// 	defer file.Close()
// 	writer := csv.NewWriter(file)
// 	defer writer.Flush()

// 	// Write column headers of the text file
// 	writer.Write([]string{"Sl. No.", "Movie Name", "Release Year", "Certificate", "Genre",
// 		"Running time", "Rating", "Number of Votes", "Gross"})

// 	// Instantiate the default Collector
// 	c := colly.NewCollector()

// 	// Before making a request, print "Visiting ..."
// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Println("Visiting: ", r.URL)
// 	})

// 	// Callback when colly finds the entry point to the DOM segment having a movie info
// 	c.OnHTML(`.lister-item-content`, func(e *colly.HTMLElement) {
// 		//Locate and extract different pieces information about each movie
// 		movie := Movies{}
// 		movie.Number = e.ChildText(".lister-item-index")
// 		movie.Name = e.ChildText(".lister-item-index ~ a")
// 		movie.Year = e.ChildText(".lister-item-year")
// 		movie.Runtime = e.ChildText(".runtime")
// 		movie.Certificate = e.ChildText(".certificate")
// 		movie.Genre = e.ChildText(".genre")
// 		movie.Rating = e.ChildText("[class='ipl-rating-star small'] .ipl-rating-star__rating")
// 		movie.Vote = e.ChildAttr("span[name=nv]", "data-value")
// 		movie.Gross = e.ChildText(".text-muted:contains('Gross') ~ span[name=nv]")

// 		// Write all scraped pieces of information to output text file
// 		writer.Write([]string{
// 			movie.Number,
// 			movie.Name,
// 			movie.Year,
// 			movie.Runtime,
// 			movie.Certificate,
// 			movie.Genre,
// 			movie.Rating,
// 			movie.Vote,
// 			movie.Gross,
// 		})
// 		println(movie.Number,
// 			movie.Name,
// 			movie.Year,
// 			movie.Runtime,
// 			movie.Certificate,
// 			movie.Genre,
// 			movie.Rating,
// 			movie.Vote,
// 			movie.Gross)
// 	})

//		// start scraping the page under the given URL
//		c.Visit(fetchURL)
//		fmt.Println("End of scraping: ", fetchURL)
//	}
package main
