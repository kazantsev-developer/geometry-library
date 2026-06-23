# Geometry Library

A comprehensive 2D geometry library built with TypeScript, featuring a high-performance native C++ addon.

## Features

* **Event-Driven Architecture:** Core system extends native `EventTarget` to provide automated notifications upon shape state changes.
* **Strict Typing:** Out-of-the-box support for TypeScript configured in strict mode.
* **Extensibility:** Built in accordance with the Open/Closed Principle; new geometric shapes (e.g., trapezoids, polygons) can be introduced without modifying the core system.
* **Standard Geometric Interfaces:** Native methods for calculating area, perimeter, and robust parameter validation.
* **Data Export:** Built-in asynchronous generation for textual reports.
* **Native C++ Addon:** Offloads triangle parameter validation to C++ for maximized computational performance.

## Quick Start

A complete usage demonstration is available in `examples/example.ts`.

### Available Commands

```bash
# Install project dependencies
npm install

# Compile the native C++ addon (required before first run)
npm run build:addon

# Compile TypeScript source code
npm run build

# Execute the main TypeScript demonstration
npm run example

# Execute the native C++ addon demonstration
npm run example:addon
```

## C++ Addon

The library includes a native addon engineered specifically for triangle validation workflows. When initializing a triangle, side verification is processed via the C++ layer, significantly increasing execution speed during high-volume operations.

**Note:** If the C++ addon is missing or uncompiled, the library automatically activates a JavaScript implementation fallback to ensure uninterrupted execution.
