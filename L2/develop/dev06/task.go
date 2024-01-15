package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//to launch write  go run github.com/AnSunn/WBTasks/L2/develop/dev06 -f="1,2" -d=";" -s=true. Or other flags
/*
=== Утилита cut ===

# Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
var (
	delimiter     string
	fields        string
	onlySeparated bool
)

func initialize() {
	//if flags are not assigned in go run, then use default values, mentioned below
	flag.StringVar(&delimiter, "d", " ", "Delimeter")
	flag.StringVar(&fields, "f", "1", "Fields to select (separated by delimiter)")
	flag.BoolVar(&onlySeparated, "s", false, "Only show lines with delimiters")
	flag.Parse()
}
func cut() {
	initialize()
	delimRows := make([][]string, 0)
	delimRows = delimRow()
	fmt.Println("The result, seperated by delimiter: ", delimRows)

	//to display only mentioned in f-flag columns consider that they have to be separated by delimiter assigned in d-flag,
	//e.g. input: {[dsd;sd;2d;ds2},[sd,f,g,h]} with d-flag="," and f-flag="2,4", the result is {[sd;ds2},[]} as only first row separated by ","
	fieldRows := make([][]string, 0)
	//if exist flag with column number
	if fields != "" {
		fieldsFlagInt, err := fieldFlagToInt(fields)
		if err != nil {
			fmt.Println(err)
			return
		}

		for i, _ := range delimRows {
			out := make([]string, 0)
			//range by f-flag values to extract required columns from delimRows
			for _, fieldsVal := range fieldsFlagInt {
				//if f flag refers to column which doesn't exist in delimRows, then [] will be displayed, e.g input: weew; and flag = 2,4;
				//there are no 2 and 4 columns in weew -> []
				if len(delimRows[i]) >= fieldsVal {
					out = append(out, delimRows[i][fieldsVal-1])
				}

			}
			fieldRows = append(fieldRows, out)
		}
	}
	fmt.Println("The result of selected in f-flag columns: ")
	//to save the order of values use for. In map order is chaotic
	for _, v := range fieldRows {
		fmt.Println(v)
	}
}

// fieldFlagToInt convert string values of flag "f" to int. Return slice of these int values
func fieldFlagToInt(field string) ([]int, error) {
	fieldsFlagInt := make([]int, 0)
	//flag with column num is string with comma separator, e.g. "2,5". We have to extract each value as int
	fieldList := strings.Split(fields, ",")
	for _, val := range fieldList {
		fieldsVal, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("Invalid field flag, unable to convert to int", err)
			return nil, err
		}
		fieldsFlagInt = append(fieldsFlagInt, fieldsVal)
	}
	return fieldsFlagInt, nil
}

// delimRow delimit row by separator
func delimRow() [][]string {
	scanner := bufio.NewScanner(os.Stdin)
	delimRows := make([][]string, 0)
	for scanner.Scan() {
		row := scanner.Text()
		if row == "" {
			break
		}
		//if s-flag = true, then the code below lets escape the rows which don't have delimiter
		if onlySeparated && !strings.Contains(row, delimiter) {
			continue
		}
		rowSplit := strings.Split(row, delimiter)
		delimRows = append(delimRows, rowSplit)
	}
	return delimRows
}

func main() {
	cut()
}
