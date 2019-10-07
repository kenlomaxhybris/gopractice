# ggexamplep

Concepts

- race conditions
- mutrex
- atomic counters

Introduction material:
- go test -v -cover -race
- tour.golang.org -> Concurrency  > sync.Mutex
- https://gobyexample.com/stateful-goroutines  .. "This channel-based approach aligns with Go’s ideas of sharing memory by communicating and having each piece of data owned by exactly 1 goroutine."

Task:
- Create a race condition using maps, go routines, channels
- Reveal with go test -race
- Fix with Mutex
- Consider pros/cons
