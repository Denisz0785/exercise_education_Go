package main

//функция меняет последовательность элементов в срезе на обратную

func SortArray(src []int) {
	leng := len(src) - 1
	for i := 0; i <= leng; i++ {
		j := leng - i
		if i < j {
			temp := src[i]
			src[i] = src[j]
			src[j] = temp

		} else {
			break
		}
	}
}

// функция сортирует срез методом вставки
func InsertSorting(src []int) {
	leng := len(src)
	for i := 1; i < leng; i++ {
		tempIndex := i
		tempValue := src[i]
		for j := tempIndex - 1; j >= 0; j-- {
			if tempValue < src[j] {
				src[j+1] = src[j]
				tempIndex--
			}

		}
		if tempIndex != i {
			src[tempIndex] = tempValue
		}
	}
}
