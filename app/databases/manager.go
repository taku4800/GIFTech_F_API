package databases

import (
	"log"
	"yuchami-app-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB // グローバルアクセス

func SetupDatabase() {
	connStr := "host=localhost user=yuchami password=yuchami0908 dbname=yuchami_api port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	var err error
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	if err = DB.AutoMigrate(&models.RemindItem{}, &models.RemindItemList{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}

func GetLists() ([]models.RemindItemList, error) {
	var lists []models.RemindItemList
	res := DB.Find(&lists)
	return lists, res.Error
}

func GetListByID(id string) (models.RemindItemList, error) {
	var list models.RemindItemList
	res := DB.Preload("RemindItems").First(&list, "id = ?", id)
	return list, res.Error
}

func CreateItem(item models.RemindItem) error {
	res := DB.Create(item)
	return res.Error
}

func CreateList(list models.RemindItemList) error {
	res := DB.Create(list)
	return res.Error
}
