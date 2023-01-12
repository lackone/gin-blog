package model

import "gorm.io/gorm"

type Auth struct {
	*Model
	AppKey    string `gorm:"size:32;not null;comment:key" json:"app_key"`
	AppSecret string `gorm:"size:64;not null;comment:secret" json:"app_secret"`
}

func (this *Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth
	err := db.Where("app_key = ? and app_secret = ? and is_del = ?", this.AppKey, this.AppSecret, 0).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}
	return auth, nil
}
