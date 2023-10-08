package services

import (
	"errors"

	"github.com/keima483/mess-menu-app/initializers"
	"github.com/keima483/mess-menu-app/models"
	"github.com/keima483/mess-menu-app/repository"
)

func GetMenuItem(messId int) ([]repository.MenuItem, error) {
	var mess repository.Mess
	initializers.DB.Find(&mess, messId)
	if mess.MessName == "" {
		return []repository.MenuItem{}, errors.New("no mess with this id")
	}
	var items []repository.MenuItem
	initializers.DB.Where("mess_id = ?", messId).Find(&items)
	return items, nil
}

func AddMenuItem(messId int, menuItem *models.MenuItemModel) (repository.MenuItem, error) {
	var mess repository.Mess
	initializers.DB.Find(&mess, messId)
	if mess.MessName == "" {
		return repository.MenuItem{}, errors.New("no mess with this id")
	}
	item := repository.MenuFromModel(*menuItem)
	mess.AddItems(&item)
	initializers.DB.Save(&mess)
	return item, nil
}

func UpdateMenuItem(itemId int, menuItem *models.MenuItemModel) (repository.MenuItem, error) {
	var item repository.MenuItem
	initializers.DB.Find(&item, itemId)
	if item.Day == "" {
		return repository.MenuItem{}, errors.New("no mess with this id")
	}
	item.BreakFast = menuItem.BreakFast
	item.Lunch = menuItem.Lunch
	item.Snacks = menuItem.Snacks
	item.Dinner = menuItem.Dinner
	initializers.DB.Save(&item)
	return item, nil
}
