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
