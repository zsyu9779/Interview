package pattern

import "fmt"

// 工厂方法模式
type PType int

const (
	ProductA PType = 1
	ProductB PType = 2
)

type Factory struct {
}

func (f *Factory) Generate(t PType) Product {
	switch t {
	case ProductA:
		return &Product1{}
	case ProductB:
		return &Product2{}
	default:
		return nil
	}

}

type Product interface {
	create()
}
type Product1 struct {
}

func (p *Product1) create() {
	fmt.Println("create product1")
}

type Product2 struct {
}

func (p *Product2) create() {
	fmt.Println("create product2")
}

func FactoryTest() {
	factory := new(Factory)
	p1 := factory.Generate(ProductA)
	p1.create()
	p2 := factory.Generate(ProductB)
	p2.create()

}

// 抽象工厂模式
type BigFactory interface {
	Generate(t PType) Product
}
type expensiveFactory struct {
}

func (e *expensiveFactory) Generate(t PType) Product {
	switch t {
	case ProductA:
		fmt.Println("generate expensive product")
		return &Product1{}
	case ProductB:
		fmt.Println("generate expensive product")
		return &Product2{}
	default:
		return nil
	}

}

type cheapFactory struct {
}

func (c *cheapFactory) Generate(t PType) Product {
	switch t {
	case ProductA:
		fmt.Println("generate cheap product")
		return &Product1{}
	case ProductB:
		fmt.Println("generate expensive product")
		return &Product2{}
	default:
		return nil
	}

}

// todo:正常来说差异应该体现在product层面，没想好怎么表示，稍后改进
func AbstractFactoryTest() {
	expensiveFactory := new(expensiveFactory)
	cheapFactory := new(cheapFactory)

	p1 := expensiveFactory.Generate(ProductA)
	p1.create()
	p2 := expensiveFactory.Generate(ProductB)
	p2.create()
	fmt.Println("==========================")
	p3 := cheapFactory.Generate(ProductA)
	p3.create()
	p4 := cheapFactory.Generate(ProductB)
	p4.create()

}
