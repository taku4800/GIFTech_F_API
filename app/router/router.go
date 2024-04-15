package router

import (
	"yuchami-tinder-app/controllers"
	"yuchami-tinder-app/databases"

	"github.com/labstack/echo/v4"
)

func ActivateRouter() {
	e := echo.New()
	databases.SetupDatabase()

	e.GET("/manager/remindItemLists", controllers.GetLists)
	e.GET("/manager/remindItemLists/:id", controllers.GetList)
	e.POST("/manager/remindItemLists", controllers.CreateList)
	e.PATCH("/manager/remindItemLists/:id", controllers.UpdateList)
	e.DELETE("/manager/remindItemLists/:id", controllers.DeleteList)

	e.Logger.Fatal(e.Start(":8080"))
}
