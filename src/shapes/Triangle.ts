import { Shape } from "../core/Shape";

export class Triangle extends Shape {
  private _a: number;
  private _b: number;
  private _c: number;

  constructor(a: number, b: number, c: number, name?: string) {
    super("triangle", name);
    this._validatePositive(a, "Side a");
    this._validatePositive(b, "Side b");
    this._validatePositive(c, "Side c");
    this._validateTriangleExists(a, b, c);
    this._a = a;
    this._b = b;
    this._c = c;
  }

  get sides(): [number, number, number] {
    return [this._a, this._b, this._c];
  }

  get a(): number {
    return this._a;
  }

  get b(): number {
    return this._b;
  }

  get c(): number {
    return this._c;
  }

  setSides(a: number, b: number, c: number): void {
    this._validatePositive(a, "Side a");
    this._validatePositive(b, "Side b");
    this._validatePositive(c, "Side c");
    this._validateTriangleExists(a, b, c);
    this._a = a;
    this._b = b;
    this._c = c;
    this._emitUpdate();
  }

  getArea(): number {
    const s = (this._a + this._b + this._c) / 2;
    return Math.sqrt(s * (s - this._a) * (s - this._b) * (s - this._c));
  }

  getPerimeter(): number {
    return this._a + this._b + this._c;
  }

  toJSON(): Record<string, unknown> {
    return {
      id: this.id,
      type: this.type,
      name: this.name,
      sides: [this._a, this._b, this._c],
      area: this.getArea(),
      perimeter: this.getPerimeter(),
    };
  }

  private _validatePositive(value: number, field: string): void {
    if (value <= 0) {
      throw new Error(`${field} must be positive, got ${value}`);
    }
  }

  private _validateTriangleExists(a: number, b: number, c: number): void {
    if (a + b <= c || a + c <= b || b + c <= a) {
      throw new Error(
        `Invalid triangle: ${a}, ${b}, ${c} do not satisfy triangle inequality`,
      );
    }
  }
}
