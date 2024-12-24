package models

import "github.com/google/uuid"

const (
	CategoryWork = 0
	CategoryStudy = 1
	CategoryTraning = 2
)

var MainCategoriesMap = map[int32]string{
	CategoryWork: "Работа",
	CategoryStudy: "Учеба",
	CategoryTraning: "Тренировки",
}

type Categories struct {
	ID uuid.UUID `db:"id"`
	MainCategory int32 `db:"main_category"`
	Name string `db:"name"`
	Description string `db:"description"`
}

