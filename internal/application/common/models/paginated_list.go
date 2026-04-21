package models

import "math"

// PaginatedList is a generic paginated result.
//
// Ported from src/Application/Common/Models/PaginatedList.cs.
type PaginatedList[T any] struct {
	Items      []T
	PageNumber int
	TotalPages int
	TotalCount int
}

// NewPaginatedList constructs a PaginatedList, computing TotalPages the same
// way as the C# implementation: Math.Ceiling(count / (double)pageSize).
func NewPaginatedList[T any](items []T, count, pageNumber, pageSize int) *PaginatedList[T] {
	var totalPages int
	if pageSize > 0 {
		totalPages = int(math.Ceil(float64(count) / float64(pageSize)))
	}
	return &PaginatedList[T]{
		Items:      items,
		PageNumber: pageNumber,
		TotalPages: totalPages,
		TotalCount: count,
	}
}

// HasPreviousPage reports whether there is a page before PageNumber.
func (p *PaginatedList[T]) HasPreviousPage() bool {
	return p.PageNumber > 1 && p.PageNumber <= p.TotalPages+1
}

// HasNextPage reports whether there is a page after PageNumber.
func (p *PaginatedList[T]) HasNextPage() bool {
	return p.PageNumber < p.TotalPages
}
