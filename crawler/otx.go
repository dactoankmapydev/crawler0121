package crawler

import (
	"encoding/json"
	"fmt"
	"ioc-provider/helper"
	"ioc-provider/model"
	"ioc-provider/repository"
	"log"
	"math"
	"strconv"
	"time"
)

type Data struct {
	Results []Results `json:"results"`
	Count   int       `json:"count"`
}

type Results struct {
	ID                string       `json:"id"`
	Name              string       `json:"name"`
	Description       string       `json:"description"`
	AuthorName        string       `json:"author_name"`
	Modified          string       `json:"modified"`
	Created           string       `json:"created"`
	Indicators        []Indicators `json:"indicators"`
	Tags              []string     `json:"tags"`
	TargetedCountries []string     `json:"targeted_countries"`
	MalwareFamilies   []string     `json:"malware_families"`
	AttackIds         []string     `json:"attack_ids"`
	References        []string     `json:"references"`
	Industries        []string     `json:"industries"`
}

type Indicators struct {
	ID        int  `json:"id"`
	Indicator string `json:"indicator"`
	Type      string `json:"type"`
	Created   string `json:"created"`
}

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func TotalPage() int {
	pathAPI := fmt.Sprintf("https://otx.alienvault.com/api/v1/pulses/subscribed?limit=50")
	fmt.Println("pathAPI->", pathAPI)
	body, err := helper.HttpClient.GetOtxWithRetries(pathAPI)
	checkError(err)
	var data Data
	json.Unmarshal(body, &data)
	countPost := data.Count
	totalPage := math.Ceil(float64(countPost) / float64(50))
	fmt.Println("totalPage->", int(totalPage))
	return int(totalPage)
}

func Subscribed(repo repository.IocRepo) {
	postList := make([]model.Post, 0)
	iocList := make([]model.Indicators, 0)
	totalPage := TotalPage()

	existsPost := repo.ExistsIndex(model.IndexNamePost)
	if !existsPost {
		repo.CreateIndex(model.IndexNamePost, model.MappingPost)
	}

	existsIoc := repo.ExistsIndex(model.IndexNameIoc)
	if !existsIoc {
		repo.CreateIndex(model.IndexNameIoc, model.MappingIoc)
	}

	for page := 1; page <= totalPage; page++ {
		pathAPI := fmt.Sprintf("https://otx.alienvault.com/api/v1/pulses/subscribed?limit=50&page=%d", page)
		body, err := helper.HttpClient.GetOtxWithRetries(pathAPI)
		checkError(err)
		var data Data
		json.Unmarshal(body, &data)

		trustType := []string{"FileHash-MD5", "FileHash-PEHASH", "FileHash-SHA256", "FileHash-SHA1", "FileHash-IMPHASH", "FileHash-MD5", "URL", "URI", "hostname", "domain", "IPv6", "IPv4", "BitcoinAddress"}
		sample := []string{"FileHash-MD5", "FileHash-PEHASH", "FileHash-SHA256", "FileHash-SHA1", "FileHash-IMPHASH", "FileHash-MD5"}
		url := []string{"URL", "URI"}
		domain := []string{"hostname", "domain"}
		ipaddress := []string{"IPv6", "IPv4", "BitcoinAddress"}

		for _, item := range data.Results {
			post := model.Post{
				ID:                item.ID,
				Name:              item.Name,
				Description:       item.Description,
				AuthorName:        item.AuthorName,
				Modified:          item.Modified,
				Created:           item.Created,
				Tags:              item.Tags,
				TargetedCountries: item.TargetedCountries,
				MalwareFamilies:   item.MalwareFamilies,
				AttackIds:         item.AttackIds,
				Industries:        item.Industries,
				References:        item.References,
				CrawledTime:       time.Now().Format(time.RFC3339),
			}
			postList = append(postList, post)
			existsIdPost := repo.ExistsDoc(model.IndexNamePost, post.ID)
			if existsIdPost {
				fmt.Println("existsIdPost", post.ID)
			} else {
				success := repo.InsertIndex(model.IndexNamePost, post.ID, post)
				if !success {
					return
				}
			}
			for _, value := range item.Indicators {

				_, foundType := Find(trustType, value.Type)
				if foundType {
					_, foundSample := Find(sample, value.Type)
					if foundSample {
						value.Type = "sample"
					}

					_, foundUrl := Find(url, value.Type)
					if foundUrl {
						value.Type = "url"
					}

					_, foundDomain := Find(domain, value.Type)
					if foundDomain {
						value.Type = "domain"
					}

					_, foundIpaddress := Find(ipaddress, value.Type)
					if foundIpaddress {
						value.Type = "ipaddress"
					}
					indicator := model.Indicators{
						//IocID:       strconv.FormatInt(value.IocID, 10),
						IocID:       strconv.Itoa(value.ID),
						Ioc:         value.Indicator,
						IocType:     value.Type,
						CreatedTime: value.Created,
						CrawledTime: time.Now().Format(time.RFC3339),
						Source:      "otx",
						Category:    item.Tags,
						PostID:      item.ID,
					}

					iocList = append(iocList, indicator)
					existsIdIoc := repo.ExistsDoc(model.IndexNameIoc, helper.Hash(indicator.IocID, indicator.PostID, indicator.CrawledTime))
					if existsIdIoc {
						fmt.Println("existsIdIoc", helper.Hash(indicator.IocID, indicator.PostID, indicator.CrawledTime))
						break
					} else {
						success := repo.InsertIndex(model.IndexNameIoc, helper.Hash(indicator.IocID, indicator.PostID, indicator.CrawledTime), indicator)
						if !success {
							return
						}
					}
				}
			}
		}
	}
}

func SubscribedAfter(repo repository.IocRepo) {
	postList := make([]model.Post, 0)
	iocList := make([]model.Indicators, 0)
    timeNow := time.Now().Format(time.RFC3339)
	pathAPI := fmt.Sprintf("https://otx.alienvault.com/api/v1/pulses/subscribed?limit=50&modified_since=%s", timeNow)
	body, err := helper.HttpClient.GetOtxWithRetries(pathAPI)
	checkError(err)
	var data Data
	json.Unmarshal(body, &data)

	trustType := []string{"FileHash-MD5", "FileHash-PEHASH", "FileHash-SHA256", "FileHash-SHA1", "FileHash-IMPHASH", "FileHash-MD5", "URL", "URI", "hostname", "domain", "IPv6", "IPv4", "BitcoinAddress"}
	sample := []string{"FileHash-MD5", "FileHash-PEHASH", "FileHash-SHA256", "FileHash-SHA1", "FileHash-IMPHASH", "FileHash-MD5"}
	url := []string{"URL", "URI"}
	domain := []string{"hostname", "domain"}
	ipaddress := []string{"IPv6", "IPv4", "BitcoinAddress"}

	for _, item := range data.Results {
			post := model.Post{
				ID:                item.ID,
				Name:              item.Name,
				Description:       item.Description,
				AuthorName:        item.AuthorName,
				Modified:          item.Modified,
				Created:           item.Created,
				Tags:              item.Tags,
				TargetedCountries: item.TargetedCountries,
				MalwareFamilies:   item.MalwareFamilies,
				AttackIds:         item.AttackIds,
				Industries:        item.Industries,
				References:        item.References,
				CrawledTime:       time.Now().Format(time.RFC3339),
			}
			postList = append(postList, post)
			existsIdPost := repo.ExistsDoc(model.IndexNamePost, post.ID)
			if existsIdPost {
				fmt.Println("existsIdPost", post.ID)
			} else {
				success := repo.InsertIndex(model.IndexNamePost, post.ID, post)
				if !success {
					return
				}
			}
			for _, value := range item.Indicators {

				_, foundType := Find(trustType, value.Type)
				if foundType {
					_, foundSample := Find(sample, value.Type)
					if foundSample {
						value.Type = "sample"
					}

					_, foundUrl := Find(url, value.Type)
					if foundUrl {
						value.Type = "url"
					}

					_, foundDomain := Find(domain, value.Type)
					if foundDomain {
						value.Type = "domain"
					}

					_, foundIpaddress := Find(ipaddress, value.Type)
					if foundIpaddress {
						value.Type = "ipaddress"
					}
					indicator := model.Indicators{
						//IocID:       strconv.FormatInt(value.IocID, 10),
						IocID:       strconv.Itoa(value.ID),
						Ioc:         value.Indicator,
						IocType:     value.Type,
						CreatedTime: value.Created,
						CrawledTime: time.Now().Format(time.RFC3339),
						Source:      "otx",
						Category:    item.Tags,
						PostID:      item.ID,
					}

					iocList = append(iocList, indicator)
					existsIdIoc := repo.ExistsDoc(model.IndexNameIoc, helper.Hash(indicator.IocID, indicator.PostID, indicator.CrawledTime))
					if existsIdIoc {
						fmt.Println("existsIdIoc", helper.Hash(indicator.IocID, indicator.PostID, indicator.CrawledTime))
						break
					} else {
						success := repo.InsertIndex(model.IndexNameIoc, helper.Hash(indicator.IocID, indicator.PostID, indicator.CrawledTime), indicator)
						if !success {
							return
						}
					}
				}
			}
		}
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
