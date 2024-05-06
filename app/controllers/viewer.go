package controllers

import (
	"fmt"
	"io"
	"net/http"
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
	// ファイルをリクエストから取得
	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	defer src.Close()
	// ファイルの内容を読み込む
	fileData, err := io.ReadAll(src)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	// 画像を登録
	item.Source = fileData
	item, err = databases.UpdateItem(item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, item)
}
