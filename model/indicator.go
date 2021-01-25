package model

type Indicators struct {
	IocID int64 `json:"id"`
	Ioc string `json:"indicator"`
	IocType string `json:"type"`
	CreatedTime string `json:"created"`
	CrawledTime string   `json:"crawled_time"`
	Source      string   `json:"source"`
	Category    []string `json:"category"`
}
