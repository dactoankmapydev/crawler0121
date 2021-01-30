package repository

import "ioc-provider/model"

type IocRepo interface {
	ExistsIndex(indexName string) bool
	CreateIndex(indexName, mapping string)
	InsertIndex(indexName, id string, doc interface{}) bool
	ExistsDoc(indexName, id string) bool

	ExistsDocIoc(indexName string, docs []model.Indicators) bool
	ExistsDocPost(indexName string, docs []model.Post) bool

	InsertManyIndexIoc(indexName string, docs []model.Indicators) bool
	InsertManyIndexPost(indexName string, docs []model.Post) bool
}
