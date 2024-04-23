package controllers

import (
	"fmt"
	"net/http"
	"yuchami-tinder-app/databases"
	"yuchami-tinder-app/models"

	"github.com/labstack/echo/v4"
)

func GetLists(c echo.Context) error {
	lists, err := databases.GetLists()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, lists)
}

func GetList(c echo.Context) error {
	id := c.Param("id")
	list, err := databases.GetListByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, list)
}

func CreateList(c echo.Context) error {
	var list models.RemindItemList
	if err := c.Bind(&list); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	var err error
	list, err = databases.CreateList(list)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusCreated, list)
}

func UpdateList(c echo.Context) error {
	id := c.Param("id")
	// Request Body
	var input models.RemindItemList
	var err error
	if err = c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	// DBにあるListを取得
	var list models.RemindItemList
	if list, err = databases.GetListByID(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	// RemindItemsに対し、Statusを"送信済"に変更
	for _, item := range list.RemindItems {
		item.Status = "送信済"
		if _, err := databases.UpdateItem(item); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}
	}
	// Listに更新情報を詰める
	list.Name = input.Name
	list.Status = input.Status
	list.IsDelete = input.IsDelete
	// Listを更新
	list, err = databases.UpdateList(list)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	// 自分以外のListのStatusを"アーカイブ"に変更
	var lists []models.RemindItemList
	if lists, err = databases.GetListsExcludingID(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	for _, l := range lists {
		l.Status = "アーカイブ"
		if _, err := databases.UpdateList(l); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}
	}
	// 最新のListの情報をとってくる
	if list, err = databases.GetListByID(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, list)
}

func DeleteList(c echo.Context) error {
	id := c.Param("id")
	// DBにあるListを取得
	var list models.RemindItemList
	var err error
	if list, err = databases.GetListByID(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	// RemindItemsを削除
	for _, item := range list.RemindItems {
		if err := databases.DeleteItem(item); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}
	}
	// Listを削除
	err = databases.DeleteList(list)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("deleted list with id = %s completed.", id)})
}
