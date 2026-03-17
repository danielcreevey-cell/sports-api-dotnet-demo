package domain

import "time"

// BaseEntity provides ID and domain event tracking.
type BaseEntity struct {
	ID           int
	domainEvents []DomainEvent
}

// AddDomainEvent adds a domain event to the entity.
func (e *BaseEntity) AddDomainEvent(event DomainEvent) {
	e.domainEvents = append(e.domainEvents, event)
}

// RemoveDomainEvent removes a domain event from the entity by reference.
func (e *BaseEntity) RemoveDomainEvent(event DomainEvent) {
	for i, ev := range e.domainEvents {
		if ev == event {
			e.domainEvents = append(e.domainEvents[:i], e.domainEvents[i+1:]...)
			return
		}
	}
}

// ClearDomainEvents removes all domain events from the entity.
func (e *BaseEntity) ClearDomainEvents() {
	e.domainEvents = nil
}

// DomainEvents returns the list of domain events.
func (e *BaseEntity) DomainEvents() []DomainEvent {
	return e.domainEvents
}

// BaseAuditableEntity adds audit tracking fields.
type BaseAuditableEntity struct {
	BaseEntity
	Created        time.Time
	CreatedBy      string
	LastModified   time.Time
	LastModifiedBy string
}
