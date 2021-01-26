package model

type Compromised struct {
	UID          string `json:"uid"`
	HostName     string `json:"hostname"`
	Src          string `json:"src"`
	VictimHash   string `json:"victim_hash"`
	CreationDate string    `json:"creation_date"`
	TimeStamp    int64    `json:"timestamp"`
	Country      string `json:"country"`
	CrawledTime       string   `json:"crawled_time"`
}
