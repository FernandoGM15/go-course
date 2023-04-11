package unittest

func Sum(a, b int) int {
	return a + b
}

func GetMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

//To get the coverage of the code testing: go test -cover
//To get the files that have a test a go coverage file: test -coverprofile=coverage.out
