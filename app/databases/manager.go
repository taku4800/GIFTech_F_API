package databases

import (
	"log"
	"os"
	"yuchami-tinder-app/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB // グローバルアクセス

func SetupDatabase() {
	godotenv.Load(".env")
	dbUrl := os.Getenv("DATABASE_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	if err = DB.AutoMigrate(&models.RemindItem{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	if err = DB.AutoMigrate(&models.RemindItemList{}); err != nil {
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

func CreateItem(item models.RemindItem) (models.RemindItem, error) {
	res := DB.Create(&item)
	return item, res.Error
}

func CreateList(list models.RemindItemList) (models.RemindItemList, error) {
	res := DB.Create(&list)
	return list, res.Error
}

func UpdateItem(item models.RemindItem) (models.RemindItem, error) {
	res := DB.Save(&item)
	return item, res.Error
}

func UpdateList(list models.RemindItemList) (models.RemindItemList, error) {
	res := DB.Save(&list)
	return list, res.Error
}
