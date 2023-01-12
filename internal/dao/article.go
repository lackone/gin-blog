package dao

import (
	"github.com/lackone/gin-blog/internal/model"
)

type ArticleRequest struct {
	Id        uint   `form:"id" binding:"required,gte=1"`
	TagId     uint   `form:"tag_id" binding:"required,gte=1"`
	Title     string `form:"title" binding:"min=2,max=100"`
	Desc      string `form:"desc" binding:"min=2,max=255"`
	Content   string `form:"content" binding:"min=2,max=4294967295"`
	Cover     string `form:"cover" binding:"url"`
	CreatedBy string `form:"created_by" binding:"required,min=2,max=100"`
	UpdatedBy string `form:"updated_by" binding:"required,min=2,max=100"`
	Status    uint8  `form:"status,default=1" binding:"oneof=0 1"`
}

func (d *Dao) GetArticleCount(tagId uint, status uint8) (int, error) {
	a := model.Article{Status: status}
	return a.Count(d.db, tagId)
}

func (d *Dao) GetArticle(id uint, status uint8) (*model.Article, error) {
	a := model.Article{Model: &model.Model{Id: id}, Status: status}
	return a.Get(d.db)
}

func (d *Dao) GetArticleList(tagId uint, status uint8, page, pageSize int) ([]*model.Article, error) {
	a := model.Article{Status: status}
	return a.List(d.db, tagId, page, pageSize)
}

func (d *Dao) CreateArticle(p *ArticleRequest) (*model.Article, error) {
	a := model.Article{
		Title:   p.Title,
		Status:  p.Status,
		Desc:    p.Desc,
		Content: p.Content,
		Cover:   p.Cover,
		Model: &model.Model{
			CreatedBy: p.CreatedBy,
		},
	}
	return a.Create(d.db)
}

func (d *Dao) UpdateArticle(p *ArticleRequest) error {
	a := model.Article{
		Model: &model.Model{
			Id: p.Id,
		},
	}
	values := map[string]interface{}{
		"title":      p.Title,
		"desc":       p.Desc,
		"content":    p.Content,
		"cover":      p.Cover,
		"updated_by": p.UpdatedBy,
		"status":     p.Status,
	}
	return a.Update(d.db, values)
}

func (d *Dao) DeleteArticle(id uint) error {
	a := model.Article{Model: &model.Model{Id: id}}
	return a.Delete(d.db)
}
