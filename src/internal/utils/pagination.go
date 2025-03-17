package utils

import "math"

func TotalPages(totalRecords, pageSize int32) int32 {
	if pageSize <= 0 {
		return 0 // защита от деления на 0
	}
	return int32(math.Ceil(float64(totalRecords) / float64(pageSize)))
}
