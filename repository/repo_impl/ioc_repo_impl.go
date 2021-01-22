package repo_impl

import (
	"context"
	"errors"
	"github.com/olivere/elastic/v7"
	"ioc-provider/db"
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

func (ioc IocRepoImpl) CreateIndex(index string, mapping string) error {
	ctx := context.Background()
	exists, err := ioc.es.Client.IndexExists(index).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		createIndex, err := ioc.es.Client.CreateIndex(index).
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

func (ioc IocRepoImpl) Index(index string, id string, doc interface{}) error {
	ctx := context.Background()
	_, err := ioc.es.Client.Index().
		Index(index).
		Id(id).
		BodyJson(doc).
		Do(ctx)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (ioc IocRepoImpl) SearchIndex(index string, search string) error {
	ctx := context.Background()
	matchQuery := elastic.NewMatchQuery(model.AppName, search)
	_, err := ioc.es.Client.Search(index).
		Index(index).
		Query(matchQuery).
		Do(ctx)
	if err != nil {
		return err
	}
    return nil
}
