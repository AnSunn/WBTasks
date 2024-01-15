package dev08

import (
	"os"
	"strings"
	"testing"
)

func TestShell(t *testing.T) {
	test := struct {
		name   string
		input  string
		output string
	}{
		name:   "fork",
		input:  "for os.Getpid()",
		output: "Hello, World!\n",
	}

	t.Run(test.name, func(t *testing.T) {
		// Save old stdin
		oldStdin := os.Stdin

		//Create new Pipe for entering
		r, w, err := os.Pipe()
		if err != nil {
			t.Fatal(err)
		}

		// change stdin to pipe
		os.Stdin = r

		// write to pipe
		w.Write([]byte(test.input))
		w.Close()

		//call Shell
		Shell(strings.NewReader(test.input))

		// Restore stdin
		os.Stdin = oldStdin
	})
}
