package unittest

import "testing"

type TestS struct {
	n1    int
	n2    int
	total int
}

func TestSum(t *testing.T) {
	table := []TestS{
		{1, 2, 3},
		{2, 2, 4},
		{25, 25, 50},
	}
	for _, item := range table {
		total := Sum(item.n1, item.n2)
		if total != item.total {
			t.Errorf("Incorrect sum, it have %d and it should have %d", total, item.total)
		}
	}
}

func TestGetMax(t *testing.T) {
	table := []TestS{
		{4, 2, 4},
		{5, 3, 5},
		{50, 785, 785},
	}
	for _, item := range table {
		max := GetMax(item.n1, item.n2)
		if max != item.total {
			t.Errorf("Incorrect get max, it have %d and it should have %d", max, item.total)
		}
	}
}

func TestFibonacci(t *testing.T) {
	table := []struct {
		n    int
		fibo int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, v := range table {
		fibo := Fibonacci(v.n)
		if fibo != v.fibo {
			t.Errorf("Incorrect Fibonacci, it have %d and it should have %d", fibo, v.fibo)
		}
	}
}
