/**
 * Абстрактный класс для геометрических фигур
 * Наследует EventTarget для поддержки событийной модели
 *
 * @emits update - когда изменяются параметры фигуры (имя, размеры)
 * @emits destroy - когда фигура помечена на удаление
 */

export abstract class Shape extends EventTarget {
  readonly id!: string;
  readonly type!: string;
  private _name!: string;

  constructor(type: string, name?: string) {
    super();
    this.id = crypto.randomUUID();
    this.type = type;
    this._name = name || `${type}_${this.id.slice(0, 4)}`;
  }

  get name(): string {
    return this._name;
  }

  set name(value: string) {
    this._name = value;
    this._emitUpdate();
  }

  protected _emitUpdate(): void {
    this.dispatchEvent(new CustomEvent('update', { detail: this }));
  }

  abstract getArea(): number;
  abstract toJSON(): Record<string, unknown>;

  destroy(): void {
    this.dispatchEvent(new CustomEvent('destroy', { detail: { id: this.id } }));
  }
}
