package pattern

import (
	"fmt"
)

/*
Implementation of "factory method" pattern - creational pattern (https://en.wikipedia.org/wiki/Factory_method_pattern)

Applicability of factory method pattern:
1. when we don't know in advance types and objects dependency (Благодаря этому, код производства можно расширять, не трогая основной. Так, чтобы добавить поддержку нового продукта,
вам нужно создать новый подкласс и определить в нём фабричный метод, возвращая оттуда экземпляр нового продукта)
2. Classes need to be decoupled from the system-specific classes
When you want to provide a framework for multiple implementations,
allowing the client code to use the framework without being tightly coupled to specific implementations.

Pros:
1. Decoupling
It decouples the client code from the specific classes it instantiates, allowing for more flexibility and easier maintenance.
2. Extensibility
It provides an extensible framework where new classes (subclasses) can be added without modifying the existing code.
3. Open/Closed Principle (Принцип открытости/закрытости)
This pattern is opened for expanding and closed for changes, so we can add new handlers without changing client's code.

Cons:
1. Complexity
Introducing multiple factory methods or creating a hierarchy of factory classes can add complexity to the code.
2. Subclass Proliferation
The pattern can lead to a proliferation of subclasses, especially if each concrete product requires its own subclass.
3. Constructor "Божественный конструктор" (всегда инициализируем объекты через один конструктор, сильная привязка)

Real examples of factory method pattern:
1. GUI Frameworks
In graphical user interface (GUI) frameworks, the creation of UI components (buttons, windows, etc.)
is often handled using the Factory Method pattern. Each UI toolkit may have its own set of factories for creating components.
2. Logging Frameworks
Logging frameworks may use the Factory Method pattern to create different types of loggers based on configuration settings.
For example, a logging framework might have different factories for creating file loggers, console loggers, or network loggers.
*/

const (
	ElectricCar = "electric"
	PetrolCar   = "petrol"
)

// product interface
type Car interface {
	GetType() string
	PrintInfo()
}

// CreateCar realizes factory
func CreateCar(typeName string) Car {
	switch typeName {
	default:
		fmt.Println("Non-existent car")
		return nil
	case ElectricCar:
		return NewElectricCar()
	case PetrolCar:
		return NewPetrolCar()
	}
}

// Electric car - concrete product
type electric struct {
	carType string
	fuel    string
}

func NewElectricCar() Car {
	return &electric{
		carType: ElectricCar,
		fuel:    "electricity",
	}
}
func (e *electric) GetType() string {
	return e.carType
}

func (e *electric) PrintInfo() {
	fmt.Printf("fuel: [%s], carType: [%s]\n", e.fuel, e.carType)
}

type petrol struct {
	carType string
	fuel    string
}

func NewPetrolCar() Car {
	return &petrol{
		carType: PetrolCar,
		fuel:    "petrol",
	}
}
func (e *petrol) GetType() string {
	return e.carType
}

func (e *petrol) PrintInfo() {
	fmt.Printf("fuel: [%s], carType: [%s]\n", e.fuel, e.carType)
}
