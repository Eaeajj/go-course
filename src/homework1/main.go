package main

import (
	"course/src/utils"
	"fmt"
	"math"
	"unsafe"
)

// Задача:
// Дан слайс
// Необходимо заменить четные числа на 1 и посчитать сумму чисел в новом слайсе
// numbers := []int{3, 5, 7, 2, 7, 8, 6, 4, 7, 0, 1, 7, 4, 8, 10, 3, 6, 8, 5, 4, 12, 3}
func task1() {
	utils.PrintTaskNumber(1)
	numbers := []int{3, 5, 7, 2, 7, 8, 6, 4, 7, 0, 1, 7, 4, 8, 10, 3, 6, 8, 5, 4, 12, 3}
	newNumbers := make([]int, len(numbers))
	for i, num := range numbers {
		if num%2 == 0 {
			newNumbers[i] = 1
		} else {
			newNumbers[i] = num
		}
	}

	sum := 0
	for _, num := range newNumbers {
		sum += num
	}

	fmt.Println("New slice:", newNumbers)
	fmt.Printf("The sum is: %d\n", sum)
}

// Задача:
// Создайте слайс целых чисел и заполните его числами от 1 до 10.
// Используя цикл, пройдите по слайсу и увеличьте каждое значение на 5, используя указатель.
// Выведите измененный слайс.
func task2() {
	utils.PrintTaskNumber(2)

	numbers := make([]int, 10)
	for i := 0; i < 10; i++ {
		numbers[i] = i + 1
	}

	// Шаг 2: Увеличение значений с помощью указателя
	ptr := unsafe.Pointer(&numbers[0])
	size := unsafe.Sizeof(int(0))

	for i := 0; i < len(numbers); i++ {
		addr := unsafe.Pointer(uintptr(ptr) + uintptr(i)*size)

		val := *(*int)(addr)
		val += 5

		*(*int)(addr) = val
	}

	fmt.Println("New slice:")
	fmt.Println(numbers)
}

// Задача:
// Дан слайс чисел, необходимо найти минимальное и максимальное значение, которое делится на 2 без остатка.
// numbers := []int{8, 44, 3, 5, 11, 8, 2, 10, 6, 77, 15, 12}
func findMinAndMaxEven(numbers []int) (int, int) {
	minEven := math.MaxInt32
	maxEven := math.MinInt32

	for _, num := range numbers {
		if num%2 == 0 {
			minEven = min(minEven, num)
			maxEven = max(maxEven, num)
		}
	}

	return minEven, maxEven
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func task3() {
	utils.PrintTaskNumber(3)
	numbers := []int{8, 44, 3, 5, 11, 8, 2, 10, 6, 77, 15, 12}

	minEven, maxEven := findMinAndMaxEven(numbers)
	fmt.Printf("Минимальное четное число: %d\n", minEven)
	fmt.Printf("Максимальное четное число: %d\n", maxEven)
}

func main() {
	task1()
	task2()
	task3()

}
