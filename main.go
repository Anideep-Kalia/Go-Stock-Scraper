package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Stock struct {
	company, price, change string
}

var ticker = []string{
	"MSFT", "IBM", "AAPL", "GOOG", "AMZN",
}

// Function to set up Colly callbacks
// below functions can't be called because these are callback function and will execute after automatically when event is triggered
func setupCallbacks(c *colly.Collector, stocks *[]Stock) {
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})
	c.OnError(func(r* colly.Response,err error){
		log.Println("Error occurred:", err)
	})
	c.OnHTML("div.container", func(e *colly.HTMLElement) {
		stock := Stock{
			company: e.ChildText("h1"),
			price:   e.ChildText("fin-streamer[data-field='regularMarketPrice']"),
			change:  e.ChildText("fin-streamer[data-field='regularMarketChangePercent']"),
		}
	
		fmt.Printf("Scraped: %+v\n", stock)
		*stocks = append(*stocks, stock)
	})
}

// Function to save the scraped data to a CSV file
func saveToCSV(stocks []Stock) {
	file, err := os.Create("stocks.csv")
	if err != nil {
		log.Fatalln("Failed to create CSV file:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header
	writer.Write([]string{"Company", "Price", "Change"})

	// Write stock data
	for _, stock := range stocks {
		writer.Write([]string{stock.company, stock.price, stock.change})
	}

	fmt.Println("Data saved to stocks.csv")
}

func main() {
	stocks := []Stock{}
	c := colly.NewCollector()

	// Set up callbacks
	setupCallbacks(c, &stocks)

	// Visit each ticker URL
	for _, t := range ticker {
		c.Visit("https://finance.yahoo.com/quote/" + t + "/")
	}

	c.Wait() // Ensure all scraping tasks are completed

	// Save the results to a CSV file
	saveToCSV(stocks)
}



