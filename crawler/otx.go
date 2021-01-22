package crawler

import (
	"encoding/json"
	"fmt"
	"ioc-provider/helper"
	"ioc-provider/model"
	"ioc-provider/repository"
	"log"
	"math"
	"runtime"
	"strconv"
)

type Data struct {
	Results []Results `json:"results"`
	Count int `json:"count"`
}

type Results struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	AuthorName string `json:"author_name"`
	Modified string `json:"modified"`
	Created string `json:"created"`
	Revision int `json:"revision"`
	Tlp string `json:"tlp"`
	Public int `json:"public"`
	Adversary string `json:"adversary"`
	Indicators []Indicators `json:"indicators"`
	Tags []string `json:"tags"`
	TargetedCountries []string `json:"targeted_countries"`
	MalwareFamilies []string `json:"malware_families"`
	AttackIds []string `json:"attack_ids"`
	References []string `json:"references"`
	Industries []string `json:"industries"`
}

type Indicators struct {
	ID int64 `json:"id"`
	Indicator string `json:"indicator"`
	Type string `json:"type"`
	Created string `json:"created"`
}

func Subscribed(repo repository.IocRepo) {
	queue := helper.NewJobQueue(runtime.NumCPU())
	queue.Start()
	defer queue.Stop()

	postList := make([]model.Post, 0)
	iocList := make([]model.Indicator, 0)
	totalPage := TotalPageOtx()
	for page := 1; page <= totalPage; page++ {
		pathAPI := fmt.Sprintf("https://otx.alienvault.com/api/v1/pulses/subscribed?limit=50&page=%d", page)
		fmt.Println("pathAPI->", pathAPI)
		body, err := helper.HttpClient.GetRequestOtx(pathAPI)
		if err != nil {
			return
		}
		var data Data
		json.Unmarshal(body, &data)

		for _, item := range data.Results {
			post := model.Post{
				PulseID:           item.ID,
				Name:              item.Name,
				Description:       item.Description,
				AuthorName:        item.AuthorName,
				Modified:          item.Modified,
				Created:           item.Created,
				Revision:          item.Revision,
				Tlp: 			   item.Tlp,
				Public:            item.Public,
				Adversary:         item.Adversary,
				Tags:              item.Tags,
				TargetedCountries: item.TargetedCountries,
				Industries:        item.Industries,
				MalwareFamilies:   item.MalwareFamilies,
				AttackIds:         item.AttackIds,
				References:        item.References,
			}
			postList = append(postList, post)
			fmt.Println("post->", post)

			for _, post := range postList {
				queue.Submit(&SubscribedProcess{
					post:    post,
					iocRepo: repo,
				})
			}

			for _, value := range item.Indicators {
				indicator := model.Indicator{
					IocID:       strconv.FormatInt(value.ID, 10),
					Ioc:         value.Indicator,
					IocType:     value.Type,
					CreatedTime: value.Created,
					CrawledTime: "",
					Source:      "otx",
					Category:    item.Tags,
				}
				iocList = append(iocList, indicator)
				fmt.Println("indicator->", indicator)

				for _, ioc := range iocList {
					queue.Submit(&SubscribedProcess{
						indicator: ioc,
						iocRepo:   repo,
					})
				}
			}

		}
	}
	fmt.Println("len postList->", len(postList))
	fmt.Println("len iocList->", len(iocList))
}

func TotalPageOtx() int {
	pathAPI := fmt.Sprintf("https://otx.alienvault.com/api/v1/pulses/subscribed?limit=50")
	fmt.Println("pathAPI->", pathAPI)
	body, err := helper.HttpClient.GetRequestOtx(pathAPI)
	if err != nil {
		return 0
	}
	var data Data
	json.Unmarshal(body, &data)
	countPost := data.Count
	totalPage := math.Ceil(float64(countPost) / float64(50))
	fmt.Println("totalPage->", int(totalPage))
	return int(totalPage)
}

type SubscribedProcess struct {
	indicator model.Indicator
	post      model.Post
	iocRepo   repository.IocRepo
}

func (process *SubscribedProcess) Process() {
	fmt.Println("Save indicator: ", process.indicator.IocID)
	process.iocRepo.CreateIndex(model.IndexNameIndicator, model.MappingIndicator)
	errSaveIndicator := process.iocRepo.InsertIndex(model.IndexNameIndicator, process.indicator.IocID, process.indicator)
	if errSaveIndicator != nil {
		log.Println(errSaveIndicator)
	}

	fmt.Println("Save post: ", process.post.PulseID)
	process.iocRepo.CreateIndex(model.IndexNamePost, model.MappingPost)
	errSavePost := process.iocRepo.InsertIndex(model.IndexNamePost, process.post.PulseID, process.post)
	if errSavePost != nil {
		log.Println(errSavePost)
	}
}
