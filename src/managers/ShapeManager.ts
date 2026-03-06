import { Shape } from '../core/Shape';
import { promises as fs } from 'fs';

export interface ShapeEventDetail {
  shape: Shape;
  timestamp: number;
}

/**
 * менеджер для управления коллекциями фигур
 * реализует паттерн observer для отслеживания изменений в фигурах
 */
export class ShapeManager extends EventTarget {
  private readonly _shapes: Map<string, Shape> = new Map();

  /** регистрация фигуры и настройка проброса событий */
  public addShape(shape: Shape): void {
    if (this._shapes.has(shape.id)) return;

    this._shapes.set(shape.id, shape);

    shape.addEventListener('update', () => {
      this.dispatchEvent(new CustomEvent('shape-updated', { detail: shape }));
    });

    this.emitEvent('shape-added', shape);
  }

  public removeShape(id: string): boolean {
    const shape = this._shapes.get(id);
    if (!shape) return false;

    const deleted = this._shapes.delete(id);
    if (deleted) this.emitEvent('shape-removed', shape);
    return deleted;
  }

  public getAllShapes(): Shape[] {
    return Array.from(this._shapes.values());
  }

  /** асинхронная генерация текстового отчета */
  public async saveReport(filePath: string): Promise<void> {
    const shapes = this.getAllShapes();
    let reportData: string[] = [];

    try {
      const existing = await fs.readFile(filePath, 'utf-8');
      if (existing.trim()) {
        reportData = existing.trim().split('\n');
        if (reportData[reportData.length - 1] !== '') reportData.push('');
      }
    } catch {
      reportData = [
        '# Отчет',
        `дата формирования: ${new Date().toLocaleString()}`,
        '='.repeat(40),
        '',
      ];
    }

    const existingCount = reportData.filter((line) =>
      line.match(/^\d+\./),
    ).length;

    shapes.forEach((shape, index) => {
      const typeLabels: Record<string, string> = {
        rectangle: 'прямоугольник',
        circle: 'круг',
        triangle: 'треугольник',
      };

      const typeName = typeLabels[shape.type] || 'фигура';
      const shapeNumber = existingCount + index + 1;

      reportData.push(`${shapeNumber}. [${typeName}] "${shape.name}"`);
      reportData.push('-'.repeat(20));
      reportData.push(shape.getFormattedDetails());
      reportData.push('-'.repeat(40), '');
    });

    await fs.writeFile(filePath, reportData.join('\n'), 'utf-8');
  }

  private emitEvent(type: string, shape: Shape): void {
    const detail: ShapeEventDetail = { shape, timestamp: Date.now() };
    this.dispatchEvent(new CustomEvent(type, { detail }));
  }
}
