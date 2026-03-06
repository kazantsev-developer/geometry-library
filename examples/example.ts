import { ShapeManager, Circle, Rectangle, Triangle } from '../dist';
import { unlink } from 'fs/promises';

async function run(): Promise<void> {
  const manager = new ShapeManager();

  manager.addEventListener('shape-updated', (event) => {
    const shape = (event as CustomEvent).detail;
    console.log(`фигура "${shape.name}" обновлена`);
  });

  const circle = new Circle(10, 'круг для дома');
  const rectangle = new Rectangle(5, 12, 'рабочий стол');
  const triangle = new Triangle(3, 4, 5, 'египетский треугольник');

  manager.addShape(circle);
  manager.addShape(rectangle);
  manager.addShape(triangle);

  console.log('добавлено 3 фигуры');

  // console.log('меняем радиус круга');
  // circle.radius = 25;

  try {
    await unlink('examples/report.txt');
  } catch {}

  await manager.saveReport('examples/report.txt');
  console.log('отчет сохранен');
}

run().catch((err) => {
  console.error('Ошибка:', err.message);
});
