import { Shape } from '../core/Shape';

export class Circle extends Shape {
  private _radius: number;

  constructor(radius: number, name?: string) {
    super('circle', name);
    this._validatePositive(radius, 'Radius');
    this._radius = radius;
  }

  get radius(): number {
    return this._radius;
  }

  set radius(value: number) {
    this._validatePositive(value, 'Radius');
    this._radius = value;
    this._emitUpdate();
  }

  getArea(): number {
    return Math.PI * this._radius ** 2;
  }

  getDiameter(): number {
    return this._radius * 2;
  }

  getCircumference(): number {
    return 2 * Math.PI * this._radius;
  }

  toJSON(): Record<string, unknown> {
    return {
      id: this.id,
      type: this.type,
      name: this.name,
      radius: this._radius,
      diameter: this.getDiameter(),
      area: this.getArea(),
      circumference: this.getCircumference(),
    };
  }

  private _validatePositive(value: number, field: string): void {
    if (value <= 0) {
      throw new Error(`${field} must be positive, got ${value}`);
    }
  }
}
