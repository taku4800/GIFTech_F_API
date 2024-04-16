package models

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
	"gorm.io/gorm"
)

type RemindItem struct {
	ID        string    `json:"id" gorm:"primary key"`
	ListID    string    `json:"list_id"`
	Order     int       `json:"order"`
	Source    []byte    `json:"source" gorm:"size:5000000"`
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

func (r *RemindItem) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == "" {
		entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
		r.ID = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
	}
	return
}

func (r *RemindItemList) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == "" {
		entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
		r.ID = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
	}
	return
}
