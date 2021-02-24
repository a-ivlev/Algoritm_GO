// LRU (Least Recently Used) — это алгоритм, при котором вытесняются значения, которые дольше всего не запрашивались.
// для решения данной задачи я использовал карту. В качестве ключа я использовал "бит времени",
// число которое позволяет определить какая карта добавлена последней и к какой карте был последний доступ.
package main

import "fmt"

var lru = make(map[int]string)
var maxLen = 3
var max, min int


func Add(data string) {
		lru[max] = data
		max++
	if  len(lru) > maxLen{
		delete(lru, min)
		min = max
		for key, _ := range lru {
			if key < min {
				min = key
			}
		}
	}
}

func Get(index int)  string {
	lru[max] = lru[index]
	delete(lru, index)
	min = max
	max++
	for key, _ := range lru {
		if key < min {
			min = key
		}
	}
	return lru[max]
}

func main()  {
	Add("A")
	Add("B")
	Add("C")
	fmt.Println("Пусть кеш расчитан на 3 элемента, а мы добавляем 4.")
	Add("D")
	fmt.Println("Смотрим что получилось.")
	fmt.Println(lru)
	fmt.Println(min)
	fmt.Println(max)
	fmt.Println("Делаем запрос к элементу с ключём 2.")
	Get(2)
	fmt.Println("Смотрим что получилось.")
	fmt.Println(lru)
	fmt.Println(min)
	fmt.Println(max)
	fmt.Println("Добавляем новый элемент.")
	Add("F")
	fmt.Println("Делаем запрос к элементу с ключём 3.")
	Get(3)
	fmt.Println(lru)
	fmt.Println(min)
	fmt.Println(max)
}
