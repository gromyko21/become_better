package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalPages(t *testing.T) {
	tests := []struct {
		name          string
		totalRecords  int32
		pageSize      int32
		expectedPages int32
	}{
		{
			name:          "normal case - exact division",
			totalRecords:  100,
			pageSize:      10,
			expectedPages: 10,
		},
		{
			name:          "normal case - with remainder",
			totalRecords:  105,
			pageSize:      10,
			expectedPages: 11, // 105/10 = 10.5 → ceil to 11
		},
		{
			name:          "page size is zero",
			totalRecords:  100,
			pageSize:      0,
			expectedPages: 0,
		},
		{
			name:          "page size is negative",
			totalRecords:  100,
			pageSize:      -5,
			expectedPages: 0,
		},
		{
			name:          "zero records",
			totalRecords:  0,
			pageSize:      10,
			expectedPages: 0,
		},
		{
			name:          "records less than page size",
			totalRecords:  5,
			pageSize:      10,
			expectedPages: 1,
		},
		{
			name:          "both zero",
			totalRecords:  0,
			pageSize:      0,
			expectedPages: 0,
		},
		{
			name:          "records less than page size - edge case",
			totalRecords:  1,
			pageSize:      10,
			expectedPages: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TotalPages(tt.totalRecords, tt.pageSize)
			assert.Equal(t, tt.expectedPages, result)
		})
	}
}
