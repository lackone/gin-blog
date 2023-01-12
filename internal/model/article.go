package model

import (
	"github.com/lackone/gin-blog/pkg/app"
	"gorm.io/gorm"
	"time"
)

type Article struct {
	*Model
	Title   string `gorm:"size:128;not null;comment:标题" json:"title"`
	Desc    string `gorm:"size:255;not null;comment:简述" json:"desc"`
	Cover   string `gorm:"size:255;not null;comment:封面" json:"cover"`
	Content string `gorm:"type:longtext;not null;comment:内容" json:"content"`
	Status  uint8  `gorm:"size:1;not null;default:1;comment:状态" json:"status"`
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func (this *Article) Count(db *gorm.DB, tagId uint) (int, error) {
	var count int64
	db = db.Where("status = ?", this.Status)
	articleIds, _ := (&ArticleTag{TagId: tagId}).GetArticleIdsByTagId(db.Session(&gorm.Session{NewDB: true}))
	db = db.Where("id in (?)", articleIds)
	err := db.Model(this).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (this *Article) Get(db *gorm.DB) (*Article, error) {
	var article Article
	err := db.Where("id = ? and status = ? and is_del = ?", this.Id, this.Status, 0).First(&article).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (this *Article) Create(db *gorm.DB) (*Article, error) {
	err := db.Create(this).Error
	if err != nil {
		return nil, err
	}
	return this, nil
}

func (this *Article) Update(db *gorm.DB, values interface{}) error {
	return db.Model(this).Where("id = ? and is_del = ?", this.Id, 0).Updates(values).Error
}

func (this *Article) Delete(db *gorm.DB) error {
	return db.Model(this).Where("id = ? and is_del = ?", this.Id, 0).Updates(map[string]interface{}{
		"deleted": time.Now().Unix(),
		"is_del":  1,
	}).Error
}

func (this *Article) List(db *gorm.DB, tagId uint, page, pageSize int) ([]*Article, error) {
	var articles []*Article
	if page > 0 && pageSize > 0 {
		db = db.Offset((page - 1) * pageSize).Limit(pageSize)
	}

	articleIds, _ := (&ArticleTag{TagId: tagId}).GetArticleIdsByTagId(db.Session(&gorm.Session{NewDB: true}))
	db = db.Where("id in (?)", articleIds)

	err := db.Where("status = ?", this.Status).Where("is_del = ?", 0).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}
