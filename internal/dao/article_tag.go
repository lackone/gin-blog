package dao

import "github.com/lackone/gin-blog/internal/model"

func (d *Dao) CreateArticleTag(articleId uint, tagId uint) error {
	at := model.ArticleTag{ArticleId: articleId, TagId: tagId}
	return at.Create(d.db)
}

func (d *Dao) DeleteArticleTagByArticleId(articleId uint) error {
	at := model.ArticleTag{ArticleId: articleId}
	return at.DeleteByArticleId(d.db)
}

func (d *Dao) DeleteArticleTagByTagId(tagId uint) error {
	at := model.ArticleTag{TagId: tagId}
	return at.DeleteByTagId(d.db)
}
