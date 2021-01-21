package repo_impl

import (
	"context"
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

func (i IocRepoImpl) CreateIndexPost(elasticIndexName string) {
	ctx := context.Background()
	const mapping = `
	{
		"settings":{
			"number_of_shards": 2,
			"number_of_replicas": 0
		},
		"mappings":{
			"properties":{
				"pulse_id":{
					"type":"text"
				},
				"name":{
					"type":"text
				},
				"description":{
					"type":"text"
				},
				"author_name":{
					"type":"text"
				},
				"modified":{
					"type":"date"
				},
				"created":{
					"type":"date"
				},
				"targeted_countries":{
					"type":"text"
				},
				"industries":{
					"type":"text"
				}
				"malware_families":{
					"type":"text"
				},
				"attack_ids":{
					"type":"text"
				},
				"references":{
					"type":"text"
				},
				"category":{
					"type":"text"
				}
			}
		}
	}`

	// Sử dụng dIndexExists để kiểm tra xem chỉ mục cụ thể có tồn tại hay không.
	exists, err := i.es.Client.IndexExists(elasticIndexName).Do(ctx)
	if err != nil {
		log.Println(err)
	}
	if !exists {
		// Tạo một chỉ mục mới.
		createIndex, err := i.es.Client.CreateIndex(elasticIndexName).BodyString(mapping).Do(ctx)
		if err != nil {
			log.Println(err)
		}
		if !createIndex.Acknowledged {}
	}
	return
}

func (i IocRepoImpl) CreateIndexIndicator(elasticIndexName string) {
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
	exists, err := i.es.Client.IndexExists(elasticIndexName).Do(ctx)
	if err != nil {
		log.Println(err)
	}
	if !exists {
		// Tạo một chỉ mục mới.
		createIndex, err := i.es.Client.CreateIndex(elasticIndexName).BodyString(mapping).Do(ctx)
		if err != nil {
			log.Println(err)
		}
		if !createIndex.Acknowledged {}
	}
	return
}

func (i IocRepoImpl) CreateIndexSample(elasticIndexName string){
	ctx := context.Background()
	const mapping = `
	{
		"settings":{
			"number_of_shards": 2,
			"number_of_replicas": 0
		},
		"mappings":{
			"properties":{
				"names":{
					"type":"text"
				},
				"sha256":{
					"type":"text
				},
				"sha1":{
					"type":"text"
				},
				"md5":{
					"type":"text"
				},
				"first_submit":{
					"type":"date"
				},
				"notification_date":{
					"type":"date"
				},
				"file_type":{
					"type":"text"
				},
				"tags":{
					"type":"text"
				},
				"engines_detected":{
					"type":"text
				},
				"detected":{
					"type":"integer"
				},
				"point":{
					"type":"integer"
				}
			}
		}
	}`

	// Sử dụng dIndexExists để kiểm tra xem chỉ mục cụ thể có tồn tại hay không.
	exists, err := i.es.Client.IndexExists(elasticIndexName).Do(ctx)
	if err != nil {
		log.Println(err)
	}
	if !exists {
		// Tạo một chỉ mục mới.
		createIndex, err := i.es.Client.CreateIndex(elasticIndexName).BodyString(mapping).Do(ctx)
		if err != nil {
			log.Println(err)
		}
		if !createIndex.Acknowledged {}
	}
	return
}

func (i IocRepoImpl) CreateIndexCompromised(elasticIndexName string) {
	ctx := context.Background()
	const mapping = `
	{
		"settings":{
			"number_of_shards": 2,
			"number_of_replicas": 0
		},
		"mappings":{
			"properties":{
				"uid":{
					"type":"text"
				},
				"hostname":{
					"type":"text
				},
				"src":{
					"type":"text"
				},
				"victim_hash":{
					"type":"text"
				},
				"creation_date":{
					"type":"date"
				},
				"timestamp":{
					"type":"date"
				},
				"country":{
					"type":"text"
				}
			}
		}
	}`

	// Sử dụng dIndexExists để kiểm tra xem chỉ mục cụ thể có tồn tại hay không.
	exists, err := i.es.Client.IndexExists(elasticIndexName).Do(ctx)
	if err != nil {
		log.Println(err)
	}
	if !exists {
		// Tạo một chỉ mục mới.
		createIndex, err := i.es.Client.CreateIndex(elasticIndexName).BodyString(mapping).Do(ctx)
		if err != nil {
			log.Println(err)
		}
		if !createIndex.Acknowledged {}
	}
	return
}


