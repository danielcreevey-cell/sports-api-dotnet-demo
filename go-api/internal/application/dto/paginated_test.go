package dto

import "testing"

func TestNewPaginatedList_BasicPagination(t *testing.T) {
	items := []string{"a", "b", "c"}
	result := NewPaginatedList(items, 10, 1, 3)

	if result.TotalCount != 10 {
		t.Errorf("TotalCount: got %d, want 10", result.TotalCount)
	}
	if result.TotalPages != 4 {
		t.Errorf("TotalPages: got %d, want 4", result.TotalPages)
	}
	if result.PageNumber != 1 {
		t.Errorf("PageNumber: got %d, want 1", result.PageNumber)
	}
	if len(result.Items) != 3 {
		t.Errorf("Items count: got %d, want 3", len(result.Items))
	}
}

func TestNewPaginatedList_FirstPage(t *testing.T) {
	result := NewPaginatedList([]int{1, 2}, 10, 1, 2)

	if result.HasPreviousPage {
		t.Error("first page should not have previous page")
	}
	if !result.HasNextPage {
		t.Error("first page of multi-page result should have next page")
	}
}

func TestNewPaginatedList_MiddlePage(t *testing.T) {
	result := NewPaginatedList([]int{3, 4}, 10, 2, 2)

	if !result.HasPreviousPage {
		t.Error("middle page should have previous page")
	}
	if !result.HasNextPage {
		t.Error("middle page should have next page")
	}
}

func TestNewPaginatedList_LastPage(t *testing.T) {
	result := NewPaginatedList([]int{9, 10}, 10, 5, 2)

	if !result.HasPreviousPage {
		t.Error("last page should have previous page")
	}
	if result.HasNextPage {
		t.Error("last page should not have next page")
	}
}

func TestNewPaginatedList_SinglePage(t *testing.T) {
	result := NewPaginatedList([]int{1, 2, 3}, 3, 1, 10)

	if result.HasPreviousPage {
		t.Error("single page should not have previous page")
	}
	if result.HasNextPage {
		t.Error("single page should not have next page")
	}
	if result.TotalPages != 1 {
		t.Errorf("TotalPages: got %d, want 1", result.TotalPages)
	}
}

func TestNewPaginatedList_EmptyItems(t *testing.T) {
	result := NewPaginatedList([]int{}, 0, 1, 10)

	if result.TotalCount != 0 {
		t.Errorf("TotalCount: got %d, want 0", result.TotalCount)
	}
	if result.TotalPages != 0 {
		t.Errorf("TotalPages: got %d, want 0", result.TotalPages)
	}
}

func TestNewPaginatedList_PageJustBeyondTotal(t *testing.T) {
	// Page 6 when TotalPages is 5 → PageNumber (6) <= TotalPages+1 (6) → HasPreviousPage = true
	result := NewPaginatedList([]int{}, 10, 6, 2)
	// TotalPages = ceil(10/2) = 5
	// HasPreviousPage = 6 > 1 && 6 <= 5+1 → true
	if !result.HasPreviousPage {
		t.Error("page just beyond total (TotalPages+1) should have previous page")
	}
	if result.HasNextPage {
		t.Error("page beyond total should not have next page")
	}
}

func TestNewPaginatedList_PageFarBeyondTotal(t *testing.T) {
	// Page 7 when TotalPages is 5 → PageNumber (7) <= TotalPages+1 (6) is false → HasPreviousPage = false
	result := NewPaginatedList([]int{}, 10, 7, 2)

	if result.HasPreviousPage {
		t.Error("page far beyond total should not have previous page")
	}
	if result.HasNextPage {
		t.Error("page far beyond total should not have next page")
	}
}
