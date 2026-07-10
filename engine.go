package geometry

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Engine orchestrates shape lifecycles and event dispatching.
// It uses a buffered channel for non-blocking event submission and
// a worker goroutine that processes events until the context is canceled.
type Engine struct {
	mu        sync.RWMutex
	shapes    map[ID]Shape
	eventCh   chan ShapeEvent
	workerCtx context.Context
	cancel    context.CancelFunc
	wg        sync.WaitGroup
}

// NewEngine creates an Engine with a given event channel buffer size.
// It does not start the worker automatically; Start must be called.
func NewEngine(bufferSize int) *Engine {
	return &Engine{
		shapes:  make(map[ID]Shape),
		eventCh: make(chan ShapeEvent, bufferSize),
	}
}

// Start launches the background event worker.
// The worker processes events until ctx is canceled or the channel is closed.
// The provided logHandler is called for each event (e.g., for logging or export).
func (e *Engine) Start(ctx context.Context, logHandler func(ShapeEvent)) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.workerCtx != nil {
		return // already started
	}
	e.workerCtx, e.cancel = context.WithCancel(ctx)
	e.wg.Add(1)
	go e.worker(e.workerCtx, logHandler)
}

// Stop gracefully shuts down the event worker and waits for it to finish.
// It cancels the context, which signals the worker to exit.
func (e *Engine) Stop() {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.cancel != nil {
		e.cancel()
	}
	e.wg.Wait()
}

// worker is the main event loop running in a separate goroutine.
func (e *Engine) worker(ctx context.Context, logHandler func(ShapeEvent)) {
	defer e.wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case ev, ok := <-e.eventCh:
			if !ok {
				return
			}
			logHandler(ev)
		}
	}
}

// AddShape thread-safely registers a shape.
// Returns an error if a shape with the same ID already exists.
func (e *Engine) AddShape(s Shape) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	if _, exists := e.shapes[s.GetID()]; exists {
		return fmt.Errorf("shape with ID %s already registered", s.GetID())
	}
	e.shapes[s.GetID()] = s
	return nil
}

// RemoveShape deletes a shape by ID.
// Returns true if the shape was removed, false if it did not exist.
func (e *Engine) RemoveShape(id ID) bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	if _, exists := e.shapes[id]; !exists {
		return false
	}
	delete(e.shapes, id)
	return true
}

// GetShape retrieves a shape by ID in a thread-safe manner.
// Returns the shape and true if found, otherwise nil and false.
func (e *Engine) GetShape(id ID) (Shape, bool) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	s, ok := e.shapes[id]
	return s, ok
}

// AllShapes returns a snapshot of all currently registered shapes.
// The returned slice is a copy to avoid concurrent modification issues.
func (e *Engine) AllShapes() []Shape {
	e.mu.RLock()
	defer e.mu.RUnlock()
	result := make([]Shape, 0, len(e.shapes))
	for _, s := range e.shapes {
		result = append(result, s)
	}
	return result
}

// TriggerUpdate submits an update event for a shape.
// The event is sent to the buffered channel; if the channel is full,
// the event is dropped to avoid blocking the caller.
// This ensures high-throughput performance under heavy load.
func (e *Engine) TriggerUpdate(ctx context.Context, s Shape) {
	select {
	case e.eventCh <- ShapeEvent{
		Ctx:       ctx,
		Target:    s,
		Timestamp: time.Now().UnixNano(),
	}:
	default:
		// Drop event if channel is saturated to prioritize throughput
	}
}
