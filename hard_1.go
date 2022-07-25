/*Формат ввода
Первая строка содержит число n(1≤n≤100000) — количество резервуаров.
Во второй строке подаётся через запятую n целых чисел.

Формат вывода
Если объемы в резервуарах можно уравнять, выведите минимальное количество операций, необходимых для этого.
Если это невозможно, выведите «-1».*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var k int
	fmt.Scan(&k)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	a := scanner.Text()

	newA := []rune(strings.ReplaceAll(a, " ", ""))
	newAlean := len(newA)
	massive := []int{}

	for i := 0; i < newAlean; i++ {
		cercleInt, _ := strconv.Atoi(string(newA[i]))
		massive = append(massive, cercleInt)
	}

	maxVal := massive[newAlean-1]

	for i := 0; i < newAlean; i++ {
		if massive[i] > maxVal {
			fmt.Println(-1)
			return
		}
	}

	min := []int{}

	for d := 1; d <= k; d++ {
		min = append(min, 0)
		for j := 0; j < k; j++ {
			cercleNum := massive[j]
			for i := massive[j]; i < maxVal; i += d {
				min[d-1] = min[d-1] + 1
				cercleNum = cercleNum + d
			}
			if cercleNum != maxVal {
				min[d-1] = -1
			}
		}
	}

	intSlice := []int{}
	for i := 0; i < len(min); i++ {
		if min[i] != -1 {
			intSlice = append(intSlice, min[i])
		}
	}

	finMin := intSlice[0]
	if len(intSlice) != 1 {
		for i := 1; i < len(intSlice); i++ {
			if intSlice[i] < finMin {
				finMin = intSlice[i]
			}
		}
	}

	fmt.Println(finMin)
}
