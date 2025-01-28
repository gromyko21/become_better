package models

const (
	CategoryWork = 0
	CategoryStudy = 1
	CategoryTraning = 2

	MinuteCategoryType = 1
	CountCategoryType = 2
)

var MainCategoriesMap = map[int32]string{
	CategoryWork: "Работа",
	CategoryStudy: "Учеба",
	CategoryTraning: "Тренировки",
}

var ProgressTypesMap = map[int32]string{
	MinuteCategoryType: "Минуты",
	CountCategoryType: "Количество раз",
}
