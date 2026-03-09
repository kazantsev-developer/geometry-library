import { ShapeManager, Circle, Rectangle, Triangle } from '../src';
import { unlink } from 'fs/promises';

async function run(): Promise<void> {
  const manager = new ShapeManager();

  manager.addEventListener('shape-updated', (event) => {
    const shape = (event as CustomEvent).detail;
    console.log(`фигура "${shape.name}" обновлена`);
  });

  const circle = new Circle(9, 'кольцо для девушки');
  const rectangle = new Rectangle(40, 60, 'прикроватный столик');
  const triangle = new Triangle(3, 4, 5, 'египетский треугольник');

  manager.addShape(circle);
  manager.addShape(rectangle);
  manager.addShape(triangle);

  console.log('добавлено 3 фигуры');

  console.log('меняем радиус круга');
  circle.radius = 10;

  try {
    await unlink('examples/report.txt');
  } catch {}

  await manager.saveReport('examples/report.txt');
  console.log('отчет сохранен');
}

run().catch((err) => {
  console.error('Ошибка:', err.message);
});
