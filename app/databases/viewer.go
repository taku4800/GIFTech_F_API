package databases

import (
	"yuchami-tinder-app/models"
)

func GetSentLists() ([]models.RemindItemList, error) {
	var lists []models.RemindItemList
	res := DB.Preload("RemindItems").Where("status = ?", "送信済").Find(&lists)
	return lists, res.Error
}

func GetItemByID(id string) (models.RemindItem, error) {
	var item models.RemindItem
	res := DB.First(&item, "id = ?", id)
	return item, res.Error
}
