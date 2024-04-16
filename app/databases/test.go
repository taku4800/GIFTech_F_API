package databases

import (
	"yuchami-tinder-app/models"
)

func GetImages() ([]models.TestImage, error) {
	var images []models.TestImage
	res := DB.Find(&images)
	return images, res.Error
}

func GetImageByID(id string) (models.TestImage, error) {
	var image models.TestImage
	res := DB.First(&image, "id = ?", id)
	return image, res.Error
}

func CreateImage(image models.TestImage) (models.TestImage, error) {
	res := DB.Create(&image)
	return image, res.Error
}
