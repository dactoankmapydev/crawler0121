package db

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"time"
)

type ElasticDB struct {
	Client *elastic.Client
	Host   string
	Port   string
}

func (es *ElasticDB) NewElasticDB() () {
	for {
		es.Client, _ = elastic.NewClient(elastic.SetURL(fmt.Sprintf("http://%s:%s", es.Host, es.Port)),
			elastic.SetSniff(false),
		)
		info, code, err := es.Client.Ping(fmt.Sprintf(`http://%s:%s`, es.Host, es.Port)).Do(context.Background())
		if err != nil {
			fmt.Printf("connection failed: %v\n", err)
			time.Sleep(3*time.Second)
		} else {
			fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
			break
		}
	}
}
