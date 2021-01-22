package repository

type IocRepo interface {
	CreateIndex(indexName, mapping string) error
	InsertIndex(indexName, id string, record interface{}) error
}
