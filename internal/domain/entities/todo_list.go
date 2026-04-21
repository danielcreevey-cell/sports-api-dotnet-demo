package entities

// Colour represents the colour of a TodoList.
type Colour string

const (
	ColourWhite  Colour = "White"
	ColourRed    Colour = "Red"
	ColourOrange Colour = "Orange"
	ColourYellow Colour = "Yellow"
	ColourGreen  Colour = "Green"
	ColourBlue   Colour = "Blue"
	ColourPurple Colour = "Purple"
	ColourGrey   Colour = "Grey"
)

// TodoList is the Go port of the Domain TodoList entity.
type TodoList struct {
	ID     int
	Title  *string
	Colour Colour
	Items  []TodoItem
}
