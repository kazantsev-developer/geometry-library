# go-geometry-library

Pure Go 2D geometry library with concurrent event engine.

## Features

- Circle, Rectangle, Triangle with strict validation
- Thread-safe engine with buffered event channel
- Context-controlled worker goroutine
- Zero-allocation area/perimeter calculations
- JSON serialization for all shapes
- Benchmarks with memory profiling

## Files

- `geometry.go` — types, interfaces, UUID generator
- `shapes.go` — Circle, Rectangle, Triangle implementations
- `engine.go` — storage, event channel, worker lifecycle
- `geometry_test.go` — unit tests and benchmarks

## Quick start

```go
engine := NewEngine(10)
engine.Start(ctx, handler)
defer engine.Stop()

circle, _ := NewCircle(5, "name")
engine.AddShape(circle)
engine.TriggerUpdate(ctx, circle)
```
