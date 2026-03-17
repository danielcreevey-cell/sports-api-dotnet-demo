package dto

// LookupDto is a generic key-value pair for dropdown lists.
type LookupDto struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
