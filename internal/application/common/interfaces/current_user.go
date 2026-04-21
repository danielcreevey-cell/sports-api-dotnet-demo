package interfaces

// CurrentUser exposes identity information about the caller of the current
// request.
//
// Ported from src/Application/Common/Interfaces/IUser.cs.
type CurrentUser interface {
	GetID() *string
	GetRoles() []string
}
