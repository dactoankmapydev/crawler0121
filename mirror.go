package main

import (
	"github.com/gocolly/colly"
	"log"
)

type Compromised struct {
	UID          string `json:"uid"`
	HostName     string `json:"hostname"`
	Src          string `json:"src"`
	VictimHash   string `json:"victim_hash"`
	CreationDate int    `json:"creation_date"`
	TimeStamp    int    `json:"timestamp"`
	Country      string `json:"country"`
}

func main() {

	c := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36"))
	err := c.SetProxy("http://127.0.0.1:3128")
    if err != nil {
    	log.Println(err)
	}
	c.OnHTML("table tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
			row.ForEach("td", func(_ int, el *colly.HTMLElement) {
				//fmt.Println(el.Text)
				//fmt.Println(el.Index)
			})
		})
	})
	c.Visit("https://mirror-h.org/archive/page/1")
}


	//c.OnHTML("body", func(body *colly.HTMLElement) {
	//	body.DOM.Find("table.table").First().Find("tr>td").Parent().Each(func(_ int, s *goquery.Selection) {
	//		symbol := s.Find("td a").Text()
	//		fmt.Println(symbol)
	//	})
	//})

