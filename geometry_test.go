package geometry

import (
	"context"
	"sync"
	"testing"
	"time"
)

// BenchmarkCircleArea measures allocation-free area computation for Circle.
func BenchmarkCircleArea(b *testing.B) {
	circle, _ := NewCircle(10.0, "bench_circle")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = circle.GetArea()
	}
}

// BenchmarkTriangleArea measures allocation-free area computation using Heron's formula.
func BenchmarkTriangleArea(b *testing.B) {
	triangle, _ := NewTriangle(3.0, 4.0, 5.0, "bench_triangle")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = triangle.GetArea()
	}
}

// BenchmarkEngineConcurrentThroughput evaluates engine performance under parallel updates.
func BenchmarkEngineConcurrentThroughput(b *testing.B) {
	engine := NewEngine(1000)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start worker with a no-op handler to simulate real processing
	engine.Start(ctx, func(ev ShapeEvent) { /* no-op */ })
	defer engine.Stop()

	circle, _ := NewCircle(15.5, "shared_circle")
	_ = engine.AddShape(circle)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			engine.TriggerUpdate(ctx, circle)
		}
	})
}

// BenchmarkEngineAllShapes measures allocation cost of taking a snapshot of all shapes.
func BenchmarkEngineAllShapes(b *testing.B) {
	engine := NewEngine(100)
	for i := 0; i < 100; i++ {
		c, _ := NewCircle(float64(i+1), "circle")
		_ = engine.AddShape(c)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = engine.AllShapes()
	}
}

// TestShapeValidation ensures that invalid shapes are rejected by constructors.
func TestShapeValidation(t *testing.T) {
	_, err := NewTriangle(1.0, 2.0, 10.0, "invalid")
	if err == nil {
		t.Error("expected validation error for broken triangle inequality")
	}

	_, err = NewCircle(-5.0, "invalid")
	if err == nil {
		t.Error("expected validation error for negative radius")
	}

	_, err = NewRectangle(0, 10, "invalid")
	if err == nil {
		t.Error("expected validation error for zero width")
	}
}

// TestEngineAddDuplicate checks that adding a duplicate shape ID returns an error.
func TestEngineAddDuplicate(t *testing.T) {
	engine := NewEngine(10)
	c1, _ := NewCircle(1.0, "c1")
	_ = engine.AddShape(c1)
	err := engine.AddShape(c1)
	if err == nil {
		t.Error("expected duplicate error")
	}
}

// TestEngineEventProcessing verifies that events are dispatched to the handler.
func TestEngineEventProcessing(t *testing.T) {
	engine := NewEngine(10)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1) // we expect at least one event

	handler := func(ev ShapeEvent) {
		wg.Done()
	}

	engine.Start(ctx, handler)
	defer engine.Stop()

	circle, _ := NewCircle(5.0, "test_circle")
	_ = engine.AddShape(circle)
	engine.TriggerUpdate(ctx, circle)

	// Wait for the event to be processed or timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	case <-done:
		// all good
	case <-time.After(2 * time.Second):
		t.Error("event not processed in time")
	}
}

// TestEngineAllShapesConsistency checks that AllShapes returns a consistent snapshot.
func TestEngineAllShapesConsistency(t *testing.T) {
	engine := NewEngine(10)
	circle, _ := NewCircle(1.0, "c1")
	rectangle, _ := NewRectangle(2, 3, "r1")
	_ = engine.AddShape(circle)
	_ = engine.AddShape(rectangle)

	shapes := engine.AllShapes()
	if len(shapes) != 2 {
		t.Errorf("expected 2 shapes, got %d", len(shapes))
	}
	// Verify that IDs match
	idMap := make(map[ID]bool)
	for _, s := range shapes {
		idMap[s.GetID()] = true
	}
	if !idMap[circle.GetID()] || !idMap[rectangle.GetID()] {
		t.Error("AllShapes did not contain all shapes")
	}
}
