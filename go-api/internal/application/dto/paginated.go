package dto

import "math"

// PaginatedList is a generic paginated collection.
type PaginatedList[T any] struct {
	Items           []T  `json:"items"`
	PageNumber      int  `json:"pageNumber"`
	TotalPages      int  `json:"totalPages"`
	TotalCount      int  `json:"totalCount"`
	HasPreviousPage bool `json:"hasPreviousPage"`
	HasNextPage     bool `json:"hasNextPage"`
}

// NewPaginatedList creates a new PaginatedList with computed pagination metadata.
func NewPaginatedList[T any](items []T, count, pageNumber, pageSize int) PaginatedList[T] {
	totalPages := int(math.Ceil(float64(count) / float64(pageSize)))
	return PaginatedList[T]{
		Items:           items,
		PageNumber:      pageNumber,
		TotalPages:      totalPages,
		TotalCount:      count,
		HasPreviousPage: pageNumber > 1 && pageNumber <= totalPages+1,
		HasNextPage:     pageNumber < totalPages,
	}
}
