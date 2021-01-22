package model

type Indicator struct {
	IocID       string  `json:"ioc_id"`
	Ioc         string   `json:"ioc"`
	IocType     string   `json:"ioc_type"`
	CreatedTime string   `json:"created_time"`
	CrawledTime string   `json:"crawled_time"`
	Source      string   `json:"source"`
	Category    []string `json:"category"`
}
