package repo_impl

import (
	"context"
	"errors"
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

func (ioc IocRepoImpl) CreateIndex(indexName, mapping string) error {
	ctx := context.Background()
	exists, err := ioc.es.Client.IndexExists(indexName).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		createIndex, err := ioc.es.Client.CreateIndex(indexName).
			BodyString(mapping).
			Do(ctx)
		if err != nil {
			return err
		}
		if !createIndex.Acknowledged {
			return errors.New("CreateIndex was not acknowledged. Check that timeout value is correct.")
		}
	}
	return nil
}

func (ioc IocRepoImpl) InsertIndex(indexName string, id string, record interface{}) error {
	ctx := context.Background()
	_, err := ioc.es.Client.Index().
		Index(indexName).
		Id(id).
		BodyJson(record).
		Do(ctx)
	if err != nil {
		log.Println(err)
	}
	return nil
}

