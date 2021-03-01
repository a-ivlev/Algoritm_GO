// LRU (Least Recently Used) — это алгоритм, при котором вытесняются значения, которые дольше всего не запрашивались.
// для решения данной задачи я использовал карту. В качестве ключа я использовал то что собираемся кешировать,
// а в качестве значения я использовал "бит времени", число которое позволяет определить какая карта была давно добавлена
// или какая карта давно не запрашивалась.
package main

import "fmt"

var lru = make(map[string]int)
var maxLen = 3
var count int = 1
var del string

func Add(data string) {
		lru[data] = count
		count++
	if  len(lru) > maxLen{
		min := count
		for key, value := range lru {
			if value < min {
				min = value
				del = key
			}
		}
		delete(lru, del)
	}
}

func Get(data string)  (string, int) {
	if lru[data] == 0 {
		return "Элемент не найден", 0
	}
	lru[data] = count
	count++
	return data, lru[data]
}

func main()  {
	Add("A")
	Add("B")
	Add("C")
	fmt.Println(lru)
	fmt.Println("Пусть кеш расчитан на 3 элемента, а мы добавляем 4 элемент \"D\".")
	Add("D")
	fmt.Println("Выводим на экран карту.\n", lru)
	fmt.Println(`Бит времени по ключу "A"`, lru["A"])
	fmt.Println(`Бит времени по ключу "B"`, lru["B"])
	fmt.Println(`Делаем запрос к элементу с ключём "B".`)
	fmt.Println(Get("B"))
	fmt.Println("Выводим на экран карту.\n", lru)
	fmt.Println("Добавляем новый элемент. \"F\"")
	Add("F")
	fmt.Println("Выводим на экран карту.\n", lru)
	fmt.Println("Делаем запрос к элементу с ключём \"C\".")
	fmt.Println(Get("C"))
	fmt.Println("Делаем запрос к элементу с ключём \"W\".")
	fmt.Println(Get("W"))
	fmt.Println("Выводим на экран карту.\n", lru)
	fmt.Println("Добавляем новый элемент. \"W\"")
	Add("W")
	fmt.Println("Выводим на экран карту.\n", lru)
}
