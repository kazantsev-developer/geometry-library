export interface ShapeData {
  id: string;
  type: string;
  name: string;
  area: number;
  perimeter: number;
  [key: string]: unknown;
}

export interface ShapeUpdateEventDetail {
  target: Shape;
  timestamp: number;
}

/**
 * базовый класс для всех фигур
 * генерирует событие update при изменении состояния
 */
export abstract class Shape extends EventTarget {
  public readonly id: string;
  public readonly type: string;
  private _name: string;

  constructor(type: string, name?: string) {
    super();
    this.id = crypto.randomUUID();
    this.type = type;
    this._name = name || `${type}_${this.id.slice(0, 4)}`;
  }

  public get name(): string {
    return this._name;
  }

  public set name(value: string) {
    this._name = value;
    this.emitUpdate();
  }

  public abstract getArea(): number;
  public abstract getPerimeter(): number;
  public abstract getFormattedDetails(): string;
  public abstract toJSON(): ShapeData;

  protected emitUpdate(): void {
    const detail: ShapeUpdateEventDetail = {
      target: this,
      timestamp: Date.now(),
    };
    this.dispatchEvent(new CustomEvent('update', { detail }));
  }

  protected validatePositive(value: number, fieldName: string): void {
    if (value <= 0) {
      throw new Error(
        `[${this.constructor.name}] ${fieldName} must be positive: ${value}`,
      );
    }
  }
}
