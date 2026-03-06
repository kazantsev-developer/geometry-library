import { ShapeManager, Circle, Rectangle, Triangle } from '../src';

async function run(): Promise<void> {
  const manager = new ShapeManager();

  const circle = new Circle(10, 'круг для дома');
  const rectangle = new Rectangle(5, 12, 'рабочий стол');
  const triangle = new Triangle(3, 4, 5, 'египетский треугольник');

  manager.addShape(circle);
  manager.addShape(rectangle);
  manager.addShape(triangle);

  manager.addEventListener('shape-updated', (event) => {
    const shape = (event as CustomEvent).detail;
    console.log(`состояние фигуры "${shape.name}" обновлено.`);
  });

  circle.radius = 25;

  try {
    await manager.saveReport('examples/report.txt');
    console.log('отчет успешно сформирован: examples/report.txt');
  } catch (error) {
    console.error(`не удалось сохранить отчет: ${(error as Error).message}`);
  }
}

run().catch((err) => {
  console.error('Fatal execution error:', err);
  process.exit(1);
});
