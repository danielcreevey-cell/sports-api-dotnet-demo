package interfaces

import (
	"context"

	"github.com/danielcreevey-cell/sports-api-dotnet-demo/internal/application/common/models"
)

// IdentityService describes identity-related operations (user lookup,
// authorization, user CRUD).
//
// Ported from src/Application/Common/Interfaces/IIdentityService.cs.
type IdentityService interface {
	GetUserName(ctx context.Context, userID string) (*string, error)
	IsInRole(ctx context.Context, userID string, role string) (bool, error)
	Authorize(ctx context.Context, userID string, policyName string) (bool, error)
	CreateUser(ctx context.Context, userName string, password string) (models.Result, string, error)
	DeleteUser(ctx context.Context, userID string) (models.Result, error)
}
