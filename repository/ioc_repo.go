package repository

import "ioc-provider/model"

type IocRepo interface {
	CreateIndexSample(sample model.Sample) model.Sample
	CreateIndexPost(post model.Post) model.Post
	CreateIndexIndicator(indexName string) model.Indicator
	CreateIndexCompromised(compromised model.Compromised) model.Compromised
}
