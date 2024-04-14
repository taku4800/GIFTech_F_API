package models

import (
	"time"
)

type RemindItem struct {
	ID        string    `json:"id" gorm:"primary key"`
	ListID    string    `json:"list_id"`
	Order     int       `json:"order"`
	Url       string    `json:"url"`
	Status    string    `json:"status"`
	IsDelete  bool      `json:"is_delete"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RemindItemList struct {
	ID          string       `json:"id" gorm:"primary key"`
	Name        string       `json:"name"`
	Status      string       `json:"status"`
	IsDelete    bool         `json:"is_delete"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	RemindItems []RemindItem `json:"remind_items" gorm:"foreignKey:ListID;references:ID"`
}
