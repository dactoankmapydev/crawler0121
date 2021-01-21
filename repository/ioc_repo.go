package repository

type IocRepo interface {
	CreateIndexSample(elasticIndexName string)
	CreateIndexPost(elasticIndexName string)
	CreateIndexIndicator(elasticIndexName string)
	CreateIndexCompromised(elasticIndexName string)
}
