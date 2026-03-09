import { Shape, ShapeData } from '../core/Shape';

export interface CircleData extends ShapeData {
  radius: number;
  diameter: number;
  circumference: number;
}

/**
 * класс для работы с кругом
 */
export class Circle extends Shape {
  private _radius: number;

  constructor(radius: number, name?: string) {
    super('circle', name);
    this.validatePositive(radius, 'Radius');
    this._radius = radius;
  }

  public get radius(): number {
    return this._radius;
  }

  public set radius(value: number) {
    this.validatePositive(value, 'Radius');
    this._radius = value;
    this.emitUpdate();
  }

  public override getArea(): number {
    return Math.PI * Math.pow(this._radius, 2);
  }

  public override getPerimeter(): number {
    return 2 * Math.PI * this._radius;
  }

  public getDiameter(): number {
    return this._radius * 2;
  }

  public override getFormattedDetails(): string {
    return [
      `радиус: ${this._radius}`,
      `диаметр: ${this.getDiameter().toFixed(2)}`,
      `площадь: ${this.getArea().toFixed(2)}`,
      `длина окружности: ${this.getPerimeter().toFixed(2)}`,
    ].join('\n');
  }

  public override toJSON(): CircleData {
    const perimeter = this.getPerimeter();
    return {
      id: this.id,
      type: this.type,
      name: this.name,
      radius: this._radius,
      diameter: this.getDiameter(),
      area: this.getArea(),
      perimeter,
      circumference: perimeter,
    };
  }
}
