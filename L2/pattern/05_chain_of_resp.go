package pattern

import "fmt"

/*
Implementation of "chain of responsibility" pattern - behavioral pattern (https://en.wikipedia.org/wiki/Factory_method_pattern)

Applicability of chain of resp pattern:
1. Importance of handlers sequence - one performs after another
2. Dynamic determination of the handler
When the set of handlers and their order may change dynamically at runtime
3. Multiple objects may handle a request
When you want to allow multiple objects to process a request, and the handler is not known in advance.

Pros:
1. Reduce dependency between the client and handler
2. Single Responsibility Principle (Принцип единственной обязанности)
Handler performs individual task and doesn't know about other handler details (Каждый обработчик концентрируется на своей задаче
и не знает о деталях обработки, выполняемых другими частями системы.)
3. Open/Closed Principle (Принцип открытости/закрытости)
This pattern is opened for expanding and closed for changes, so we can add new handlers without changing client's code.

Cons:
1. Guarantee of Handling
There's no guarantee that a request will be handled. If the end of the chain is reached and the request is still unhandled,
it may not be clear how to proceed.

Real examples of Chain of Responsibility pattern:
1.Event Handling in GUIs
The Chain of Responsibility is often used in graphical user interface frameworks for handling events.
Components in the UI hierarchy can be arranged in a chain, where each component has the option to handle an event or pass it to the next component in the chain.
2. Logging Systems
In logging frameworks, different loggers can be organized in a chain.
Each logger decides whether it can handle a log message or should pass it to the next logger in the chain.
*/

type ApplicantForVisa struct {
	Name                string
	applyOnlineDone     bool
	sendDocToOfficeDone bool
	paymentDone         bool
}

// handler interface
type Handler interface {
	Execute(*ApplicantForVisa)
}

// onlineApplication - concrete handler
type OnlineApplication struct {
	Next Handler
}

func (onlineApp *OnlineApplication) Execute(applicant *ApplicantForVisa) {
	if applicant.applyOnlineDone {
		fmt.Println("Online application is already done")
		onlineApp.Next.Execute(applicant)
		return
	}

	fmt.Println("Online application is successful")
	applicant.applyOnlineDone = true
	if onlineApp.Next != nil {
		onlineApp.Next.Execute(applicant)
	}
	return
}

//officer document verification - concrete handler

type OfficerDocVerification struct {
	Next Handler
}

func (doc *OfficerDocVerification) Execute(applicant *ApplicantForVisa) {
	if applicant.sendDocToOfficeDone {
		fmt.Println("The docs are already checked by the officer")
		doc.Next.Execute(applicant)
		return
	}

	fmt.Println("Docs are successfully verified by the officer")
	applicant.sendDocToOfficeDone = true
	if doc.Next != nil {
		doc.Next.Execute(applicant)
	}
	return
}

// payment - concrete handler
type Payment struct {
	next Handler
}

func (payment *Payment) Execute(applicant *ApplicantForVisa) {
	if applicant.paymentDone {
		fmt.Println("The payment is already proceeded")
		payment.next.Execute(applicant)
		return
	}

	fmt.Println("The payment is successful")
	applicant.paymentDone = true
	if payment.next != nil {
		payment.next.Execute(applicant)
	}
	return
}
