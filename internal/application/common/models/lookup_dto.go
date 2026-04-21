package models

// LookupDto is a minimal lookup projection for entities with an identifier
// and a title.
//
// Ported from src/Application/Common/Models/LookupDto.cs.
type LookupDto struct {
	ID    int
	Title string
}
