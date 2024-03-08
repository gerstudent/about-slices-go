package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func changeSlc(funcSlc []int) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&funcSlc))
	fmt.Println("--------- FUNC START ---------")
	fmt.Println("(IN FUNC) header was copied BY VALUE:", header)
	funcSlc[0] = 1
	funcSlc = append(funcSlc, 2)
	fmt.Println("(IN FUNC) header after append:", header)
	fmt.Println("(IN FUNC) funcSlc:", funcSlc)
	fmt.Println("--------- FUNC END ---------")
}

func main() {
	slc := make([]int, 1, 2)
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slc))
	fmt.Println("slc:", slc, header)
	changeSlc(slc)
	fmt.Println("slc:", slc, header)

	/*
		slc и funcSlc имеют одно хранилище данных, поэтому       расширение slc до cap(slc) позволит получить элементы,       добавленные в функции.
	*/
	slc = slc[:cap(slc)]
	fmt.Println("slc after cap grow:", slc, header)
	fmt.Println()

	/*
		После превышения емкости при добавлении элементов
		funcSlc изменит хранилище данных (SliceHeader),
		поэтому slc останется прежним, увеличение cap не поможет.
	*/
	changeSlc(slc)
	fmt.Println("slc:", slc, header)
	fmt.Println("slc after cap grow:", slc, header)
	fmt.Println()

	/*
		Для копирования слайсов используется функция copy.
		После копирования два слайса будут иметь разные
		хранилища данных (структуры SliceHeader).
	*/
	fmt.Println("--------- Slice copy ---------")
	slcCopy := make([]int, len(slc))
	copy(slcCopy, slc)
	headerCopy := (*reflect.SliceHeader)(unsafe.Pointer(&slcCopy))
	fmt.Println("slc, header:", slc, header)
	fmt.Println("slcCopy, headerCopy:", slcCopy, headerCopy)

	slc[0] = 100
	fmt.Println("slc after change:", slc)
	fmt.Println("slcCopy after change:", slcCopy)
}
