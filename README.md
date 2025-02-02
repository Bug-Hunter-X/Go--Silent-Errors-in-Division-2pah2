# Go: Silent Errors in Division

This example demonstrates a common error in Go programs: silently ignoring errors that might occur during division.

The `Calculate` function correctly handles division by zero by returning an error. However, the `main` function only prints the error if it exists, omitting the error's details. This can make debugging challenging.

The solution demonstrates the proper way to handle the error, including printing its details.

## How to Run
1. Save the code in `bug.go` and `bugSolution.go`
2. Run `go run bug.go` to see the original bug.
3. Run `go run bugSolution.go` to see the solution.
