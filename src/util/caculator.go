package util

func Add(x int, y int) int {
	return x + y
}

func Sum(numbers ...int) int {
	var total int
	for _, number := range numbers {
		total += number
	}
	return total
}

func Subtract(x int, y int) int {
	return x - y
}

func Multiple(x int, y int) int {
	return x * y
}

func Divide(x int, y int) int {
	return x / y
}

