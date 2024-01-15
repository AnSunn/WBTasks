package pattern

/*
	Implementation of "builder" pattern - creational pattern (https://en.wikipedia.org/wiki/Builder_pattern)

Applicability of builder pattern:
1. Complex Object Construction
2. Step-by-Step Construction
3. Variability in Object Representation

Pros:
1. Separation of Concerns
It separates the construction of a complex object from its representation, providing better organization
and maintainability of the code.
2. Step-by-Step Construction
3. Reusability
The individual components (ConcreteBuilders) can be reused across different director classes, promoting code reuse.
4. Encapsulation
The details of the construction process are encapsulated within the ConcreteBuilder classes, providing a clean and
clear interface for clients

Cons:
1. Complexity
Introducing the Builder pattern may increase the overall complexity of the codebase, especially for simpler objects
where the benefits of the pattern might be unnecessary.
2. Overhead
In cases where there are only a few variations of the final object, using the Builder pattern might be overkill,
leading to additional code overhead.

Real examples of using builder pattern:
1. when dealing with the construction of complex objects (meal builder, vehicle builder, etc.)
2. document building
In a word processing application, when creating a complex document, the Builder pattern can be employed.
Different builders can be used for creating documents with various structures (e.g., reports, letters, or articles).
*/
type Toy struct {
	model    string
	material string
	form     string
}

// PlasticBuilder - create plastic builder and its functions
type PlasticBuilder struct {
	toy Toy
}

func NewPlasticBuilder() *PlasticBuilder {
	return &PlasticBuilder{}
}

func (p *PlasticBuilder) setModel(str string) {
	p.toy.model = str
}

func (p *PlasticBuilder) setMaterial() {
	p.toy.material = "plastic"
}
func (p *PlasticBuilder) setForm(str string) {
	p.toy.form = str
}
func (p *PlasticBuilder) getToy() Toy {
	return Toy{
		model:    p.toy.model,
		material: p.toy.material,
		form:     p.toy.form,
	}
}

// WoodBuilder - create wood builder and its functions
type WoodBuilder struct {
	toy Toy
}

func NewWoodBuilder() *WoodBuilder {
	return &WoodBuilder{}
}

func (p *WoodBuilder) setModel(str string) {
	p.toy.model = str
}

func (p *WoodBuilder) setMaterial() {
	p.toy.material = "wood"
}
func (p *WoodBuilder) setForm(str string) {
	p.toy.form = str
}
func (p *WoodBuilder) getToy() Toy {
	return Toy{
		model:    p.toy.model,
		material: p.toy.material,
		form:     p.toy.form,
	}
}

type builder interface {
	setModel(str string)
	setMaterial()
	setForm(str string)
	getToy() Toy
}

// Director struct to manage builders
type Director struct {
	builder builder
}

func NewDirector(build builder) *Director {
	return &Director{
		builder: build,
	}
}
func (d *Director) Construct() Toy {
	d.builder.setModel("AW")
	d.builder.setMaterial()
	d.builder.setForm("car")
	return d.builder.getToy()
}
