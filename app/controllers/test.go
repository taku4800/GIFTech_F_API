package controllers

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"net/http"
	"os"
	"yuchami-tinder-app/databases"
	"yuchami-tinder-app/models"

	"github.com/labstack/echo/v4"
)

func GetImages(c echo.Context) error {
	images, err := databases.GetImages()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, images)
}

func GetImage(c echo.Context) error {
	id := c.Param("id")
	image, err := databases.GetImageByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, image)
}

func UploadImage(c echo.Context) error {
	var image models.TestImage
	imgName := c.Param("name")
	imgPath := fmt.Sprintf("../src/images/%s.jpeg", imgName)
	file, err := os.Open(imgPath)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	defer file.Close()
	img, err := jpeg.Decode(file)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	// バッファを用意してバイトスライスを取得する
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, img, nil); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	// 画像を保存
	image.Source = buf.Bytes()
	image, err = databases.CreateImage(image)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusCreated, image)
}
