package controllers

import (
	"bytes"
	"fmt"
	"image/png"
	"net/http"
	"os"
	"yuchami-tinder-app/databases"
	"yuchami-tinder-app/models"

	"github.com/labstack/echo/v4"
)

func GetSentLists(c echo.Context) error {
	lists, err := databases.GetSentLists()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, lists)
}

func UpdateItem(c echo.Context) error {
	id := c.Param("id")
	// Request Body
	var input models.RemindItem
	var err error
	if err = c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	// DBにあるItemを検索
	var item models.RemindItem
	if item, err = databases.GetItemByID(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	// itemに更新情報を詰める
	item.Name = input.Name
	item.Order = input.Order
	item.Source = input.Source
	item.Status = input.Status
	item.IsDelete = input.IsDelete
	// itemを更新
	item, err = databases.UpdateItem(item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, item)
}

func DeleteItem(c echo.Context) error {
	id := c.Param("id")
	var item models.RemindItem
	var err error
	if item, err = databases.GetItemByID(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	if err = databases.DeleteItem(item); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("deleted item with id = %s completed.", id)})
}

func SubscribeImage(c echo.Context) error {
	id := c.Param("id")
	// DBにあるItemを検索
	var item models.RemindItem
	var err error
	if item, err = databases.GetItemByID(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	// 指定画像をローカルから開く
	imgName := c.Param("name")
	imgPath := fmt.Sprintf("../src/images/%s.png", imgName)
	file, err := os.Open(imgPath)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	// バッファを用意してバイトスライスを取得する
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	// 画像を保存
	item.Source = buf.Bytes()
	item, err = databases.UpdateItem(item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, item)
}
