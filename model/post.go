package model

type Post struct {
	PulseID           string   `json:"pulse_id"`
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	AuthorName        string   `json:"author_name"`
	Modified          string   `json:"modified"`
	Created           string   `json:"created"`
	Revision          int      `json:"revision"`
	Tlp               string   `json:"tlp"`
	Public            int      `json:"public"`
	Adversary         string   `json:"adversary"`
	Tags              []string `json:"tags"`
	TargetedCountries []string `json:"targeted_countries"`
	Industries        []string `json:"industries"`
	MalwareFamilies   []string `json:"malware_families"`
	AttackIds         []string `json:"attack_ids"`
	References        []string `json:"references"`
}
