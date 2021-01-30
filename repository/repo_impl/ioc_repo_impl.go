package repo_impl

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"ioc-provider/db"
	"ioc-provider/helper"
	"ioc-provider/model"
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

func (ioc IocRepoImpl) InsertManyIndexIoc(indexName string, docs []model.Indicators) bool {
	bulkRequest := ioc.es.Client.Bulk()
	for _, doc := range docs {
		req := elastic.NewBulkIndexRequest().Index(indexName).Id(helper.Hash(doc.IocID, doc.PostID)).Doc(doc)
		bulkRequest = bulkRequest.Add(req)
	}
	bulkResponse, err := bulkRequest.Do(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	if bulkResponse != nil {

	}
	return true
}

func (ioc IocRepoImpl) InsertManyIndexPost(indexName string, docs []model.Post) bool {
	bulkRequest := ioc.es.Client.Bulk()
	for _, doc := range docs {
		req := elastic.NewBulkIndexRequest().Index(indexName).Id(helper.Hash(doc.ID, doc.Modified)).Doc(doc)
		bulkRequest = bulkRequest.Add(req)
	}
	bulkResponse, err := bulkRequest.Do(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	if bulkResponse != nil {

	}
	return true
}

func (ioc IocRepoImpl) ExistsDocIoc(indexName string, docs []model.Indicators) bool {
	for _, doc := range docs {
		exists, err := ioc.es.Client.Exists().
			Index(indexName).Id(helper.Hash(doc.IocID, doc.PostID)).
			Do(context.Background())
		if err != nil {
			log.Println(err)
		}
		return exists
	}
	return true
}

func (ioc IocRepoImpl) ExistsDocPost(indexName string, docs []model.Post) bool {
	for _, doc := range docs {
		exists, err := ioc.es.Client.Exists().
			Index(indexName).Id(helper.Hash(doc.ID, doc.Modified)).
			Do(context.Background())
		if err != nil {
			log.Println(err)
		}
		return exists
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
