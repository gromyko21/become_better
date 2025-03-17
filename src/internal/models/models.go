package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID           uuid.UUID `db:"id"`
	MainCategory int32     `db:"main_category"`
	Name         string    `db:"name"`
	Description  string    `db:"description"`
	ProgressType int32     `db:"progress_type"`
}

type Progress struct {
	ID           uuid.UUID `db:"id"`
	UserID       uuid.UUID `db:"user_id"`
	CategoryID   uuid.UUID `db:"category_id"`
	ProgressType int32     `db:"progress_type"`
	Result       int32     `db:"result_int"`
	Description  string    `db:"result_description"`
	Date         time.Time `db:"date"`
}
