package repository

import (
	"github.com/keima483/mess-menu-app/models"
	"gorm.io/gorm"
)

type Mess struct {
	gorm.Model
	MessName string
	BreakFastTime string
	LunchTime string
	SnacksTime string
	DinnerTime string
	Items    []*MenuItem `gorm:"foreignKey:MessId"`
}

func MessFromModel(model models.MessModel) Mess {
	return Mess{
		MessName: model.MessName,
		BreakFastTime: model.BreakFastTime,
		LunchTime: model.LunchTime,
		SnacksTime: model.SnacksTime,
		DinnerTime: model.DinnerTime,	
	}
}

func (m *Mess) AddItems(item *MenuItem) {
	items := []*MenuItem {item}
	if m.Items != nil {
		items = append(items, items...)
	}
	m.Items = items
}