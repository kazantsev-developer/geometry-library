import { Shape } from "../core/Shape";

export class ShapeManager extends EventTarget {
  private _shapes: Map<string, Shape> = new Map();

  addShape(shape: Shape): void {
    this._shapes.set(shape.id, shape);

    shape.addEventListener("update", () => {
      this.dispatchEvent(new CustomEvent("shape-updated", { detail: shape }));
    });

    this.dispatchEvent(new CustomEvent("shape-added", { detail: shape }));
  }

  removeShape(id: string): boolean {
    const deleted = this._shapes.delete(id);

    if (deleted) {
      this.dispatchEvent(new CustomEvent("shape-removed", { detail: { id } }));
    }

    return deleted;
  }

  getShape(id: string): Shape | undefined {
    return this._shapes.get(id);
  }

  getAllShapes(): Shape[] {
    return Array.from(this._shapes.values());
  }

  getByTypes<T extends Shape>(type: string): T[] {
    return this.getAllShapes().filter((s) => s.type === type) as T[];
  }

  clear(): void {
    this._shapes.clear();
    this.dispatchEvent(new CustomEvent("shapes-cleared"));
  }

  get count(): number {
    return this._shapes.size;
  }

  async saveToFile(filename: string): Promise<void> {
    return new Promise((resolve) => {
      setTimeout(() => {
        const data = JSON.stringify(
          this.getAllShapes().map((s) => s.toJSON()),
          null,
          2,
        );
        console.log(`Saving to ${filename}:`, data);
        resolve();
      }, 1000);
    });
  }
}
