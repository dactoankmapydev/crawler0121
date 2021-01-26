package crawler

import (
	"fmt"
	"github.com/gocolly/colly"
	"ioc-provider/helper"
	"ioc-provider/model"
	"ioc-provider/repository"
	"log"
	"runtime"
	"strings"
	"time"
)

func MirrorPost(repo repository.IocRepo) {
	compromises := make([]model.Compromised, 0)
	c := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36"))
	c.SetRequestTimeout(30 * time.Second)
	//err := c.SetProxy("http://127.0.0.1:3128")
	//if err != nil {
	//	log.Println(err)
	//}
	c.OnHTML("table tbody", func(e *colly.HTMLElement) {
		var mirrorPost model.Compromised
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
			rows := make([]string, 0)
			row.ForEach("td", func(_ int, el *colly.HTMLElement) {
				rows = append(rows, el.Text)})
			mirrorPost.HostName = rows[0]
			mirrorPost.Country = strings.Replace(strings.Replace(strings.TrimSpace(rows[1]), "(", "", -1), ")", "", -1)
			mirrorPost.UID = rows[2]
			mirrorPost.Src = rows[3]
			mirrorPost.CreationDate = convertUTCTime(rows[4])
			mirrorPost.TimeStamp = convertTimestamp(rows[4])
			mirrorPost.VictimHash = helper.Hash(mirrorPost.TimeStamp, mirrorPost.UID, mirrorPost.HostName)
			mirrorPost.CrawledTime = time.Now().Format(time.RFC3339)
			compromises = append(compromises, mirrorPost)
		})
	})

	c.OnScraped(func(r *colly.Response) {
		queue := helper.NewJobQueue(runtime.NumCPU())
		queue.Start()
		defer queue.Stop()

		for _, compromise := range compromises {
			queue.Submit(&MirrorProcess{
				compromised: compromise,
				iocRepo: repo,
			})
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	for i := 1; i < 75187; i++ {
		fullURL := fmt.Sprintf("https://mirror-h.org/archive/page/%d", i)
		c.Visit(fullURL)
		fmt.Println(fullURL)
	}
}

type MirrorProcess struct {
	compromised     model.Compromised
	iocRepo    repository.IocRepo
}

func (process *MirrorProcess) Process() {
	existsCompro := process.iocRepo.ExistsIndex(model.IndexNameCompromised)
	if !existsCompro {
		process.iocRepo.CreateIndex(model.IndexNameCompromised, model.MappingCompromised)
	}
	existsID := process.iocRepo.ExistsDoc(model.IndexNameCompromised, process.compromised.VictimHash)
	if existsID {
		return
	} else {
		success := process.iocRepo.InsertIndex(model.IndexNameCompromised, process.compromised.VictimHash, process.compromised)
		if !success {
			return
		}
	}
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