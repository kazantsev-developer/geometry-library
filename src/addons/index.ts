import { Triangle } from '../shapes/Triangle';

interface NativeValidator {
  validateTriangle: (a: number, b: number, c: number) => boolean;
}

let nativeValidator: NativeValidator | null = null;

try {
  const addon = require('../../build/Release/validator.node');
  if (addon && typeof addon.validateTriangle === 'function') {
    nativeValidator = addon as NativeValidator;
    console.log('Native C++ validator loaded successfully');
  } else {
    console.log(
      'Native validator has incorrect interface, using JavaScript fallback',
    );
  }
} catch (error) {
  console.log('Native validator not found, using JavaScript fallback');
}

export function validateTriangle(a: number, b: number, c: number): boolean {
  if (a <= 0 || b <= 0 || c <= 0) {
    return false;
  }

  if (nativeValidator) {
    return nativeValidator.validateTriangle(a, b, c);
  }

  return a + b > c && a + c > b && b + c > a;
}

export function createTriangleSafe(
  a: number,
  b: number,
  c: number,
  name?: string,
): Triangle | null {
  if (!validateTriangle(a, b, c)) {
    console.error(`Invalid triangle sides: ${a}, ${b}, ${c}`);
    return null;
  }
  return new Triangle(a, b, c, name);
}
