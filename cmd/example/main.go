package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/kazantsev-developer/geometry-library"
)

func main() {
	ctx := context.Background()
	engine := geometry.NewEngine(10)

	engine.Start(ctx, func(ev geometry.ShapeEvent) {
		fmt.Printf("[event] shape %q updated at %d\n", ev.Target.GetName(), ev.Timestamp)
	})
	defer engine.Stop()

	circle, _ := geometry.NewCircle(9.0, "кольцо для девушки")
	rect, _ := geometry.NewRectangle(40, 60, "прикроватный столик")
	triangle, _ := geometry.NewTriangle(3, 4, 5, "египетский треугольник")

	_ = engine.AddShape(circle)
	_ = engine.AddShape(rect)
	_ = engine.AddShape(triangle)
	fmt.Println("добавлено 3 фигуры")

	fmt.Println("меняем радиус круга")
	_ = engine.RemoveShape(circle.GetID())
	newCircle, _ := geometry.NewCircle(10.0, circle.GetName())
	_ = engine.AddShape(newCircle)

	time.Sleep(100 * time.Millisecond)

	if err := saveReport(engine, "report.txt"); err != nil {
		fmt.Printf("ошибка сохранения отчёта: %v\n", err)
	} else {
		fmt.Println("отчет сохранен")
	}
}

func saveReport(e *geometry.Engine, filePath string) error {
	shapes := e.AllShapes()
	if len(shapes) == 0 {
		return fmt.Errorf("no shapes")
	}

	lines := []string{
		"# Отчет",
		fmt.Sprintf("дата формирования: %s", time.Now().Format("1/2/2006, 3:04:05 PM")),
		"",
	}

	typeNames := map[geometry.ShapeType]string{
		geometry.TypeCircle:    "круг",
		geometry.TypeRectangle: "прямоугольник",
		geometry.TypeTriangle:  "треугольник",
	}

	for i, s := range shapes {
		tname := typeNames[s.GetType()]
		lines = append(lines, fmt.Sprintf("%d. [%s] \"%s\"", i+1, tname, s.GetName()))
		lines = append(lines, "-----")

		data := s.ToJSON()
		switch s.GetType() {
		case geometry.TypeCircle:
			lines = append(lines,
				fmt.Sprintf("радиус: %.0f", data.Metadata["radius"]),
				fmt.Sprintf("диаметр: %.2f", data.Metadata["diameter"]),
				fmt.Sprintf("площадь: %.2f", data.Area),
				fmt.Sprintf("длина окружности: %.2f", data.Metadata["circumference"]),
			)
		case geometry.TypeRectangle:
			lines = append(lines,
				fmt.Sprintf("ширина: %.0f", data.Metadata["width"]),
				fmt.Sprintf("высота: %.0f", data.Metadata["height"]),
				fmt.Sprintf("площадь: %.0f", data.Area),
				fmt.Sprintf("периметр: %.0f", data.Perimeter),
			)
		case geometry.TypeTriangle:
			sides := data.Metadata["sides"].([]float64)
			lines = append(lines,
				fmt.Sprintf("Стороны: a=%.0f, b=%.0f, c=%.0f", sides[0], sides[1], sides[2]),
				fmt.Sprintf("Периметр: %.0f", data.Perimeter),
				fmt.Sprintf("Площадь: %.2f", data.Area),
			)
		}
		lines = append(lines, "-----", "")
	}

	return os.WriteFile(filePath, []byte(joinLines(lines)), 0o644)
}

func joinLines(lines []string) string {
	res := ""
	for _, line := range lines {
		res += line + "\n"
	}
	return res
}
