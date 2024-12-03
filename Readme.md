# Stock Scraper with Colly

This project is a Go-based web scraper that fetches stock data (company name, price, and percentage change) from Yahoo Finance using the [Colly](https://github.com/gocolly/colly) web scraping framework. The scraped data is saved to a `stocks.csv` file.

## Features
- Fetches stock information for multiple companies (tickers defined in code).
- Extracts company name, stock price, and percentage change.
- Saves the data in a CSV file (`stocks.csv`).

---

## Learnings
What is Colly?
Colly is a powerful and easy-to-use web scraping library for Go. It provides a variety of methods to interact with web pages and extract data from them.

### Key features of Colly:

Asynchronous: Handles multiple requests concurrently.
Flexible Selectors: Use CSS selectors to pinpoint the data you want to scrape.
Event-Driven: Built around callback functions triggered by events like request completion, error occurrence, or HTML parsing.

## Colly Callback Functions

### OnRequest:
Triggered every time a new request is made.
Example:
```go
c.OnRequest(func(r *colly.Request) {
    fmt.Println("Visiting:", r.URL)
})
```


### OnError:
Triggered when an error occurs during a request.
Example:
```go
c.OnError(func(r *colly.Response, err error) {
    log.Println("Error occurred:", err)
})
```


### OnHTML:
Triggered when a specified HTML element is found in the response.
```go
c.OnHTML("div.container", func(e *colly.HTMLElement) {
    fmt.Println("Element Text:", e.Text)
})
```


### Wait:
Waits for all pending asynchronous tasks to complete before moving forward.
```go
c.Wait()
```


### Post:
This makes an HTTP POST request to submit the login form.
```go
e.Request.Post(e.Request.URL.String(), map[string]string{
		"username": "your_username",
		"password": "your_password",
	})
```


### Limit:
function used to set rate-limiting rules for the collector. It defines how the collector should behave when sending requests to a server to avoid overloading the server or getting blocked.
```go
c := colly.NewCollector(
    colly.Limit(&colly.LimitRule{
        DomainGlob:  "*example.com*",
        Delay:       2 * time.Second,       // delay between each request
        RandomDelay: 1 * time.Second,       // adds a random delay before each request, helps make the scraping behavior less predictable  
    }),
)
```


### Scrape Links:
Extract all links (`<a>` tags) from a webpage.moving forward.
```go
c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println("Link found:", link)
	})
```


### Scrape Images:
Collect all image URLs from a webpage.
```go
c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		imageURL := e.Attr("src")
		fmt.Println("Image found:", imageURL)
	})
```


### Handle AJAX Content:
Use colly.Async for scraping websites with dynamic content loaded via JavaScript.
```go
c := colly.NewCollector(
		colly.Async(true),
	)

	c.OnHTML("div.dynamic-content", func(e *colly.HTMLElement) {
		fmt.Println("Dynamic content found:", e.Text)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Error:", err)
	})

	c.Visit("https://example.com")
	c.Wait() // Wait for all asynchronous requests to finish
```


###  Login and Scrape Protected Pages:
Handle login forms and scrape data from pages that require authentication.
```go
c := colly.NewCollector()

	// Handle login form submission
	c.OnHTML("form", func(e *colly.HTMLElement) {
        // URL of the current page,  map where we specify the login form fields (username and password).
		e.Request.Post(e.Request.URL.String(), map[string]string{  
			"username": "your_username",
			"password": "your_password",
		})
	})

	// Scrape data after login
	c.OnHTML("div.profile-info", func(e *colly.HTMLElement) {
		fmt.Println("Profile Info:", e.Text)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Error:", err)
	})

	// Visit the login page
	c.Visit("https://example.com/login")
```


### Set Custom Headers:
Send requests with custom headers, such as a User-Agent or cookies.
```go
c := colly.NewCollector()

	// Set custom headers
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "MyCustomScraper/1.0")
		r.Headers.Set("Authorization", "Bearer YOUR_TOKEN")
		fmt.Println("Visiting:", r.URL)
	})

	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println("Page Title:", e.Text)
	})

	c.Visit("https://example.com")
```

### Rate Limiting:
Prevent getting banned by setting rate limits for requests.
```go
c := colly.NewCollector(
		colly.Limit(&colly.LimitRule{
			DomainGlob:  "*example.com*",
			Delay:       2 * time.Second,
			RandomDelay: 1 * time.Second,
		}),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		fmt.Println("Heading found:", e.Text)
	})

	c.Visit("https://example.com")
```

