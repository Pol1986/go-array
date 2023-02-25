package main

import "fmt"

func main() {
//Массив
	array1 := []int{1, 2, 6, 9}
	array2 := []int{3, 4, 5, 7, 8}
	arrSort := []int{9, 7, 5, 4, 6, 1}
//Печать массивов
	fmt.Println(array1)
	fmt.Println(array2)
//Функция
	arrayMerch(array1, array2)

	fmt.Println("Несортированный массив:", arrSort)
	bubbleSort(arrSort)

}
//Сортировка
func bubbleSort(arrSort []int) {
	//len:=len(arrSort)
	for i := 0; i < len(arrSort)-1; i++ {
		for j := 0; j < len(arrSort)-i-1; j++ {
			if arrSort[j] > arrSort[j+1] {
				arrSort[j], arrSort[j+1] = arrSort[j+1], arrSort[j]
			}
		}
	}
	fmt.Println("Отсортированный массив:", arrSort)
}

func arrayMerch(array1 []int, array2 []int) {
	var array3 [9]int

	countArray1 := len(array1)
	countArray2 := len(array2)
	xInt, yInt := 0, 0
	for i := 0; i < len(array3); i++ {
		if xInt >= countArray1 {
			array3[i] = array2[yInt]
			yInt++
			continue
		} else if yInt >= countArray2 {
			array3[i] = array1[xInt]
			xInt++
			continue
		}

		if array1[xInt] > array2[yInt] {
			array3[i] = array2[yInt]
			yInt++
		} else {
			array3[i] = array1[xInt]
			xInt++
		}

	}
	fmt.Println("Объединенный массив: ", array3)
}
