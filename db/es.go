package db

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type ElasticDB struct {
	Client *elastic.Client
	Host string
	Port string
}

func (es *ElasticDB) NewElasticDB() {
	client, err := elastic.NewClient()
	if err != nil {
		fmt.Printf("connection failed: %v\n", err)
	}
	info, code, err := client.Ping("http://" + es.Host + ":" + es.Port).Do(context.Background())
	if err != nil {
		fmt.Printf("connection failed: %v\n", err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
}
