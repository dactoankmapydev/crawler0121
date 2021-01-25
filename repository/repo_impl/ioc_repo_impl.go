package repo_impl

import (
	"context"
	"fmt"
	"ioc-provider/db"
	"ioc-provider/repository"
	"log"
)

type IocRepoImpl struct {
	es *db.ElasticDB
}

func NewIocRepo(es *db.ElasticDB) repository.IocRepo {
	return &IocRepoImpl{
		es: es,
	}
}

func (ioc IocRepoImpl) ExistsIndex(indexName string) bool {
	ctx := context.Background()
	exists, err := ioc.es.Client.IndexExists(indexName).Do(ctx)
	if err != nil {
		log.Println(err)
	}
	return exists
}

func (ioc IocRepoImpl) CreateIndex(indexName, mapping string) {
	ctx := context.Background()
	_, _ = ioc.es.Client.CreateIndex(indexName).
		Body(mapping).
		Do(ctx)
}

func (ioc IocRepoImpl) InsertIndex(indexName string, id string, doc interface{}) bool {
	_, err := ioc.es.Client.Index().
		Index(indexName).
		Id(id).
		BodyJson(doc).
		Do(context.Background())
	if err != nil {
		fmt.Println(doc)
		log.Println("client.Index() ERROR: %v", err)
	}
	return true
}

func (ioc IocRepoImpl) ExistsDoc(indexName, id string) bool {
	ctx := context.Background()
	exists, err := ioc.es.Client.Exists().
		Index(indexName).Id(id).
		Do(ctx)
	if err != nil {
		log.Println(err)
	}
	return exists
}