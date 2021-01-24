package model

type Data struct {
	Results []Results `json:"results"`
	Count int `json:"count"`
}

type Results struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	AuthorName string `json:"author_name"`
	//Modified string `json:"modified"`
	//Created string `json:"created"`
	Indicators []Indicators `json:"indicators"`
	Tags []string `json:"tags"`
	TargetedCountries []string `json:"targeted_countries"`
	MalwareFamilies []string `json:"malware_families"`
	AttackIds []string `json:"attack_ids"`
	References []string `json:"references"`
	Industries []string `json:"industries"`
}

type Indicators struct {
	IocID int64 `json:"id"`
	Ioc string `json:"indicator"`
	IocType string `json:"type"`
	//Created string `json:"created"`
	CrawledTime string   `json:"crawled_time"`
	Source      string   `json:"source"`
	Category    []string `json:"category"`
}