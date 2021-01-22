package repository

type IocRepo interface {
	CreateIndex(index string, mapping string) error
	Index(index string, id string, doc interface{}) error
	SearchIndex(index string, search string) error
}
