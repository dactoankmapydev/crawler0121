package db

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"time"
)

var timeout = time.Duration(15 * time.Second)

type ElasticDB struct {
	Client *elastic.Client
	Host   string
	Port   string
}

func (es *ElasticDB) NewElasticDB() {
	es.Client, _ = elastic.Dial(elastic.SetURL(fmt.Sprintf("http://%s:%s", es.Host, es.Port)),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
	)

	info, code, err := es.Client.Ping(fmt.Sprintf(`http://%s:%s`, es.Host, es.Port)).Do(context.Background())
	if err != nil {
		fmt.Printf("connection failed: %v\n", err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
}
