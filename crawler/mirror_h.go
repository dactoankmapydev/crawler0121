package crawler

//import (
//	"fmt"
//	"github.com/gocolly/colly"
//	"ioc-provider/helper"
//	"ioc-provider/model"
//	"ioc-provider/repository"
//	"log"
//	"runtime"
//	"strings"
//	"time"
//)
//
//func MirrorPost(repo repository.IocRepo) {
//	loc, _ := time.LoadLocation("Europe/London")
//	compromises := make([]model.Compromised, 0)
//	c := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36"))
//	c.SetRequestTimeout(30 * time.Second)
//	//err := c.SetProxy("http://127.0.0.1:3128")
//	//if err != nil {
//	//	log.Println(err)
//	//}
//	c.OnHTML("table tbody", func(e *colly.HTMLElement) {
//		var mirrorPost model.Compromised
//		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
//			rows := make([]string, 0)
//			row.ForEach("td", func(_ int, el *colly.HTMLElement) {
//				rows = append(rows, el.Text)})
//			mirrorPost.HostName = rows[0]
//			mirrorPost.Country = strings.Replace(strings.Replace(strings.TrimSpace(rows[1]), "(", "", -1), ")", "", -1)
//			mirrorPost.UID = rows[2]
//			mirrorPost.Src = rows[3]
//			mirrorPost.CreationDate = convertUTCTime(rows[4])
//			mirrorPost.TimeStamp = convertTimestamp(rows[4])
//			mirrorPost.VictimHash = helper.Hash(mirrorPost.TimeStamp, mirrorPost.UID, mirrorPost.HostName)
//			mirrorPost.CrawledTime = time.Now().In(loc).Format(time.RFC3339)
//			compromises = append(compromises, mirrorPost)
//		})
//	})
//
//
//
//	c.OnScraped(func(r *colly.Response) {
//		queue := helper.NewJobQueue(runtime.NumCPU())
//		queue.Start()
//		defer queue.Stop()
//		for _, compromise := range compromises {
//			queue.Submit(&MirrorProcess{
//				compromised: compromise,
//				iocRepo: repo,
//			})
//		}
//	})
//
//	c.OnError(func(r *colly.Response, err error) {
//		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
//	})
//
//	for i := 1; i <= 10; i++ {
//		fullURL := fmt.Sprintf("https://mirror-h.org/archive/page/%d", i)
//		c.Visit(fullURL)
//		fmt.Println(fullURL)
//	}
//}
//
//type MirrorProcess struct {
//	compromised     model.Compromised
//	iocRepo    repository.IocRepo
//}
//
//func (process *MirrorProcess) Process() {
//	existsCompro := process.iocRepo.ExistsIndex(model.IndexNameCompromised)
//	if !existsCompro {
//		process.iocRepo.CreateIndex(model.IndexNameCompromised, model.MappingCompromised)
//	}
//	existsID := process.iocRepo.ExistsDoc(model.IndexNameCompromised, process.compromised.VictimHash)
//	if !existsID {
//		success := process.iocRepo.InsertIndex(model.IndexNameCompromised, process.compromised.VictimHash, process.compromised)
//		fmt.Println("Add", process.compromised.VictimHash)
//		if !success {
//			return
//		}
//	}
//}
//

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
	"ioc-provider/helper"
	"ioc-provider/helper/rabbit"
	"ioc-provider/model"
	"ioc-provider/repository"
	"log"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const urlBase = "https://mirror-h.org/archive"

type Scraper struct {
	doc *goquery.Document
}

func convertUTCTime(strTime string) string{
	layout := "02/01/2006"
	t,_ := time.Parse(layout, strTime)
	return t.Format(time.RFC3339)
}

func convertTimestamp(strTime string) int64 {
	layout := "02/01/2006"
	t,_ := time.Parse(layout, strTime)
	return t.Unix()
}

func findYear(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func getTotalPage() int {
	response, err := helper.HttpClient.GetMirrorWithRetries(urlBase)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Println(err)
	}

	lastPageLink, _ := doc.Find("ul.pagination li:last-child a").Attr("href")
	split := strings.Split(lastPageLink, "/")[3]
	totalPages, _ := strconv.Atoi(split)
	fmt.Println("totalPage->", totalPages)
	return totalPages
}

func getOnePage(pathURL string) ([]model.Compromised, error) {
	response, err := helper.HttpClient.GetMirrorWithRetries(pathURL)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Println(err)
	}
	
	loc, _ := time.LoadLocation("Europe/London")
	compromisesList := make([]model.Compromised, 0)
	trustYear := []string{"2021", "2020", "2019"}

	doc.Find("table tbody").Each(func(index int, tableHtml *goquery.Selection) {
		var mirrorPost model.Compromised
		tableHtml.Find("tr").Each(func(indexTr int, rowHtml *goquery.Selection) {
			row := make([]string, 0)
			rowHtml.Find("td").Each(func(ndexTd int, tableCell *goquery.Selection) {
				row = append(row, tableCell.Text())
			})
			_, foundYear := findYear(trustYear, strings.Split(row[4], "/")[2])
			if foundYear {
				mirrorPost.HostName = row[0]
				mirrorPost.Country = strings.Replace(strings.Replace(strings.TrimSpace(row[1]), "(", "", -1), ")", "", -1)
				mirrorPost.UID = row[2]
				mirrorPost.Src = row[3]
				mirrorPost.CreationDate = strings.Replace(convertUTCTime(row[4]), "Z", "", -1)
				mirrorPost.TimeStamp = convertTimestamp(row[4])
				mirrorPost.VictimHash = helper.Hash(mirrorPost.TimeStamp, mirrorPost.UID, mirrorPost.HostName)
				mirrorPost.CrawledTime = strings.Replace(time.Now().In(loc).Format(time.RFC3339), "Z", "", -1)
				compromisesList = append(compromisesList, mirrorPost)
			}
		})
	})
	return compromisesList, nil
}

func Mirror(repo repository.IocRepo) {
	sem := semaphore.NewWeighted(int64(25*runtime.NumCPU()))
	group, ctx := errgroup.WithContext(context.Background())

	totalPage := getTotalPage()
	for page := 24800; page <= totalPage; page ++ {
		pathURL := fmt.Sprintf("https://mirror-h.org/archive/page/%d", page)
		err := sem.Acquire(ctx, 1)
		if err != nil {
			fmt.Printf("Acquire err = %+v\n", err)
			continue
		}
		group.Go(func() error {
			defer sem.Release(1)

			// do work
			compromisesList, err := getOnePage(pathURL)
			if err != nil {
				log.Println(err)
			}

			queue := helper.NewJobQueue(runtime.NumCPU())
			queue.Start()
			defer queue.Stop()

			queue.Submit(&MirrorProcess{
				compromisesList: compromisesList,
				iocRepo: repo,
			})
			return nil
		})
	}
	if err := group.Wait(); err != nil {
		fmt.Printf("g.Wait() err = %+v\n", err)
	}
	fmt.Println("done!")
}

type MirrorProcess struct {
	compromisesList []model.Compromised
	iocRepo  repository.IocRepo
}

func (process *MirrorProcess) Process() {
	existsCompromises := process.iocRepo.ExistsIndex(model.IndexNameCompromised)
	if !existsCompromises {
		process.iocRepo.CreateIndex(model.IndexNameCompromised, model.MappingCompromised)
	}

	existsIdCompromises := process.iocRepo.ExistsDocCompromised(model.IndexNameCompromised, process.compromisesList)
	if !existsIdCompromises {
		success := process.iocRepo.InsertManyIndexCompromised(model.IndexNameCompromised, process.compromisesList)
		if !success {
			return
		}
		rabbit.PublishCompromised("compromises", process.compromisesList)
	}
}