
#  ggexamplea

  

Concepts:
 - Go Test Framework
	 - Tests and t.Error 
	 - ExampleXXX : go test ./... -v -cover -race
	 - BenchmarkXXX: go test -bench=. 
	 - Test coverage, 
	 - Test racing conditions
	 - Test Helper
- Go first steps
	- packages
	- imports
	- function
	- main func
	- exported names (package public)
- Go CLI:
	- go run xxx.go
	- go build xxx.go
	- go test -v -cover -race
	- go test -bench=.
	- godoc -http ":8081"
	- gobenchui -last 10 .

Introduction material:
- tour.golang.org -> Basics > packages
- tour.golang.org -> Basics > imports
- tour.golang.org -> Basics > functions
- tour.golang.org -> Basics > exported names
- https://gobyexample.com/functions

Tasks:
- Wrap some (trivial) business function with a functional test, exampleTest, benchmarkTest and documentation.
- Confirm the documentation is showing your exampleTest
- Try out all the CLI options listed.

Further reading:
- https://golang.org/pkg/testing/
