
## [Exercise 3.1](https://github.com/Property-Finder-Patika/week-1-homework-1-emretask1n/tree/main/Chapter3/exercises/ex3.1)

If the function f returns a non-finite `float64` value, the SVG file will contain invalid `<polygon>` elements (although many SVG renderers handle this gracefully). Modify the program to skip invalid polygons.


## [Exercise 3.2](https://github.com/Property-Finder-Patika/week-1-homework-1-emretask1n/tree/main/Chapter3/exercises/ex3.2)

Experiment with visualizations of other functions from the math package. Can you produce an `egg box`, `moguls`, or a `saddle`?


## [Exercise 3.3](https://github.com/Property-Finder-Patika/week-1-homework-1-emretask1n/tree/main/Chapter3/exercises/ex3.3)

Color each polygon based on its height, so that the peaks are colored red (#ff0000) and the valleys blue (#0000ff)


## [Exercise 3.4](https://github.com/Property-Finder-Patika/week-1-homework-1-emretask1n/tree/main/Chapter3/exercises/ex3.4)

Following the approach of the `Lissajous` example in Section 1.7, construct a web server that computes surfaces and writes SVG data to the client. The server must set the `Content-Type` header like this:

```go
w.Header().Set("Content-Type", "image/svg+xml")
```


## [Exercise 3.5](https://github.com/Property-Finder-Patika/week-1-homework-1-emretask1n/tree/main/Chapter3/exercises/ex3.5)

Implement a full-color `Mandelbrot set` using the function `image.NewRGBA` and the type `color.RGBA` or `color.YCbCr`.


## [Exercise 3.6](https://github.com/Property-Finder-Patika/week-1-homework-1-emretask1n/tree/main/Chapter3/exercises/ex3.6)

Supersampling is a technique to reduce the effect of pixelation by computing the color value at several points within each pixel and taking the average. The simplestmethod is to divide each pixel into four "subpixels". Implement it.


## [Exercise 3.7](https://github.com/Property-Finder-Patika/week-1-homework-1-emretask1n/tree/main/Chapter3/exercises/ex3.7)

Another simple fractal uses `Newton's method` to find complex solutions to a function such as z^4-1=0. Shade each starting point by the number of iterations required to get close to one of the four roots. Color each point by the root it approaches.


## [Exercise 3.8](https://github.com/Property-Finder-Patika/week-1-homework-1-emretask1n/tree/main/Chapter3/exercises/ex3.8)

Rendering fractals at high zoom levels demands great arithmetic precision. Implement the same fractal using four different representations of numbers: `complex64`, `complex128`, `big.Float`, and `big.Rat.` (The latter two types are found in the `math/big` package. `Float` uses arbitrary but bounded-precision floating-point; `Rat` uses unbounded-precision rational numbers) How do theycompare in performance and memory usage? At what zoom levels do rendering artifacts become visible?


## [Exercise 3.9](https://github.com/Property-Finder-Patika/week-1-homework-1-emretask1n/tree/main/Chapter3/exercises/ex3.9)

Write a web server that renders fractals and writes tye image data to the client. Allow the client to specify the `x,y`, and `zoom` values as parameters to the HTTP request.


