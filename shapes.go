package geometry

import (
	"fmt"
	"math"
)

// Circle implements the Shape interface for 2D circles.
// It stores radius and provides area, perimeter, and diameter via metadata.
type Circle struct {
	id     ID
	stype  ShapeType
	name   string
	radius float64
}

// NewCircle instantiates a Circle with positive radius validation.
// If name is empty, a generated UUID prefix is used as a fallback.
func NewCircle(radius float64, name string) (*Circle, error) {
	if radius <= 0 {
		return nil, fmt.Errorf("[Circle] radius must be positive: %f", radius)
	}
	if name == "" {
		name = fmt.Sprintf("circle_%s", generateUUID()[:4])
	}
	return &Circle{
		id:     ID(generateUUID()),
		stype:  TypeCircle,
		name:   name,
		radius: radius,
	}, nil
}

// GetID returns the unique identifier of the circle.
func (c *Circle) GetID() ID { return c.id }

// GetType returns the shape type (circle).
func (c *Circle) GetType() ShapeType { return c.stype }

// GetName returns the human-readable name of the circle.
func (c *Circle) GetName() string { return c.name }

// SetName updates the circle's name.
func (c *Circle) SetName(name string) { c.name = name }

// GetArea computes the area πr².
func (c *Circle) GetArea() float64 { return math.Pi * c.radius * c.radius }

// GetPerimeter computes the circumference 2πr.
func (c *Circle) GetPerimeter() float64 { return 2 * math.Pi * c.radius }

// ToJSON returns a ShapeData struct with all computed fields and circle-specific metadata.
func (c *Circle) ToJSON() ShapeData {
	return ShapeData{
		ID:        c.id,
		Type:      c.stype,
		Name:      c.name,
		Area:      c.GetArea(),
		Perimeter: c.GetPerimeter(),
		Metadata: map[string]interface{}{
			"radius":        c.radius,
			"diameter":      c.radius * 2,
			"circumference": c.GetPerimeter(),
		},
	}
}

// Rectangle implements the Shape interface for 2D rectangles.
// It stores width and height.
type Rectangle struct {
	id     ID
	stype  ShapeType
	name   string
	width  float64
	height float64
}

// NewRectangle instantiates a Rectangle with positive dimension validation.
func NewRectangle(width, height float64, name string) (*Rectangle, error) {
	if width <= 0 || height <= 0 {
		return nil, fmt.Errorf("[Rectangle] dimensions must be positive: w=%f, h=%f", width, height)
	}
	if name == "" {
		name = fmt.Sprintf("rectangle_%s", generateUUID()[:4])
	}
	return &Rectangle{
		id:     ID(generateUUID()),
		stype:  TypeRectangle,
		name:   name,
		width:  width,
		height: height,
	}, nil
}

// GetID returns the unique identifier.
func (r *Rectangle) GetID() ID { return r.id }

// GetType returns the shape type (rectangle).
func (r *Rectangle) GetType() ShapeType { return r.stype }

// GetName returns the rectangle's name.
func (r *Rectangle) GetName() string { return r.name }

// SetName updates the rectangle's name.
func (r *Rectangle) SetName(name string) { r.name = name }

// GetArea computes width * height.
func (r *Rectangle) GetArea() float64 { return r.width * r.height }

// GetPerimeter computes 2*(width+height).
func (r *Rectangle) GetPerimeter() float64 { return 2 * (r.width + r.height) }

// ToJSON returns ShapeData with rectangle-specific metadata.
func (r *Rectangle) ToJSON() ShapeData {
	return ShapeData{
		ID:        r.id,
		Type:      r.stype,
		Name:      r.name,
		Area:      r.GetArea(),
		Perimeter: r.GetPerimeter(),
		Metadata: map[string]interface{}{
			"width":  r.width,
			"height": r.height,
		},
	}
}

// Triangle implements the Shape interface using three sides.
// It validates the triangle inequality strictly.
type Triangle struct {
	id    ID
	stype ShapeType
	name  string
	a     float64
	b     float64
	c     float64
}

// NewTriangle instantiates a Triangle with side validation (positive and inequality).
func NewTriangle(a, b, c float64, name string) (*Triangle, error) {
	if a <= 0 || b <= 0 || c <= 0 {
		return nil, fmt.Errorf("[Triangle] sides must be positive: %f, %f, %f", a, b, c)
	}
	if a+b <= c || a+c <= b || b+c <= a {
		return nil, fmt.Errorf("[Triangle] sides %f, %f, %f do not satisfy triangle inequality", a, b, c)
	}
	if name == "" {
		name = fmt.Sprintf("triangle_%s", generateUUID()[:4])
	}
	return &Triangle{
		id:    ID(generateUUID()),
		stype: TypeTriangle,
		name:  name,
		a:     a,
		b:     b,
		c:     c,
	}, nil
}

// GetID returns the unique identifier.
func (t *Triangle) GetID() ID { return t.id }

// GetType returns the shape type (triangle).
func (t *Triangle) GetType() ShapeType { return t.stype }

// GetName returns the triangle's name.
func (t *Triangle) GetName() string { return t.name }

// SetName updates the triangle's name.
func (t *Triangle) SetName(name string) { t.name = name }

// GetArea computes area using Heron's formula.
func (t *Triangle) GetArea() float64 {
	s := (t.a + t.b + t.c) / 2
	return math.Sqrt(s * (s - t.a) * (s - t.b) * (s - t.c))
}

// GetPerimeter returns the sum of sides.
func (t *Triangle) GetPerimeter() float64 { return t.a + t.b + t.c }

// ToJSON returns ShapeData with triangle-specific metadata (sides slice).
func (t *Triangle) ToJSON() ShapeData {
	return ShapeData{
		ID:        t.id,
		Type:      t.stype,
		Name:      t.name,
		Area:      t.GetArea(),
		Perimeter: t.GetPerimeter(),
		Metadata: map[string]interface{}{
			"sides": []float64{t.a, t.b, t.c},
		},
	}
}
