# Quad (Go)

This Go package prints ASCII-art rectangles with different border styles. It provides five functions: `QuadA`, `QuadB`, `QuadC`, `QuadD`, and `QuadE`. Each one generates a rectangle of size `x` by `y` using different corner and edge characters.

## Usage

```go
package main

import "quad"

func main() {
    quad.QuadA(5, 3)
    quad.QuadB(5, 3)
    quad.QuadC(5, 3)
    quad.QuadD(5, 3)
    quad.QuadE(5, 3)
}
