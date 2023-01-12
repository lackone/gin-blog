package dao

import (
	"github.com/lackone/gin-blog/internal/model"
)

func (d *Dao) CountTag(name string, status uint8) (int, error) {
	tag := model.Tag{Name: name, Status: status}
	return tag.Count(d.db)
}

func (d *Dao) GetTagList(name string, status uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, Status: status}
	return tag.List(d.db, page, pageSize)
}

func (d *Dao) CreateTag(name string, status uint8, createdBy string) error {
	tag := model.Tag{
		Name:   name,
		Status: status,
		Model: &model.Model{
			CreatedBy: createdBy,
		},
	}
	return tag.Create(d.db)
}

func (d *Dao) UpdateTag(id uint, name string, status uint8, updatedBy string) error {
	tag := model.Tag{
		Model: &model.Model{
			Id: id,
		},
	}
	values := map[string]interface{}{
		"status":     status,
		"updated_by": updatedBy,
	}
	if name != "" {
		values["name"] = name
	}
	return tag.Update(d.db, values)
}

func (d *Dao) DeleteTag(id uint) error {
	tag := model.Tag{Model: &model.Model{Id: id}}
	return tag.Delete(d.db)
}
