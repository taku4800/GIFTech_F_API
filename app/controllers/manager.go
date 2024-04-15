package controllers

import (
	"net/http"
	"yuchami-app-api/databases"
	"yuchami-app-api/models"

	"github.com/labstack/echo/v4"
)

func GetLists(c echo.Context) error {
	lists, err := databases.GetLists()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, lists)
}

func GetList(c echo.Context) error {
	id := c.Param("id")
	list, err := databases.GetListByID(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, list)
}

func CreateList(c echo.Context) error {
	var list models.RemindItemList
	if err := c.Bind(&list); err != nil {
		return err
	}
	var err error
	list, err = databases.CreateList(list)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, list)
}

func UpdateList(c echo.Context) error {
	id := c.Param("id")
	// Request Body
	var input models.RemindItemList
	var err error
	if err = c.Bind(&input); err != nil {
		return err
	}
	// DBにあるListを取得
	var list models.RemindItemList
	if list, err = databases.GetListByID(id); err != nil {
		return err
	}
	// 既存のRemindItemsに対し、is_delete=trueに更新
	for _, item := range list.RemindItems {
		item.IsDelete = true
		if _, err := databases.UpdateItem(item); err != nil {
			return err
		}
	}
	// Listに更新情報を詰める
	list.ID = id
	list.Name = input.Name
	list.Status = input.Status
	list.IsDelete = input.IsDelete
	// inputからのItemsを追加
	list.RemindItems = nil
	for _, item := range input.RemindItems {
		newItem := models.RemindItem{
			ListID:   list.ID,
			Order:    item.Order,
			Url:      item.Url,
			Status:   item.Status,
			IsDelete: item.IsDelete,
		}
		list.RemindItems = append(list.RemindItems, newItem)
	}
	// Listを更新
	list, err = databases.UpdateList(list)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, list)
}

func DeleteList(c echo.Context) error {
	id := c.Param("id")
	// DBにあるListを取得
	var list models.RemindItemList
	var err error
	if list, err = databases.GetListByID(id); err != nil {
		return err
	}
	// 既存のRemindItemsに対し、is_delete=trueに更新
	for _, item := range list.RemindItems {
		item.IsDelete = true
		if _, err := databases.UpdateItem(item); err != nil {
			return err
		}
	}
	// Listを更新
	list.IsDelete = true
	list, err = databases.UpdateList(list)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, list)
}
