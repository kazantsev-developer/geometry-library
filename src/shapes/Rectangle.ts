import { Shape, ShapeData } from '../core/Shape';

export interface RectangleData extends ShapeData {
  width: number;
  height: number;
}

/**
 * класс для работы с прямоугольником
 */
export class Rectangle extends Shape {
  private _width: number;
  private _height: number;

  constructor(width: number, height: number, name?: string) {
    super('rectangle', name);
    this.validatePositive(width, 'Width');
    this.validatePositive(height, 'Height');
    this._width = width;
    this._height = height;
  }

  public get width(): number {
    return this._width;
  }

  public set width(value: number) {
    this.validatePositive(value, 'Width');
    this._width = value;
    this.emitUpdate();
  }

  public get height(): number {
    return this._height;
  }

  public set height(value: number) {
    this.validatePositive(value, 'Height');
    this._height = value;
    this.emitUpdate();
  }

  public override getArea(): number {
    return this._width * this._height;
  }

  public override getPerimeter(): number {
    return 2 * (this._width + this._height);
  }

  public override getFormattedDetails(): string {
    return [
      `ширина: ${this._width}`,
      `высота: ${this._height}`,
      `площадь: ${this.getArea()}`,
      `периметр: ${this.getPerimeter()}`,
    ].join('\n');
  }

  public override toJSON(): RectangleData {
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
}
