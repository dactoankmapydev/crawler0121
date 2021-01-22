package crawler

import (
	"encoding/json"
	"fmt"
	"ioc-provider/helper"
	"ioc-provider/model"
	"ioc-provider/repository"
	"math"
)

type OtxResult struct {
	Results []struct {
		PulseID           string   `json:"id"`
		Name              string   `json:"name"`
		Description       string   `json:"description"`
		AuthorName        string   `json:"author_name"`
		Modified          string   `json:"modified"`
		Created           string   `json:"created"`
		Tags              []string `json:"tags"`
		TargetedCountries []string `json:"targeted_countries"`
		Industries        []string `json:"industries"`
		MalwareFamilies   []string `json:"malware_families"`
		AttackIds         []string `json:"attack_ids"`
		References        string   `json:"references"`
		Indicators        []struct {
			IocID   string `json:"id"`
			Ioc     string `json:"indicator"`
			IocType string `json:"type"`
			Created string `json:"created"`
		} `json:"indicators"`
	} `json:"results"`
	Count int64 `json:"count"`
}

func Subscribed(repo repository.IocRepo) {
	//fmt.Println("Subscribed")
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
		var or OtxResult
		json.Unmarshal(body, &or)

		for _, item := range or.Results {
			post := model.Post{
				PulseID:           item.PulseID,
				Name:              item.Name,
				Description:       item.Description,
				AuthorName:        item.AuthorName,
				Modified:          item.Modified,
				Created:           item.Created,
				TargetedCountries: item.TargetedCountries,
				Industries:        item.Industries,
				MalwareFamilies:   item.MalwareFamilies,
				AttackIds:         item.AttackIds,
				References:        item.References,
				Category:          item.Tags,
			}
			postList = append(postList, post)
			fmt.Println("post->", post)
			/*repo.CreateIndex("post", model.MappingPost)
			repo.Index("ioc", post.PulseID, post)*/

			for _, value := range item.Indicators {
				var indicator = model.Indicator{
					IocID:       value.IocID,
					Ioc:         value.Ioc,
					IocType:     value.IocType,
					CreatedTime: value.Created,
					CrawledTime: "",
					Source:      "otx",
					Category:    item.Tags,
				}
				iocList = append(iocList, indicator)
				fmt.Println("indicator->", indicator)
				/*repo.CreateIndex("ioc", model.MappingSample)
				repo.Index("ioc", indicator.IocID, indicator)*/
			}

		}
	}
	fmt.Println("len postList->", len(postList))
	fmt.Println("len iocList->", len(iocList))

	/*queue := helper.NewJobQueue(runtime.NumCPU())
	queue.Start()
	defer queue.Stop()
	for _, ioc := range ioc_list {
		queue.Submit(&SubscribedProcess{
			indicator: ioc,
			iocRepo:   repo,
		})
	}

	for _, post := range post_list {
		queue.Submit(&SubscribedProcess{
			post:    post,
			iocRepo: repo,
		})
	}*/
}

func TotalPageOtx() int {
	pathAPI := fmt.Sprintf("https://otx.alienvault.com/api/v1/pulses/subscribed?limit=50")
	fmt.Println("pathAPI->", pathAPI)
	body, err := helper.HttpClient.GetRequestOtx(pathAPI)
	if err != nil {
		return 0
	}
	var otxResult OtxResult
	json.Unmarshal(body, &otxResult)
	countPost := otxResult.Count
	totalPage := math.Ceil(float64(countPost) / float64(50))
	fmt.Println("totalPage->", int(totalPage))
	return int(totalPage)
}

/*type SubscribedProcess struct {
	indicator model.Indicator
	post      model.Post
	iocRepo   repository.IocRepo
}

func (process *SubscribedProcess) Process() {
    err := process.iocRepo.SearchIndex("ioc", process.indicator.IocID)
    if err != nil {
    	fmt.Println("Add: ", process.indicator.IocID)
    	err := process.iocRepo.Index("ioc", process.indicator.IocID, process.indicator)
    	if err != nil {
    		log.Println(err)
		}
		return
	}
}*/
