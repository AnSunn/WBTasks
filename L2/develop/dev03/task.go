package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
===Утилита sort===
Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры): на входе подается файл из несортированными строками, на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:

-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки
*/
var (
	columnSort                           int
	numericSort, reverseSort, uniqueSort bool
)

func main() {
	// flag definition
	flag.IntVar(&columnSort, "k", 0, "Column index for sorting (default is entire line)")
	flag.BoolVar(&numericSort, "n", false, "Sort numerically")
	flag.BoolVar(&reverseSort, "r", false, "Sort in reverse order")
	flag.BoolVar(&uniqueSort, "u", false, "Output only unique lines")

	flag.Parse()

	// open file
	//file, err := os.Open(flag.Arg(0))
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Read string from the file
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// Функция сравнения для сортировки
	comparator := func(i, j int) bool {
		fl, err := compare(lines[i], lines[j], 2, false, false, false)
		if err != nil {
			fmt.Println(err)
		}
		return fl
	}

	// Сортировка строк
	sort.SliceStable(lines, comparator)
	// Вывод строк
	printLines(lines, true)
}
func compare(line1, line2 string, columnSort int, numericSort, reverseSort, uniqueSort bool) (bool, error) {
	//return slice of words from line1 (the slice is result of line1 separation by " ")
	fields1 := strings.Fields(line1)
	fields2 := strings.Fields(line2)

	var col1, col2 string
	if columnSort > 0 && columnSort <= len(fields1) {
		col1 = fields1[columnSort-1]
	} else {
		return false, errors.New("invalid argument for -k")
	}
	if columnSort > 0 && columnSort <= len(fields2) {
		col2 = fields2[columnSort-1]
	} else {
		return false, errors.New("invalid argument for -k")
	}

	if numericSort {
		val1, err1 := strconv.Atoi(col1)
		val2, err2 := strconv.Atoi(col2)

		if err1 == nil && err2 == nil {
			if reverseSort {
				return val1 > val2, nil
			}
			return val1 < val2, nil
		} else if err1 != nil {
			return false, err1
		} else if err2 != nil {
			return false, err2
		}
	}
	if reverseSort {
		return line1 > line2, nil
	}
	return line1 < line2, nil
}

func printLines(lines []string, unique bool) {
	if unique {
		// Логика для удаления повторяющихся строк
		uniqueLines := make(map[string]bool)
		for _, line := range lines {
			uniqueLines[line] = true
		}

		// Вывод уникальных строк
		for line := range uniqueLines {
			fmt.Println(line)
		}
	} else {
		// Вывод всех строк
		for _, line := range lines {
			fmt.Println(line)
		}
	}
}
