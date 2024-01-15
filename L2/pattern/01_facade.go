package pattern

import "fmt"

/*
	Implementation of "facade" pattern - structural pattern (https://en.wikipedia.org/wiki/Facade_pattern)

Applicability of facade pattern:
	1. Simplifying Complex Systems
	2. Reducing Dependency
	3. Providing High-Level Interface

Pros:
1. Simplified Interface
2. Decoupling
Facade helps in decoupling client code from the intricate
details of subsystems, which improves maintainability by reducing dependencies.
3. Promoting Consistency
It promotes consistency in interactions with the subsystems, as all communication is channeled through the facade.
4. Easier Testing and Maintenance

Cons:
1. Limited Flexibility
The main drawback is that the facade may provide a simplified interface, but it can also limit the flexibility for
advanced users who might need to interact with the subsystems more directly.
2. Additional Abstraction Layer
3.Not a One-Size-Fits-All Solution
The Facade pattern may not be suitable for every scenario. In simpler systems or when flexibility is a priority,
it might be overkill.

Real examples of using facade pattern:
1. Database Connection
In a database management system, a facade can simplify the process of connecting to different types of databases
(MySQL, PostgreSQL, etc.) by providing a unified interface for database connections.
2. Graphics Libraries
It often uses the Facade pattern to provide a high-level interface for drawing shapes, managing colors,
and handling transformations without exposing the low-level details.
3. Operating System Interfaces
Operating systems use facades to provide simplified interfaces for complex operations. For example, file system operations,
network communication, and process management can be abstracted through a system interface facade.
4. Web Frameworks
It often employs facades to simplify the interaction with various subsystems like handling HTTP requests,
managing sessions, and dealing with databases.
*/

// subservice taxi park
type taxiParkService struct {
}

func (*taxiParkService) findTaxi() {
	fmt.Println("Find the taxi in the taxi park")
}

// subservice driver
type driverService struct {
}

func (*driverService) approveOrder() {
	fmt.Println("Approve the order by driver")
}

// subservice payment
type paymentService struct {
}

func (*paymentService) blockMoney() {
	fmt.Println("Block Money on the card")
}

// subservice application
type appService struct {
}

func (*appService) notify() {
	fmt.Println("Notify customer sending car details")
}

// TaxiService facade unified interface to work with subservices
type TaxiService struct {
	taxiParkService *taxiParkService
	driverService   *driverService
	paymentService  *paymentService
	appService      *appService
}

// NewTaxi service with initialized subservices
func NewTaxi() *TaxiService {
	return &TaxiService{
		taxiParkService: &taxiParkService{},
		driverService:   &driverService{},
		paymentService:  &paymentService{},
		appService:      &appService{},
	}
}

func (t *TaxiService) OrderTaxi() {
	t.taxiParkService.findTaxi()
	t.driverService.approveOrder()
	t.paymentService.blockMoney()
	t.appService.notify()
}
