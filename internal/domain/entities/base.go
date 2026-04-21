// Package entities contains aggregate roots and entities for the
// CleanArchitecture domain layer.
package entities

import (
	"time"

	"github.com/danielcreevey-cell/sports-api-dotnet-demo/internal/domain/events"
)

// BaseEntity is the common base for all domain entities.
//
// Ported from src/Domain/Common/BaseEntity.cs.
type BaseEntity struct {
	ID int

	domainEvents []events.DomainEvent
}

// DomainEvents returns the events currently queued on this entity.
//
// The returned slice is a snapshot; callers must not mutate the
// underlying storage (equivalent to IReadOnlyCollection<BaseEvent>).
func (b *BaseEntity) DomainEvents() []events.DomainEvent {
	out := make([]events.DomainEvent, len(b.domainEvents))
	copy(out, b.domainEvents)
	return out
}

// AddDomainEvent appends a domain event to the entity's queue.
func (b *BaseEntity) AddDomainEvent(e events.DomainEvent) {
	b.domainEvents = append(b.domainEvents, e)
}

// RemoveDomainEvent removes the first occurrence of e from the queue,
// if present. Included for parity with the .NET BaseEntity.
func (b *BaseEntity) RemoveDomainEvent(e events.DomainEvent) {
	for i, existing := range b.domainEvents {
		if existing == e {
			b.domainEvents = append(b.domainEvents[:i], b.domainEvents[i+1:]...)
			return
		}
	}
}

// ClearDomainEvents removes all queued events.
func (b *BaseEntity) ClearDomainEvents() {
	b.domainEvents = nil
}

// BaseAuditableEntity extends BaseEntity with audit metadata that is
// populated automatically by the infrastructure layer.
//
// Ported from src/Domain/Common/BaseAuditableEntity.cs.
type BaseAuditableEntity struct {
	BaseEntity

	Created        time.Time
	CreatedBy      *string
	LastModified   time.Time
	LastModifiedBy *string
}
