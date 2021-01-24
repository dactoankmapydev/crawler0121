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
)

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func TotalPageOtx() int {
	pathAPI := fmt.Sprintf("https://otx.alienvault.com/api/v1/pulses/subscribed?limit=50")
	fmt.Println("pathAPI->", pathAPI)
	body, err := helper.HttpClient.GetRequestOtx(pathAPI)
	checkError(err)
	var data model.Data
	json.Unmarshal(body, &data)
	countPost := data.Count
	totalPage := math.Ceil(float64(countPost) / float64(50))
	fmt.Println("totalPage->", int(totalPage))
	return int(totalPage)
}

func getDataOnePage(pathAPI string) ([]model.Results, []model.Indicators, error) {
	postList := make([]model.Results, 0)
	iocList := make([]model.Indicators, 0)
	body, err := helper.HttpClient.GetRequestOtx(pathAPI)
	checkError(err)
	var data model.Data
	json.Unmarshal(body, &data)

	for _, item := range data.Results {
		post := model.Results{
			ID:                item.ID,
			Name:              item.Name,
			Description:       item.Description,
			AuthorName:        item.AuthorName,
			//Modified:          item.Modified,
			//Created:           item.Created,
			Tags:              item.Tags,
			TargetedCountries: item.TargetedCountries,
			MalwareFamilies:   item.MalwareFamilies,
			AttackIds:         item.AttackIds,
			Industries:        item.Industries,
			References:        item.References,
		}
		postList = append(postList, post)
		//fmt.Println("post->", post)

		for _, value := range item.Indicators {
			indicator := model.Indicators{
				//IocID:       strconv.FormatInt(value.IocID, 10),
				IocID:       value.IocID,
				Ioc:         value.Ioc,
				IocType:     value.IocType,
				//CreatedTime: value.Created,
				CrawledTime: "",
				Source:      "otx",
				Category:    item.Tags,
			}
			iocList = append(iocList, indicator)
			//fmt.Println("indicator->", indicator)
		}
	}
	return postList, iocList, nil
}

func GetAllDataSubscribed(repo repository.IocRepo) () {
	eg := errgroup.Group{}

	exists := repo.ExistsIndex(model.IndexNamePost)
	if !exists {
		fmt.Println("not exists")
		repo.CreateIndex(model.IndexNamePost, model.MappingPost)
	}

	totalPage := TotalPageOtx()
	var totalPost int = 0
	var totalIoc int = 0
	posts := make([]model.Results, 0)
	iocs := make([]model.Indicators, 0)

	if totalPage > 0 {
		for page := 1; page <= totalPage; page++ {
			pathAPI := fmt.Sprintf("https://otx.alienvault.com/api/v1/pulses/subscribed?limit=50&page=%d", page)
			eg.Go(func() error {
				postList, iocList, err := getDataOnePage(pathAPI)
				checkError(err)

				totalPost += len(postList)
				totalIoc += len(iocList)

				for _, post := range postList {
					onePost := model.Results{
						ID:                post.ID,
						Name:              post.Name,
						Description:       post.Description,
						AuthorName:        post.AuthorName,
						//Modified:          post.Modified,
						//Created:           post.Created,
						Tags:              post.Tags,
						TargetedCountries: post.TargetedCountries,
						Industries:        post.Industries,
						MalwareFamilies:   post.MalwareFamilies,
						AttackIds:         post.AttackIds,
						References:        post.References,
					}

					//fmt.Println(onePost)
					posts = append(posts, onePost)


					existsID := repo.ExistsDoc(model.IndexNamePost, onePost.ID)
					if existsID {
						fmt.Println(existsID)
						break
					} else {
						fmt.Println("insert")
						success := repo.InsertIndex(model.IndexNamePost, onePost.ID, onePost)
						if !success {
							return nil
						}
					}
				}

				for _, ioc := range iocList {

					oneIoc := model.Indicators{
						IocID:       ioc.IocID,
						Ioc:         ioc.Ioc,
						IocType:     ioc.IocType,
						//CreatedTime: ioc.CreatedTime,
						CrawledTime: "",
						Source:      "otx",
						Category:    ioc.Category,
					}

					iocs = append(iocs, oneIoc)

				}
				return nil
			})
		}
		err := eg.Wait()
		checkError(err)
	}
	fmt.Println("len postList:", len(posts))
	fmt.Println("len iocList:", len(iocs))
	fmt.Println("-----------------------------")
	fmt.Println("total post:", totalPost)
	fmt.Println("total ioc:", totalIoc)
}
