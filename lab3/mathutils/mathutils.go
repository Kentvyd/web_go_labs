package mathutils

// Factorial вычисляет факториал числа n (n!).
// Для n <= 1 возвращает 1 (0! = 1! = 1 по определению).
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}

	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
