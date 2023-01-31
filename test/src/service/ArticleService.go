package service

import (
	"github.com/xingyunyang01/cucloud-gin/test/src/daos"
	"github.com/xingyunyang01/cucloud-gin/test/src/models"
	"gorm.io/gorm"
)

type ArticleService struct {
	ArticleDao *daos.ArticleDAO `inject:"-"`
}

func NewArticleService() *ArticleService {
	return &ArticleService{}
}

func (this *ArticleService) GetArticleDetail(news *models.ArticleModel, db *gorm.DB) error {
	return this.ArticleDao.GetArticleDetail(news, db)
}

func (this *ArticleService) UpdateViews(param interface{}, db *gorm.DB) {
	this.ArticleDao.UpdateViews(param, db)
}
