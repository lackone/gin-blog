package service

import (
	"github.com/lackone/gin-blog/internal/model"
	"github.com/lackone/gin-blog/pkg/app"
)

type CountTagRequest struct {
	Name   string `form:"name" binding:"max=100"`
	Status uint8  `form:"status,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name   string `form:"name" binding:"max=100"`
	Status uint8  `form:"status,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=2,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=2,max=100"`
	Status    uint8  `form:"status,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	Id        uint   `form:"id" binding:"required,gte=1"`
	Name      string `form:"name" binding:"max=100"`
	Status    uint8  `form:"Status" binding:"oneof=0 1"`
	UpdatedBy string `form:"updated_by" binding:"required,min=2,max=100"`
}

type DeleteTagRequest struct {
	Id uint `form:"id" binding:"required,gte=1"`
}

func (s *Service) CountTag(p *CountTagRequest) (int, error) {
	return s.dao.CountTag(p.Name, p.Status)
}

func (s *Service) GetTagList(p *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return s.dao.GetTagList(p.Name, p.Status, pager.Page, pager.PageSize)
}

func (s *Service) CreateTag(p *CreateTagRequest) error {
	return s.dao.CreateTag(p.Name, p.Status, p.CreatedBy)
}

func (s *Service) UpdateTag(p *UpdateTagRequest) error {
	return s.dao.UpdateTag(p.Id, p.Name, p.Status, p.UpdatedBy)
}

func (s *Service) DeleteTag(p *DeleteTagRequest) error {
	return s.dao.DeleteTag(p.Id)
}
