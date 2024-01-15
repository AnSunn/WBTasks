package pattern

import "fmt"

/*
Implementation of "command" pattern - behavioral pattern (https://en.wikipedia.org/wiki/Command_pattern)

Applicability of visitor pattern:
1. Parameterize objects
You want to parameterize objects by operations.
2. Queue requests
You want to queue operations, schedule their execution, or execute them remotely.
3. Support undo operations
You need to support undo/redo functionality in your system.
4. Support transactions
You want to support transactions where a sequence of operations is treated as a single operation.
5. Decouple sender and receiver
You want to decouple the sender (invoker) of a request from the receiver (object that performs the operation).

Pros:
1. Decoupling
It decouples the sender and receiver of a request, making the system more flexible and extensible.
2. Parameterization
It allows you to parameterize objects with operations, making it easy to implement various commands without modifying the client code.
3. Undo/Redo
It provides support for undo/redo functionality by keeping track of command history.
4. Queueing and Scheduling
It enables the queuing of requests and scheduling their execution.
5. Transaction Support
It supports transactions by grouping a set of operations into a single command.

Cons:
1. Increased Number of Classes
The pattern may lead to an increased number of classes in the system, which might be seen as an overhead
in small-scale applications.

Real examples of visitor pattern
1. Game Engines
Game engines often use the command pattern for handling player input. Each player action,
like moving characters or firing weapons, can be represented as a command.
2. Text Editors
Text editors use the command pattern to implement undo/redo functionality.
Each edit operation, such as insert or delete, is encapsulated in a command object.
3. GUI Applications
GUI frameworks often use the command pattern for handling user actions.
Each menu item, button click, or key press can be represented as a command object, allowing for undo/redo functionality.
*/

type Command interface {
	Execute()
}

// receiver interface
type Receiver interface {
	action()
}

// concrete command
type UpdateCommand struct {
	Receiver Receiver
}

func (upd *UpdateCommand) Execute() {
	upd.Receiver.action()
}

// receiver jpg
type JpgFile struct {
	Name string
}

func (pic *JpgFile) action() {
	fmt.Println("Draw image in", pic.Name)
}

// receiver word
type DocFile struct {
	Name string
}

func (doc *DocFile) action() {
	fmt.Println("Add word in", doc.Name)
}

type Invoker struct {
	history []Command
}

func (i *Invoker) ExecuteCommand(command Command) {
	command.Execute()
	i.history = append(i.history, command)
}

func (i *Invoker) UndoCommand() {
	if len(i.history) != 0 {
		i.history = i.history[:len(i.history)-1]
	}
}
