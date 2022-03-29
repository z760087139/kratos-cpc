package internal

import (
	"encoding/json"
	"kratos-project/app/policy/internal/biz"
)

type User struct {
	ID       int64  `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	Password string `gorm:"password"`
}

type ImportPolicy struct {
	ID         int64  `gorm:"column:id"`
	Name       string `gorm:"column:name"`
	TemplateID int64  `gorm:"template_id"`
	Args       string `gorm:"column:args"`
}

func (i ImportPolicy) FromInternal() (*biz.ImportPolicy, error) {
	var args map[string]interface{}
	if err := json.Unmarshal([]byte(i.Args), &args); err != nil {
		return nil, err
	}
	return &biz.ImportPolicy{
		BaseInfo: biz.BaseInfo{
			ID:   i.ID,
			Name: i.Name,
		},
		TemplateID: i.TemplateID,
		Args:       args,
	}, nil
}
