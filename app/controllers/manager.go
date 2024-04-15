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
