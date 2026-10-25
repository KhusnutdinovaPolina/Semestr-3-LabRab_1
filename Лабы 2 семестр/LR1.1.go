package main

import (
	"fmt"
	"math"
)

// Исходная функция f(x) = x^2 + e^x - 2
func f(x float64) float64 {
	return x*x + math.Exp(x) - 2
}

func main() {
	var a, b, eps float64

	fmt.Println("Метод половинного деления")
	fmt.Print("Введите левую границу a: ")
	fmt.Scan(&a)
	fmt.Print("Введите правую границу b: ")
	fmt.Scan(&b)
	fmt.Print("Введите точность eps: ")
	fmt.Scan(&eps)

	// Проверка знаков
	if f(a)*f(b) >= 0 {
		fmt.Println("Ошибка: f(a) и f(b) одного знака!")
		return
	}

	// Срезы (как векторы в C++) для хранения истории
	var a_vals []float64
	var b_vals []float64
	var len_vals []float64

	// Сохраняем начальные значения
	a_vals = append(a_vals, a)
	b_vals = append(b_vals, b)
	len_vals = append(len_vals, b-a)

	// Основной итерационный процесс
	for (b - a) > eps {
		c := (a + b) / 2.0

		if f(a)*f(c) < 0 {
			b = c
		} else {
			a = c
		}

		a_vals = append(a_vals, a)
		b_vals = append(b_vals, b)
		len_vals = append(len_vals, b-a)
	}

	// Вывод таблицы
	fmt.Println("\n------------------------------------------------")
	fmt.Println(" N |      a_n      |      b_n      |   b_n - a_n  ")
	fmt.Println("------------------------------------------------")

	for i := 0; i < len(a_vals); i++ {
		fmt.Printf("%2d | %12.6f | %12.6f | %12.6f\n", i, a_vals[i], b_vals[i], len_vals[i])
	}
	fmt.Println("------------------------------------------------")

	root := (a_vals[len(a_vals)-1] + b_vals[len(b_vals)-1]) / 2.0
	fmt.Printf("Корень уравнения: %.6f\n", root)
	fmt.Printf("Проверка f(x) = %.6f\n", f(root))
	fmt.Printf("Число итераций: %d\n", len(a_vals)-1)
}
