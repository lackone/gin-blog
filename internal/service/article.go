package service

import (
	"github.com/lackone/gin-blog/internal/dao"
	"github.com/lackone/gin-blog/internal/model"
	"github.com/lackone/gin-blog/pkg/app"
)

type GetArticleRequest struct {
	Id     uint  `form:"id" binding:"required,gte=1"`
	Status uint8 `form:"status,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	TagId  uint  `form:"tag_id" binding:"gte=1"`
	Status uint8 `form:"status,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	TagId     uint   `form:"tag_id" binding:"required,gte=1"`
	Title     string `form:"title" binding:"required,min=2,max=100"`
	Desc      string `form:"desc" binding:"required,min=2,max=255"`
	Content   string `form:"content" binding:"required,min=2,max=4294967295"`
	Cover     string `form:"cover" binding:"required,url"`
	CreatedBy string `form:"created_by" binding:"required,min=2,max=100"`
	Status    uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	Id        uint   `form:"id" binding:"required,gte=1"`
	TagId     uint   `form:"tag_id" binding:"required,gte=1"`
	Title     string `form:"title" binding:"min=2,max=100"`
	Desc      string `form:"desc" binding:"min=2,max=255"`
	Content   string `form:"content" binding:"min=2,max=4294967295"`
	Cover     string `form:"cover" binding:"url"`
	UpdatedBy string `form:"updated_by" binding:"required,min=2,max=100"`
	Status    uint8  `form:"status,default=1" binding:"oneof=0 1"`
}

type DeleteArticleRequest struct {
	Id uint `form:"id" binding:"required,gte=1"`
}

func (s *Service) GetArticle(p *GetArticleRequest) (*model.Article, error) {
	return s.dao.GetArticle(p.Id, p.Status)
}

func (s *Service) GetArticleList(p *ArticleListRequest, pager *app.Pager) ([]*model.Article, int, error) {
	list, err := s.dao.GetArticleList(p.TagId, p.Status, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}
	count, err := s.dao.GetArticleCount(p.TagId, p.Status)
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (s *Service) CreateArticle(p *CreateArticleRequest) error {
	article, err := s.dao.CreateArticle(&dao.ArticleRequest{
		Title:     p.Title,
		Desc:      p.Desc,
		Content:   p.Content,
		Cover:     p.Cover,
		CreatedBy: p.CreatedBy,
		Status:    p.Status,
	})
	if err != nil {
		return err
	}
	err = s.dao.CreateArticleTag(article.Id, p.TagId)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateArticle(p *UpdateArticleRequest) error {
	err := s.dao.UpdateArticle(&dao.ArticleRequest{
		Id:        p.Id,
		Title:     p.Title,
		Desc:      p.Desc,
		Content:   p.Content,
		Cover:     p.Cover,
		UpdatedBy: p.UpdatedBy,
		Status:    p.Status,
	})
	if err != nil {
		return err
	}
	err = s.dao.DeleteArticleTagByArticleId(p.Id)
	if err != nil {
		return err
	}
	err = s.dao.CreateArticleTag(p.Id, p.TagId)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteArticle(p *DeleteArticleRequest) error {
	err := s.dao.DeleteArticle(p.Id)
	if err != nil {
		return err
	}
	err = s.dao.DeleteArticleTagByArticleId(p.Id)
	if err != nil {
		return err
	}
	return nil
}
