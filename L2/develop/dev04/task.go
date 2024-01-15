package main

import (
	"fmt"
	"reflect"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

*/

func anagramms(s []string) map[string][]string {
	m := make(map[string][]string)
	mRunes := make([]map[rune]int, 0)

	for i := range s {
		//counts is a map which display word as a set of letters with num of them in the word.
		//each word convert to the map "count" with char as a key and amount of chars as a value
		counts := make(map[rune]int)
		//to lower case
		s[i] = strings.ToLower(s[i])
		rs := []rune(s[i])
		for _, char := range rs {
			counts[char]++
		}
		//mRunes is a representation of a word as a set of letters indicating the number of letters in the word
		mRunes = append(mRunes, counts)
	}
	//flag to display unique elements
	exists := make(map[string]bool)
	for i := 0; i < len(mRunes); i++ {
		row := make([]string, 0)
		for j := i + 1; j < len(mRunes); j++ {
			//!exists added to display only words which haven't been added earlier to row or as a key to m.
			//so it is to escape this result:  пятак:[пятка тяпка] пятка:[тяпка]
			if reflect.DeepEqual(mRunes[i], mRunes[j]) && !exists[s[j]] {
				row = append(row, s[j])
				m[s[i]] = row
				exists[s[j]] = true
			}
		}
		//sets of one word have not be reflected. To reflect it uncomment code below
		/*if len(row) == 0 && !exists[s[i]] {
			m[s[i]] = row
			exists[s[i]] = true
		}*/
	}
	return m
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "Листок", "слиток", "столик", "ав", "ыыва", "тяпка"}
	anagrams := anagramms(words)
	fmt.Println(anagrams)
}
