package models

import "time"

type TmpEntity struct {
	Id        *int64     `json:"id"`
	Title     *string    `json:"title"`
	CreatedAt *time.Time `json:"created_at"`
	IsDeleted *bool      `json:"is_deleted"`
}
