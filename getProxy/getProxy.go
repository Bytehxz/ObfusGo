package getProxy

import (
	"fmt"
	"log"
	"github.com/gocolly/colly"
)

// GetProxies 
func GetProxies() []string {
	var proxies []string

	c := colly.NewCollector()

	c.OnHTML("div.table-responsive.fpl-list table tbody tr", func(e *colly.HTMLElement) {
		ip := e.ChildText("td:nth-child(1)")   
		port := e.ChildText("td:nth-child(2)") 
		anonymity := e.ChildText("td:nth-child(5)") 

		if anonymity == "elite proxy" {
			proxy := fmt.Sprintf("%s:%s", ip, port)
			proxies = append(proxies, proxy)
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Error :", err)
	})

	err := c.Visit("https://free-proxy-list.net/")
	if err != nil {
		log.Fatal(err)
	}

	return proxies
}

