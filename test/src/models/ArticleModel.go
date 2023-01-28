package models

import (
	"fmt"
	"time"
)

type ArticleModel struct {
	NewsId      int       `json:"id" gorm:"column:id" uri:"id" binding:"required,gt=0"`
	NewsTitle   string    `json:"title" gorm:"column:newstitle"`
	NewsContent string    `json:"content" gorm:"column:newscontent"`
	Views       int       `json:"views" gorm:"column:views"`
	Addtime     time.Time `json:"addtime" gorm:"column:addtime"`
}

func NewArticleModel() *ArticleModel {
	return &ArticleModel{}
}

// 为了继承Model接口
func (this *ArticleModel) String() string {
	return fmt.Sprintf("NewsId:%d", this.NewsId)
}
