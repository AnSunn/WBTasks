package dev08

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

// To do fork (https://github.com/kraken-hpc/go-fork/blob/main/README.md)
// to do ps
//to launch: go test -v. Or change the pack name to main and add here func main(), remove parameters in shell function to be shell()
/*
=== Взаимодействие с ОС ===
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:
- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*
Также требуется поддерживать функционал fork/exec-команд
Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).
*/
var commands []string

// argument in shell is required for task_test as we enter stdin there
func Shell(input io.Reader) {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			fmt.Println("you don't input anything")
			os.Exit(0)
		}
		fmt.Fprintln(os.Stderr, "error reading input", err)
	}

	commands = strings.Fields(str)
	switch commands[0] {
	case "cd":
		if len(commands) < 2 {
			fmt.Fprintln(os.Stderr, "Incorrect mask, please use this one: cd <path>")
		}
		cd(commands[1])
	case "pwd":
		dir := pwd()
		fmt.Println("Current working directory: ", dir)
	case "echo":
		fmt.Println(strings.Join(commands[1:], " "))
	case "kill":
		kill(commands[1])
	/*case "fork":
	fmt.Printf("main() pid: %d\n", os.Getpid())
	if err := Fork("child", 1); err != nil {
		log.Fatalf("failed to fork: %v", err)
	}*/
	default:
		cmd := exec.Command(commands[0], commands[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
}

// cd function changes directory
func cd(command string) {
	// Calling the function Chdir() to change the directory
	err := os.Chdir(command)
	if err != nil {
		_, errPr := fmt.Fprintf(os.Stderr, "%v\n", err)
		if errPr != nil {
			fmt.Printf("The error printing err value to os.Stderr %v\n", errPr)
		}
	}
	dir := pwd()
	fmt.Println("New working directory: ", dir)

}
func pwd() string {
	Dir, err := os.Getwd()
	if err != nil {
		_, errPr := fmt.Fprintf(os.Stderr, "%v\n", err)
		if errPr != nil {
			fmt.Printf("The error printing err value to os.Stderr %v\n", errPr)
		}
	}
	return Dir
}
func kill(command string) {
	cmd := exec.Command("sleep", command)
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// Kill it:
	if err := cmd.Process.Kill(); err != nil {
		log.Fatal("failed to kill process: ", err)
	}
}

//to test issue here remove argument in shell function, change package name and uncomment code below
/*func main(){
	Shell()
}*/
