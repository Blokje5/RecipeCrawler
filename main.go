package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.tasteofhome.com", "tasteofhome.com"),

		colly.CacheDir("./tasteofhome"),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		// Filter anything that does not start with /recipes
		strippedLink := strings.Replace(link, "https://www.tasteofhome.com", "", 1)
		if !strings.HasPrefix(strippedLink, "/recipes") { // TODO implement as URL filters
			return
		}

		e.Request.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnHTML("script[type='application/ld+json']", func(e *colly.HTMLElement) {
		fmt.Println("object: ", e)
	})

	c.Visit("https://www.tasteofhome.com/recipes/")
}
