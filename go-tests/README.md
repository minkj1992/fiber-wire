# go-testing
> https://tutorialedge.net/courses/go-testing-bible/

```bash
$ go test ./... -v
$ cd calculator && go test -run=TestCalculateIsArmstrong -v
$ cd calculator && go test -run=TestCalculateIsArmstrongWithTable -v
```

- coverage
```console
$ go test ./... --cover
$ go test ./... -coverprofile=coverage.out
$ go tool cover -html=coverage.out
```

- Benchmarking
```console
$ go test ./... -run=Benchmark -bench=.
goos: darwin
goarch: arm64
pkg: github.com/minkj1992/go-playground/go-tests/calculator
BenchmarkCalculateIsArmstrong370-10     50413156                23.64 ns/op
BenchmarkCalculateIsArmstrong371-10     51319698                21.84 ns/op
BenchmarkCalculateIsArmstrong0-10       100000000               11.02 ns/op
PASS
ok      github.com/minkj1992/go-playground/go-tests/calculator  4.571s
PASS
ok      github.com/minkj1992/go-playground/go-tests/ytml        0.277s
```

## Index
- Fundamentals
  - [x] Sub-tests in Go
  - [x] Table Driven Testing in Go
  - [x] Test Coverage
  - [x] Testdata Directory for Tests (fixtures)
  - [x] Benchmarking Your Code
  - [x] Linting and Formatting
- Intermediate
  - [ ] TestMain
  - [ ] Testify
  - [ ] Test HTTP endpoints with HTTPTest
  - [ ] Mocking with Testify
  - [ ] Mocking with vektra/mockery
- Advanced
  - [ ] Integration Testing
  - [ ] Differentiating Tests with Build Tags
  - [ ] Testing for Race Conditions with --race
  - [ ] Integration tests for REST API
  - [ ] E2E Testing
