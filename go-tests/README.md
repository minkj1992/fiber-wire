# go-testing
> https://tutorialedge.net/courses/go-testing-bible/

```bash
$ go test ./... -v
$ cd calculator && go test -run=TestCalculateIsArmstrong -v
$ cd calculator && go test -run=TestCalculateIsArmstrongWithTable -v

# cov
$ go test ./... --cover
$ go test ./... -coverprofile=coverage.out
$ go tool cover -html=coverage.out
```

## Index
- Fundamentals
  - [ ] Sub-tests in Go
  - [ ] Table Driven Testing in Go
  - [ ] Test Coverage
  - [ ] Testdata Directory for Tests (fixtures)
  - [ ] Benchmarking Your Code
  - [ ] Linting and Formatting
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
