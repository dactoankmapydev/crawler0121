package model

type Sample struct {
	Name             string   `json:"names"`
	Sha256           string   `json:"sha256"`
	Sha1             string   `json:"sha1"`
	Md5              string   `json:"md5"`
	FirstSubmit      string      `json:"first_submit"`
	NotificationDate string      `json:"notification_date"`
	FileType         string   `json:"file_type"`
	Tags             []string `json:"tags"`
	EnginesDetected  []string `json:"engines_detected"`
	Detected         int      `json:"detected"`
	Point            int      `json:"point"`
	CrawledTime       string   `json:"crawled_time"`
}