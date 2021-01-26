package model

type Indicators struct {
	IocID       string    `json:"ioc_id"`
	Ioc         string   `json:"ioc"`
	IocType     string   `json:"type"`
	CreatedTime string   `json:"created"`
	CrawledTime string   `json:"crawled_time"`
	Source      string   `json:"source"`
	Category    []string `json:"category"`
	PostID     string   `json:"post_id"`
}
