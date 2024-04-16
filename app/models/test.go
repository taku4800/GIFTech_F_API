package models

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
	"gorm.io/gorm"
)

type TestImage struct {
	ID        string    `json:"id" gorm:"primary key"`
	Source    []byte    `json:"image" gorm:"size:5000000"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (r *TestImage) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == "" {
		entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
		r.ID = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
	}
	return
}
