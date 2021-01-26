package repository

import "ioc-provider/model"

type IocRepo interface {
	ExistsIndex(indexName string) bool
	CreateIndex(indexName, mapping string)
	InsertIndex(indexName, id string, doc interface{}) bool
	ExistsDoc(indexName, id string) bool

	InsertManyIndexIoc(indexName, id string, docs []model.Indicators) bool
	InsertManyIndexPost(indexName, id string, docs []model.Post) bool
}
