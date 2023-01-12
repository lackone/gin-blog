package model

import (
	"github.com/lackone/gin-blog/pkg/app"
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	*Model
	Name   string `gorm:"size:32;not null;comment:标签名" json:"name"`
	Status uint8  `gorm:"size:1;not null;default:1;comment:状态" json:"status"`
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (this *Tag) Get(db *gorm.DB) (*Tag, error) {
	var tag Tag
	err := db.Where("id = ? and is_del = ?", this.Id, 0).First(&tag).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (this *Tag) Count(db *gorm.DB) (int, error) {
	var count int64
	if this.Name != "" {
		db = db.Where("name = ?", this.Name)
	}
	db = db.Where("status = ?", this.Status)
	if err := db.Model(this).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (this *Tag) Create(db *gorm.DB) error {
	return db.Create(this).Error
}

func (this *Tag) Update(db *gorm.DB, values interface{}) error {
	return db.Model(this).Where("id = ? and is_del = ?", this.Id, 0).Updates(values).Error
}

func (this *Tag) Delete(db *gorm.DB) error {
	return db.Model(this).Where("id = ? and is_del = ?", this.Id, 0).Updates(map[string]interface{}{
		"deleted": time.Now().Unix(),
		"is_del":  1,
	}).Error
}

func (this *Tag) List(db *gorm.DB, page, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	if page > 0 && pageSize > 0 {
		db = db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	if this.Name != "" {
		db = db.Where("name = ?", this.Name)
	}
	db = db.Where("status = ?", this.Status)
	if err := db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (this *Tag) ListByIds(db *gorm.DB, ids []uint) ([]*Tag, error) {
	var tags []*Tag
	err := db.Where("id in (?) and is_del = ?", ids, 0).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}
