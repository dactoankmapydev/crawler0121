package repository

type IocRepo interface {
	ExistsIndex(indexName string) bool
	CreateIndex(indexName, mapping string)
	InsertIndex(indexName, id string, doc interface{}) bool
	ExistsDoc(indexName, id string) bool
}