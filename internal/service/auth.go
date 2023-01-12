package service

import (
	"errors"
)

type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

func (s *Service) CheckAuth(p *AuthRequest) error {
	auth, err := s.dao.GetAuth(p.AppKey, p.AppSecret)
	if err != nil {
		return err
	}
	if auth.Id > 0 {
		return nil
	}
	return errors.New("auth not exist")
}
