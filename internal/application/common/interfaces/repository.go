package interfaces

import (
	"context"

	"github.com/danielcreevey-cell/sports-api-dotnet-demo/internal/application/common/models"
	"github.com/danielcreevey-cell/sports-api-dotnet-demo/internal/domain/entities"
)

// TodoItemBriefDto is a compact projection of a TodoItem for list views.
type TodoItemBriefDto struct {
	ID     int
	ListID int
	Title  string
	Done   bool
}

// TodoListRepository describes persistence operations for TodoList aggregates.
//
// Ported from the TodoLists surface of IApplicationDbContext in
// src/Application/Common/Interfaces/IApplicationDbContext.cs.
type TodoListRepository interface {
	GetAll(ctx context.Context) ([]entities.TodoList, error)
	GetByID(ctx context.Context, id int) (*entities.TodoList, error)
	Create(ctx context.Context, list *entities.TodoList) (int, error)
	Update(ctx context.Context, list *entities.TodoList) error
	Delete(ctx context.Context, id int) error
	PurgeAll(ctx context.Context) error
}

// TodoItemRepository describes persistence operations for TodoItem entities.
//
// Ported from the TodoItems surface of IApplicationDbContext in
// src/Application/Common/Interfaces/IApplicationDbContext.cs.
type TodoItemRepository interface {
	GetByID(ctx context.Context, id int) (*entities.TodoItem, error)
	GetPaginated(ctx context.Context, listID, pageNumber, pageSize int) (*models.PaginatedList[TodoItemBriefDto], error)
	Create(ctx context.Context, item *entities.TodoItem) (int, error)
	Update(ctx context.Context, item *entities.TodoItem) error
	Delete(ctx context.Context, id int) error
}
