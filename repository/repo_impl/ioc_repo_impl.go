package repo_impl

import (
	"context"
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

func (i IocRepoImpl) InsertPost(post model.Post) model.Post {
	return post
}

func (i IocRepoImpl) InsertIndicator(indexName string) model.Indicator {
	ctx := context.Background()
	const mapping = `
	{
		"settings":{
			"number_of_shards": 2,
			"number_of_replicas": 0
		},
		"mappings":{
			"properties":{
				"ioc_id":{
					"type":"text"
				},
				"ioc":{
					"type":"text
				},
				"ioc_type":{
					"type":"text"
				},
				"created_time":{
					"type":"date"
				},
				"crawled_time":{
					"type":"date"
				},
				"source":{
					"type":"text"
				},
				"category":{
					"type":"text"
				}
			}
		}
	}`

	// Sử dụng dIndexExists để kiểm tra xem chỉ mục cụ thể có tồn tại hay không.
	exists, err := i.es.Client.IndexExists(indexName).Do(ctx)
	if err != nil {
		log.Println(err)
	}
	if !exists {
		// Tạo một chỉ mục mới.
		createIndex, err := i.es.Client.CreateIndex(indexName).BodyString(mapping).Do(ctx)
		if err != nil {
			log.Println(err)
		}
		if !createIndex.Acknowledged {}
	}
	return 
}

func (i IocRepoImpl) InsertSample(sample model.Sample) model.Sample {
	return sample
}

func (i IocRepoImpl) InsertCompromised(compromised model.Compromised) model.Compromised {
	return compromised
}
