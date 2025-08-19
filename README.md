![Final Image](images/out8.jpg)


# go-tracer
A ray tracer written in Go - adapted from [Ray Tracing In One Weekend](https://raytracing.github.io/books/RayTracingInOneWeekend.html)

## Goals
- Get better at Go programming
- Learn about the basics of computer graphics (e.g., Tracing, Refraction, Linear Algebra applications)
- Learn about concurrency and using multiple CPU cores for rendering

## Features
- Multi-threaded rendering
- Materials (Glass, Metal, etc) 
- Unit Tests

## Running the Ray Tracer

### Setup
1. `go install`

### Running Options
The ray tracer can be run in either single-threaded or multi-threaded mode:

```bash
# Run with multi-threading (default)
go run main.go

# Run with single threading
go run main.go -multi=false
```

### Performance
Performance measurements for the sample scene (on average):
- Single-threaded mode: ~38 seconds
- Multi-threaded mode: ~20 seconds

The multi-threaded implementation provides approximately a 1.9x speed-up in rendering time compared to the single-threaded version! 
(Results may vary based on your CPU, system resources, etc)

