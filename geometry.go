// Package geometry provides a high-performance 2D computational geometry engine
// with concurrency support and event-driven updates.
package geometry

import (
	"context"
	"crypto/rand"
	"fmt"
)

// ID represents a unique identifier for a shape, stored as a UUID string.
type ID string

// ShapeType enumerates the supported geometric primitives.
type ShapeType string

const (
	TypeCircle    ShapeType = "circle"
	TypeRectangle ShapeType = "rectangle"
	TypeTriangle  ShapeType = "triangle"
)

// Shape defines the contract for all geometric figures.
// All methods are designed to be safe for concurrent read operations;
// write operations (SetName) are expected to be serialized by the caller.
type Shape interface {
	GetID() ID
	GetType() ShapeType
	GetName() string
	SetName(name string)
	GetArea() float64
	GetPerimeter() float64
	// ToJSON returns a serializable representation for reporting and exporting.
	ToJSON() ShapeData
}

// ShapeData is a plain data container used for JSON marshaling and reporting.
// It includes computed properties (Area, Perimeter) and a metadata map for
// shape-specific fields (e.g., radius, sides).
type ShapeData struct {
	ID        ID                     `json:"id"`
	Type      ShapeType              `json:"type"`
	Name      string                 `json:"name"`
	Area      float64                `json:"area"`
	Perimeter float64                `json:"perimeter"`
	Metadata  map[string]interface{} `json:"metadata"`
}

// ShapeEvent represents a state change notification for a shape.
// It carries a context for cancellation/deadline propagation, the target shape,
// and a high-resolution timestamp (nanoseconds) for ordering and telemetry.
// The Error field is used to propagate validation failures during updates.
type ShapeEvent struct {
	Ctx       context.Context `json:"-"`
	Target    Shape           `json:"target"`
	Timestamp int64           `json:"timestamp"` // Unix nano
	Error     error           `json:"error,omitempty"`
}

// generateUUID creates a version 4‑like UUID using crypto/rand.
// It is used as a fallback identifier when the caller does not provide a name.
// The function is defined here to keep all ID generation logic centralized.
func generateUUID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
