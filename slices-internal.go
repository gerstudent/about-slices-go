package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func slicesCompare(slc1 []int, slc2 []int, name1, name2 string) {
	fmt.Printf("-------------- %s and %s --------------\n", name1, name2)
	fmt.Printf("%s: ", name1)
	fmt.Println(slc1, (*reflect.SliceHeader)(unsafe.Pointer(&slc1)))

	fmt.Printf("%s: ", name2)
	fmt.Println(slc2, (*reflect.SliceHeader)(unsafe.Pointer(&slc2)))
	fmt.Println("-----------------------------------------------")
	fmt.Println()
}

func main() {
	slc := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		slc = append(slc, i)
	}

	/*
		Первое число cuttedSlc на 16 больше, чем аналогичное у slc,
		так как структура SliceHeader cuttedSlc скопировалась.
		Она содержит два поля типа int, которые на 64-битной машине занимают
		8 + 8 = 16 байт. Отсюда и увеличение на 16.
	*/
	cuttedSlc := slc[2:4]
	slicesCompare(slc, cuttedSlc, "slc", "cuttedSlc")
	// _ = cuttedSlc[5] // index out of range [2] with length 2

	// Расширение слайса cuttedSlc
	cuttedSlc = cuttedSlc[:cap(cuttedSlc)]
	slicesCompare(slc, cuttedSlc, "slc", "cuttedSlc")
	// _ = cuttedSlc[5] // ошибок нет

	/*
		Изменение нулевого элемента слайса slc не затронет
		слайс cuttedSlc, а вот изменение нулевого элемента cuttedSlc
		(второго элемента slc) отразится на slc.
	*/
	slc[0] = 100
	cuttedSlc[0] = 55
	slicesCompare(slc, cuttedSlc, "slc", "cuttedSlc")

	/*
		После append емкость cuttedSlc увеличилась вдвое,
		Теперь slc и cuttedSlc имеют разные адреса в памяти.
	*/
	cuttedSlc = append(cuttedSlc, 66)
	slicesCompare(slc, cuttedSlc, "slc", "cuttedSlc")
}
