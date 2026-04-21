package entities

import (
	"github.com/danielcreevey-cell/sports-api-dotnet-demo/internal/domain/valueobjects"
)

// TodoList is the aggregate root grouping TodoItem entries.
//
// Ported from src/Domain/Entities/TodoList.cs.
type TodoList struct {
	BaseAuditableEntity

	Title  *string
	Colour valueobjects.Colour
	// Items are held as pointers so that mutations made through a
	// *TodoItem (for example SetDone, which appends a domain event)
	// are visible to other holders of the same item.
	Items []*TodoItem
}

// NewTodoList constructs a TodoList with the default Colour (White) and
// an empty Items slice, matching the .NET default field initializers.
func NewTodoList() *TodoList {
	return &TodoList{
		Colour: valueobjects.ColourWhite,
		Items:  []*TodoItem{},
	}
}
