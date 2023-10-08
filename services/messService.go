package services

import (
	"errors"

	"github.com/keima483/mess-menu-app/initializers"
	"github.com/keima483/mess-menu-app/models"
	"github.com/keima483/mess-menu-app/repository"
)

func GetMess() []repository.Mess {
	var messes []repository.Mess
	initializers.DB.Find(&messes)
	return messes
}

func GetMessById(messId int) (repository.Mess, error) {
	var mess repository.Mess
	initializers.DB.Find(&mess, messId)
	if mess.MessName == "" {
		return repository.Mess{}, errors.New("no mess with this id")
	}
	return mess, nil
}

func AddMess(me *models.MessModel) (repository.Mess, error) {
	if me.MessName == "" {
		return repository.Mess{}, errors.New("enter a mess name atleast")
	}
	mess := repository.MessFromModel(*me)
	initializers.DB.Save(&mess)
	return mess, nil
}

func UpdateMess(messId int, me models.MessModel) (repository.Mess, error) {
	var mess repository.Mess
	initializers.DB.Find(&mess, messId)
	if mess.MessName == "" {
		return repository.Mess{}, errors.New("no mess with this id")
	}
	mess.MessName = me.MessName
	mess.BreakFastTime = me.BreakFastTime
	mess.LunchTime = me.LunchTime
	mess.SnacksTime = me.SnacksTime
	mess.DinnerTime = me.DinnerTime
	initializers.DB.Save(&mess)
	return mess, nil
}

func DeleteMess(messId int) (bool, error) {
	var mess repository.Mess
	initializers.DB.Find(&mess, messId)
	if mess.MessName == "" {
		return false, errors.New("no mess with this id")
	}
	initializers.DB.Delete(&mess)
	return true, nil
}