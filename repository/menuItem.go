package repository

import (
	"github.com/keima483/mess-menu-app/models"
	"gorm.io/gorm"
)

type MenuItem struct {
	gorm.Model
	MessId    int
	Day       string
	BreakFast string
	Lunch     string
	Snacks    string
	Dinner    string
}

func MenuFromModel(model models.MenuItemModel) MenuItem {
	return MenuItem{
		Day:       model.Day,
		BreakFast: model.BreakFast,
		Lunch:     model.Lunch,
		Snacks:    model.Snacks,
		Dinner:    model.Dinner,
	}
}
