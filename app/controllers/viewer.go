package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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

func SendNotification(c echo.Context) error {
	user := c.Param("user")
	userUpper := strings.ToUpper(user)
	sendTo := os.Getenv(userUpper)
	// リクエストするJSONデータ
	data := []map[string]string{
		{
			"to":    sendTo,
			"sound": "default",
			"title": "モッテンダー",
			"body":  "マネさんから忘れ物リストが届きました！",
		},
	}
	// JSONデータをバイト配列にエンコード
	jsonData, err := json.Marshal(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	// HTTP POSTリクエストを作成
	url := "https://exp.host/--/api/v2/push/send"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	// ヘッダーを設定
	req.Header.Set("Content-Type", "application/json")
	// HTTPクライアントを作成してリクエストを送信
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	defer resp.Body.Close()
	// レスポンスボディを読み出し
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	// レスポンスボディを返却
	return c.JSON(resp.StatusCode, string(body))
}
