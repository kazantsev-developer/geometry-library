import { Shape } from '../core/Shape';

export class Rectangle extends Shape {
  private _width: number;
  private _height: number;

  constructor(width: number, height: number, name?: string) {
    super('rectangle', name);
    this._validatePositive(width, 'Width');
    this._validatePositive(height, 'Height');
    this._width = width;
    this._height = height;
  }

  get width(): number {
    return this._width;
  }

  set width(value: number) {
    this._validatePositive(value, 'Width');
    this._width = value;
    this._emitUpdate();
  }

  get height(): number {
    return this._height;
  }

  set height(value: number) {
    this._validatePositive(value, 'Height');
    this._height = value;
    this._emitUpdate();
  }

  getArea(): number {
    return this._width * this._height;
  }

  getPerimeter(): number {
    return 2 * (this._width + this._height);
  }

  toJSON(): Record<string, unknown> {
    return {
      id: this.id,
      type: this.type,
      name: this.name,
      width: this._width,
      height: this._height,
      area: this.getArea(),
      perimeter: this.getPerimeter(),
    };
  }

  private _validatePositive(value: number, field: string): void {
    if (value <= 0) {
      throw new Error(`${field} must be positive, got ${value}`);
    }
  }
}
