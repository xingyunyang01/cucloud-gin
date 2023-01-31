package configuration

import (
	"github.com/xingyunyang01/cucloud-gin/test/src/daos"
	"github.com/xingyunyang01/cucloud-gin/test/src/service"
)

type ServiceConfig struct{}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{}
}

func (this *ServiceConfig) ArticleDAO() *daos.ArticleDAO {
	return daos.NewArticleDAO()
}

func (this *ServiceConfig) ArticleService() *service.ArticleService {
	return service.NewArticleService()
}
