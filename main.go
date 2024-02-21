package main

import "fmt"

func search(num, i int) {
	if i != 0 && num%i == 0 {
		fmt.Println("Число", num, "делится на", i, "без остатка")
	} else if i != 0 {
		temp := num % i
		fmt.Printf("Остаток от деления = %v \n", temp)
	} else {
		fmt.Println("Данную операцию не возможно осуществить")
	}
}

func search2(num, i int) {
	if i != 0 {
		switch temp := num % i; temp {
		case 0:
			fmt.Printf("Число %v делится на %v без остатка\n", num, i)
		case 1:
			fmt.Printf("Остаток от деления %d на %d = %v\n", num, i, temp)
		case 2:
			fmt.Printf("Остаток от деления %d на %d = %v\n", num, i, temp)
		case 3:
			fmt.Printf("Остаток от деления %d на %d = %v\n", num, i, temp)
		case 4:
			fmt.Printf("Остаток от деления %d на %d = %v\n", num, i, temp)
		case 5:
			fmt.Printf("Остаток от деления %d на %d = %v\n", num, i, temp)
		case 6:
			fmt.Printf("Остаток от деления %d на %d = %v\n", num, i, temp)
		case 7:
			fmt.Printf("Остаток от деления %d на %d = %v\n", num, i, temp)
		case 8:
			fmt.Printf("Остаток от деления %d на %d = %v\n", num, i, temp)
		case 9:
			fmt.Printf("Остаток от деления %d на %d = %v\n", num, i, temp)
		default:
			fmt.Println("Операция не может быть осуществлена")

		}
	} else {
		fmt.Println("Деление на 0 невозможно")
	}
}

func search3(num, i int) {
	if i != 0 {
		temp := num % i
		switch {
		case temp == 0:
			fmt.Println("Остаток от деления = 0")
		case temp < 5:
			fmt.Println("Остаток меньше 5")
		case 5 <= temp && temp < 7:
			fmt.Println("Остаток больше 4х и меньше 7")
		case 7 <= temp:
			fmt.Println("Остаток больше 6")
		default:
			fmt.Println("Операция не может быть осуществлена")

		}

	} else {
		fmt.Println("Операция не может быть осуществлена")
	}

}

func main() {

	search(8, 5)

	src := make([]int, 6, 10)

	for i := 0; i < 6; i++ {
		src[i] = i * 3
	}

	for i := 0; i < len(src)-1; i++ {
		search2(src[i], 2)
	}

	for _, v := range src {
		if v%2 == 0 || v%4 == 0 {
			fmt.Printf("Число %v делится без остатка\n", v)
		} else {
			fmt.Printf("Остаток не равен нулю\n")

		}

	}

	fmt.Println("search3:")
	search3(95, 3)
}
