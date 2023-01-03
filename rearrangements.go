/*Скрипт, который поместит + (2+3), - (3-2), или ничего ( ) в промежутках между цифрами от 9 до 0 (в таком порядке) так, чтобы в результате получилось 200. Например: 98+76-5+43-2-10=200*/

package main

import (
	"fmt"
	"math"
)

func main() {
	intSlice := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0} //данный слайс является стартовым
	relevantElems := [][]int{}                      //в данный слайс слайсов будут занесены все возможные элементы, которые релевантны для перебора
	for i := 0; i < len(intSlice); i++ {
		relevantElems = append(relevantElems, elemsFromOneSlice(intSlice[i:]))
	}

	counters := []int{} //для перебора всех релевантных элементов мы создадим слайс со счётчиком
	for i := 0; i < len(relevantElems); i++ {
		counters = append(counters, 0)
	}

	finalSlice := [][]int{}
	for { //в бесконечном цикле мы будем перебирать все комбинации до предельного значения счетчика
		sum := 0                              //переменная для проверки равенства суммы элементов цифре '200'
		ans := []int{}                        //в данную переменную падают потенциальные ответы
		for i := 0; i < len(relevantElems); { //в цикле мы проходимся по релевантным значениям. если скрипт видит несочетаемые элементы, он из не записывает в "sum"
			if math.Sqrt(float64(relevantElems[i][counters[i]]*relevantElems[i][counters[i]]))/100 >= 1 {
				sum += relevantElems[i][counters[i]]
				ans = append(ans, relevantElems[i][counters[i]])
				i += 3
			} else if math.Sqrt(float64(relevantElems[i][counters[i]]*relevantElems[i][counters[i]]))/10 >= 1 {
				sum += relevantElems[i][counters[i]]
				ans = append(ans, relevantElems[i][counters[i]])
				i += 2
			} else {
				sum += relevantElems[i][counters[i]]
				ans = append(ans, relevantElems[i][counters[i]])
				i += 1
			}
		}
		if sum == 200 { //если сумма равна '200', проверяем на дубликат
			finalSlice = checkDoubles(finalSlice, ans)
		}
		counters[0] += 1                          //после прохождения цикла сдвигаем элемент счётчика
		for s := 0; s < len(relevantElems); s++ { //проверяем на достижение предельных значений счётчика
			if counters[s] == len(relevantElems[s]) {
				if counters[s] == counters[len(counters)-1] { //по достижении предельного значения счетчика завершаем скрипт
					fmt.Println("Комбинации чисел, сумма которых даёт '200'", "\n")
					for i := 0; i < len(finalSlice); i++ {
						fmt.Println(finalSlice[i])
					}
					fmt.Println("\n", "END")
					return
				} else { //если элемент счётчика достиг промежуточного предельного значения, сдвигаем элемент счётчика
					counters[s] = 0
					counters[s+1] += 1
				}
			}
		}

	}
}

func checkDoubles(finalSlice [][]int, someSlice []int) [][]int {
	check := true                          //если слайса длинны "someSlice" нет в "finalSlice", будет "true"
	for i := 0; i < len(finalSlice); i++ { //ищем слайсы аналогичной длинны
		if len(someSlice) == len(finalSlice[i]) {
			check = false
			for j := 0; j < len(finalSlice[i]); j++ {
				if someSlice[j] != finalSlice[i][j] {
					check = true
				}
			}
			if check == false { //если есть хоть один инентичный слайс, будет "false"
				return finalSlice
			}
		}
	}
	finalSlice = append(finalSlice, someSlice) //в ином случае мы объединяем слайс в слайс с ответами
	return finalSlice
}

func elemsFromOneSlice(intSlice []int) []int { //функция генерирует все возможные цифры из принимаемого слайса, которые меньше или равны числу "1000" по модулю
	allVars := []int{}    //данный слайс будет результатом функции
	if intSlice[0] == 0 { //костыль на случай, если передан слайс из одного нуля
		allVars = []int{0}
		return allVars
	}
	allVars = append(allVars, intSlice[0])
	for i := 0; i < len(intSlice)-1; i++ {
		allVars = append(allVars, allVars[len(allVars)-1]*(-1))             //добавили отрицательное значение
		allVars = append(allVars, allVars[len(allVars)-2]*10+intSlice[i+1]) //добавили "склеенное" значение
		if i == 1 {                                                         // прерываем на моменте, когда более релевантных чисел мы не получим
			break
		}
	}
	element := allVars[len(allVars)-1]
	allVars = append(allVars, element*(-1))
	return allVars
}
