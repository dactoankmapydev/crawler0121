package model

type Post struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	AuthorName string `json:"author_name"`
	Modified string `json:"modified"`
	Created string `json:"created"`
	Tags []string `json:"tags"`
	TargetedCountries []string `json:"targeted_countries"`
	MalwareFamilies []string `json:"malware_families"`
	AttackIds []string `json:"attack_ids"`
	References []string `json:"references"`
	Industries []string `json:"industries"`
}


