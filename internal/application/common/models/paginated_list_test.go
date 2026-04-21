package models_test

import (
	"testing"

	"github.com/danielcreevey-cell/sports-api-dotnet-demo/internal/application/common/models"
)

func TestHasPreviousPage_ShouldBeFalse_WhenOnFirstPage(t *testing.T) {
	list := models.NewPaginatedList[int]([]int{}, 100, 1, 10)

	if list.HasPreviousPage() {
		t.Fatalf("expected HasPreviousPage to be false on first page")
	}
}

func TestHasPreviousPage_ShouldBeTrue_WhenOnSecondPage(t *testing.T) {
	list := models.NewPaginatedList[int]([]int{}, 100, 2, 10)

	if !list.HasPreviousPage() {
		t.Fatalf("expected HasPreviousPage to be true on second page")
	}
}

func TestHasPreviousPage_ShouldBeTrue_WhenOnePageBeyondLastPage(t *testing.T) {
	list := models.NewPaginatedList[int]([]int{}, 100, 11, 10)

	if !list.HasPreviousPage() {
		t.Fatalf("expected HasPreviousPage to be true one page beyond last page")
	}
}

func TestHasPreviousPage_ShouldBeFalse_WhenTwoPagesOrMoreBeyondLastPage(t *testing.T) {
	list := models.NewPaginatedList[int]([]int{}, 100, 12, 10)

	if list.HasPreviousPage() {
		t.Fatalf("expected HasPreviousPage to be false two pages beyond last page")
	}
}

func TestHasNextPage_ShouldBeFalse_WhenOnLastPage(t *testing.T) {
	list := models.NewPaginatedList[int]([]int{}, 100, 10, 10)

	if list.HasNextPage() {
		t.Fatalf("expected HasNextPage to be false on last page")
	}
}

func TestHasNextPage_ShouldBeTrue_WhenNotOnLastPage(t *testing.T) {
	list := models.NewPaginatedList[int]([]int{}, 100, 9, 10)

	if !list.HasNextPage() {
		t.Fatalf("expected HasNextPage to be true when not on last page")
	}
}
