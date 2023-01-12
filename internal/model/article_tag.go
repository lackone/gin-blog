package model

import (
	"gorm.io/gorm"
	"time"
)

type ArticleTag struct {
	*Model
	ArticleId uint `gorm:"size:32;not null;comment:文章ID" json:"article_id"`
	TagId     uint `gorm:"size:32;not null;comment:标签ID" json:"tag_id"`
}

func (this *ArticleTag) Create(db *gorm.DB) error {
	return db.Create(this).Error
}

func (this *ArticleTag) GetArticleIdsByTagId(db *gorm.DB) ([]uint, error) {
	var articleIds []uint
	err := db.Model(this).Where("tag_id = ? and is_del = ?", this.TagId, 0).Pluck("article_id", &articleIds).Error
	if err != nil {
		return nil, err
	}
	return articleIds, nil
}

func (this *ArticleTag) GetTagIdsByArticleId(db *gorm.DB) ([]uint, error) {
	var tagIds []uint
	err := db.Model(this).Where("article_id = ? and is_del = ?", this.ArticleId, 0).Pluck("tag_id", &tagIds).Error
	if err != nil {
		return nil, err
	}
	return tagIds, nil
}

func (this *ArticleTag) DeleteByArticleId(db *gorm.DB) error {
	return db.Model(this).Where("article_id = ? and is_del = ?", this.ArticleId, 0).Updates(map[string]interface{}{
		"deleted": time.Now().Unix(),
		"is_del":  1,
	}).Error
}

func (this *ArticleTag) DeleteByTagId(db *gorm.DB) error {
	return db.Model(this).Where("tag_id = ? and is_del = ?", this.TagId, 0).Updates(map[string]interface{}{
		"deleted": time.Now().Unix(),
		"is_del":  1,
	}).Error
}

func (this *ArticleTag) Delete(db *gorm.DB) error {
	return db.Model(this).Where("id = ? and is_del = ?", this.Id, 0).Updates(map[string]interface{}{
		"deleted": time.Now().Unix(),
		"is_del":  1,
	}).Error
}
