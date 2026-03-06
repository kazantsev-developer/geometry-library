import { Shape, ShapeData } from '../core/Shape';

export interface TriangleData extends ShapeData {
  sides: [number, number, number];
}

/**
 * реализация треугольника по трем сторонам
 */
export class Triangle extends Shape {
  private _a: number;
  private _b: number;
  private _c: number;

  constructor(a: number, b: number, c: number, name?: string) {
    super('triangle', name);
    this.validateSides(a, b, c);
    this._a = a;
    this._b = b;
    this._c = c;
  }

  public get sides(): [number, number, number] {
    return [this._a, this._b, this._c];
  }

  public setSides(a: number, b: number, c: number): void {
    this.validateSides(a, b, c);
    this._a = a;
    this._b = b;
    this._c = c;
    this.emitUpdate();
  }

  /** вычисление площади по формуле Герона */
  public override getArea(): number {
    const s = (this._a + this._b + this._c) / 2;
    return Math.sqrt(s * (s - this._a) * (s - this._b) * (s - this._c));
  }

  public override getPerimeter(): number {
    return this._a + this._b + this._c;
  }

  public override getFormattedDetails(): string {
    return [
      `Стороны: a=${this._a}, b=${this._b}, c=${this._c}`,
      `Периметр: ${this.getPerimeter()}`,
      `Площадь: ${this.getArea().toFixed(2)}`,
    ].join('\n');
  }

  public override toJSON(): TriangleData {
    return {
      id: this.id,
      type: this.type,
      name: this.name,
      sides: [this._a, this._b, this._c],
      area: this.getArea(),
      perimeter: this.getPerimeter(),
    };
  }

  /** проверка существования треугольника и положительных значений */
  private validateSides(a: number, b: number, c: number): void {
    this.validatePositive(a, 'Side A');
    this.validatePositive(b, 'Side B');
    this.validatePositive(c, 'Side C');

    if (a + b <= c || a + c <= b || b + c <= a) {
      throw new Error(
        `[Triangle] Sides ${a}, ${b}, ${c} do not satisfy triangle inequality.`,
      );
    }
  }
}
