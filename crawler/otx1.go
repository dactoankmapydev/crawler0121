package crawler

import (
	"encoding/json"
	"fmt"
	"golang.org/x/sync/errgroup"
	"ioc-provider/helper"
	"ioc-provider/model"
	"ioc-provider/repository"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

type Data1 struct {
	Results1 []Results1 `json:"results"`
	Count   int       `json:"count"`
}

type Results1 struct {
	ID                string       `json:"id"`
	Name              string       `json:"name"`
	Description       string       `json:"description"`
	AuthorName        string       `json:"author_name"`
	Modified          string       `json:"modified"`
	Created           string       `json:"created"`
	Indicators1        []Indicators1 `json:"indicators"`
	Tags              []string     `json:"tags"`
	TargetedCountries []string     `json:"targeted_countries"`
	MalwareFamilies   []string     `json:"malware_families"`
	AttackIds         []string     `json:"attack_ids"`
	References        []string     `json:"references"`
	Industries        []string     `json:"industries"`
}

type Indicators1 struct {
	ID        int  `json:"id"`
	Indicator string `json:"indicator"`
	Type      string `json:"type"`
	Created   string `json:"created"`
}

func checkError1(err error) {
	if err != nil {
		log.Println(err)
	}
}

func TotalPage1() int {
	pathAPI := fmt.Sprintf("https://otx.alienvault.com/api/v1/pulses/subscribed?limit=50")
	fmt.Println("pathAPI->", pathAPI)
	body, err := helper.HttpClient.GetOtxWithRetries(pathAPI)
	checkError(err)
	var data Data1
	json.Unmarshal(body, &data)
	countPost := data.Count
	totalPage := math.Ceil(float64(countPost) / float64(50))
	fmt.Println("totalPage->", int(totalPage))
	return int(totalPage)
}

func getDataOnePage(pathAPI string) ([]model.Post, []model.Indicators, error) {
	loc, _ := time.LoadLocation("Europe/London")
	postList := make([]model.Post, 0)
	iocList := make([]model.Indicators, 0)
	body, err := helper.HttpClient.GetOtxWithRetries(pathAPI)
	checkError(err)
	var data Data1
	json.Unmarshal(body, &data)

	trustType := []string{"FileHash-MD5", "FileHash-PEHASH", "FileHash-SHA256", "FileHash-SHA1", "FileHash-IMPHASH", "FileHash-MD5", "URL", "URI", "hostname", "domain", "IPv6", "IPv4", "BitcoinAddress"}
	sample := []string{"FileHash-MD5", "FileHash-PEHASH", "FileHash-SHA256", "FileHash-SHA1", "FileHash-IMPHASH", "FileHash-MD5"}
	url := []string{"URL", "URI"}
	domain := []string{"hostname", "domain"}
	ipaddress := []string{"IPv6", "IPv4", "BitcoinAddress"}

	for _, item := range data.Results1 {

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
			CrawledTime:       strings.Replace(time.Now().In(loc).Format(time.RFC3339), "Z", "", -1),
		}
		postList = append(postList, post)
		//fmt.Println("post->", post)

		for _, value := range item.Indicators1 {
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
					CrawledTime: strings.Replace(time.Now().In(loc).Format(time.RFC3339), "Z", "", -1),
					Source:      "otx",
					Category:    item.Tags,
					PostID:      item.ID,
				}
				iocList = append(iocList, indicator)
				//fmt.Println("indicator->", indicator)
			}


		}
	}
	return postList, iocList, nil
}

func Subscribed1(repo repository.IocRepo) {
	eg := errgroup.Group{}
	existsPost := repo.ExistsIndex(model.IndexNamePost1)
	if !existsPost {
		repo.CreateIndex(model.IndexNamePost1, model.MappingPost)
	}

	existsIoc := repo.ExistsIndex(model.IndexNameIoc1)
	if !existsIoc {
		repo.CreateIndex(model.IndexNameIoc1, model.MappingIoc)
	}

	totalPage := TotalPage1()
	var totalPost int = 0
	var totalIoc int = 0
	posts := make([]model.Post, 0)
	iocs := make([]model.Indicators, 0)
	//posts := []model.Post{}
	//iocs := []model.Indicators{}

	if totalPage > 0 {
		for page := 1; page <= totalPage; page++ {
			pathAPI := fmt.Sprintf("https://otx.alienvault.com/api/v1/pulses/subscribed?limit=50&page=%d", page)
			eg.Go(func() error {
				postList, iocList, err := getDataOnePage(pathAPI)
				checkError(err)
				totalPost += len(postList)
				totalIoc += len(iocList)

				for _, post := range postList {
					onePost := model.Post{
						ID:                post.ID,
						Name:              post.Name,
						Description:       post.Description,
						AuthorName:        post.AuthorName,
						Modified:          post.Modified,
						Created:           post.Created,
						Tags:              post.Tags,
						TargetedCountries: post.TargetedCountries,
						Industries:        post.Industries,
						MalwareFamilies:   post.MalwareFamilies,
						AttackIds:         post.AttackIds,
						References:        post.References,
						CrawledTime:       post.CrawledTime,
					}
					//fmt.Println(onePost)
					posts = append(posts, onePost)
					existsIdPost := repo.ExistsDoc(model.IndexNamePost1, helper.Hash(post.ID, post.Modified))
					if existsIdPost {
						fmt.Println("existsIdPost", helper.Hash(post.ID, post.Modified))
					} else {
						success := repo.InsertIndex(model.IndexNamePost1, helper.Hash(post.ID, post.Modified), post)
						if !success {
							fmt.Println(success)
						}
					}
				}

				for _, ioc := range iocList {
					oneIoc := model.Indicators{
						IocID:       ioc.IocID,
						Ioc:         ioc.Ioc,
						IocType:     ioc.IocType,
						CreatedTime: ioc.CreatedTime,
						CrawledTime: ioc.CrawledTime,
						Source:      "otx",
						Category:    ioc.Category,
						PostID:      ioc.PostID,
					}
					//fmt.Println(oneIoc)
					iocs = append(iocs, oneIoc)
					existsIdIoc := repo.ExistsDoc(model.IndexNameIoc1, helper.Hash(oneIoc.IocID, oneIoc.PostID, oneIoc.CrawledTime))
					if existsIdIoc {
						fmt.Println("existsIdIoc", helper.Hash(oneIoc.IocID, oneIoc.PostID, oneIoc.CrawledTime))
						break
					} else {
						time.Sleep(100*time.Millisecond)
						success := repo.InsertIndex(model.IndexNameIoc1, helper.Hash(oneIoc.IocID, oneIoc.PostID, oneIoc.CrawledTime), oneIoc)
						time.Sleep(100*time.Millisecond)
						if !success {
							fmt.Println(success)
						}
					}
				}
				return nil
			})
		}
		//err := eg.Wait()
		//checkError(err)
		if err := eg.Wait(); err != nil {
			return
		}
	}
	fmt.Println("len postList:", len(posts))
	fmt.Println("len iocList:", len(iocs))
	fmt.Println("-----------------------------")
	fmt.Println("total post:", totalPost)
	fmt.Println("total ioc:", totalIoc)
}

func SubscribedAfter1(repo repository.IocRepo) {
	loc, _ := time.LoadLocation("Europe/London")
	postList := make([]model.Post, 0)
	iocList := make([]model.Indicators, 0)

	timeNow := strings.Replace(time.Now().In(loc).Format(time.RFC3339), "Z", "", -1)
	pathAPI := fmt.Sprintf("https://otx.alienvault.com/api/v1/pulses/subscribed?limit=50&modified_since=%s", timeNow)
	body, err := helper.HttpClient.GetOtxWithRetries(pathAPI)
	checkError(err)
	var data Data1
	json.Unmarshal(body, &data)
	countPost := data.Count

	trustType := []string{"FileHash-MD5", "FileHash-PEHASH", "FileHash-SHA256", "FileHash-SHA1", "FileHash-IMPHASH", "FileHash-MD5", "URL", "URI", "hostname", "domain", "IPv6", "IPv4", "BitcoinAddress"}
	sample := []string{"FileHash-MD5", "FileHash-PEHASH", "FileHash-SHA256", "FileHash-SHA1", "FileHash-IMPHASH", "FileHash-MD5"}
	url := []string{"URL", "URI"}
	domain := []string{"hostname", "domain"}
	ipaddress := []string{"IPv6", "IPv4", "BitcoinAddress"}

	if countPost != 0 {
		for _, item := range data.Results1 {
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
				CrawledTime:       strings.Replace(time.Now().In(loc).Format(time.RFC3339), "Z", "", -1),
			}
			postList = append(postList, post)
			existsIdPost := repo.ExistsDoc(model.IndexNamePost1, helper.Hash(post.ID, post.Modified))
			if existsIdPost {
				fmt.Println("existsIdPost", helper.Hash(post.ID, post.Modified))
			} else {
				success := repo.InsertIndex(model.IndexNamePost1, helper.Hash(post.ID, post.Modified), post)
				if !success {
					return
				}
			}
			for _, value := range item.Indicators1 {

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
						CrawledTime: strings.Replace(time.Now().In(loc).Format(time.RFC3339), "Z", "", -1),
						Source:      "otx",
						Category:    item.Tags,
						PostID:      item.ID,
					}

					iocList = append(iocList, indicator)
					existsIdIoc := repo.ExistsDoc(model.IndexNameIoc1, helper.Hash(indicator.IocID, indicator.PostID, indicator.CrawledTime))
					if existsIdIoc {
						fmt.Println("existsIdIoc", helper.Hash(indicator.IocID, indicator.PostID, indicator.CrawledTime))
						break
					} else {

						success := repo.InsertIndex(model.IndexNameIoc1, helper.Hash(indicator.IocID, indicator.PostID, indicator.CrawledTime), indicator)
						if !success {
							return
						}
					}
				}
			}
		}
	} else {
		fmt.Println("Not pulse new")
	}
}

func Find1(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
