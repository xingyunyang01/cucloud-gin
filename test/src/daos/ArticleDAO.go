package daos

import (
	"github.com/xingyunyang01/cucloud-gin/test/src/models"
	"gorm.io/gorm"
)

type ArticleDAO struct{}

func NewArticleDAO() *ArticleDAO {
	return &ArticleDAO{}
}

func (this *ArticleDAO) GetArticleDetail(news *models.ArticleModel, db *gorm.DB) error {
	return db.Table("mynews").Where("id=?", news.NewsId).Find(news).Error
}

func (this *ArticleDAO) UpdateViews(param interface{}, db *gorm.DB) {
	db.Table("mynews").Where("id=?", param).Update("views", gorm.Expr("views+1"))
}
