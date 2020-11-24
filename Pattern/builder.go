package pattern

import (
	"fmt"
)



type Item interface {
	Name() string
	Price() float64
	Packing() Packing
}
type Packing interface {
	Pack() string
}

type Wrapper struct {
	Packing
}

func (w *Wrapper) Pack() string {
	return "wrapper"
}

type Bottle struct {
	Packing
}

func (b *Bottle) Pack() string {
	return "wrapper"
}

type Bugger struct {
	Item
}

func (b *Bugger) Packing() Packing {
	return new(Wrapper)
}

type VegBugger struct {
	Bugger
}

func (v *VegBugger) Price() float64 {
	return 25
}
func (v *VegBugger) Name() string {
	return "vegBugger"
}

type ChickenBugger struct {
	Bugger
}

func (c *ChickenBugger) Price() float64 {
	return 50.5
}
func (c *ChickenBugger) Name() string {
	return "ChickenBugger"
}

type ColdDrink struct {
	Item
}

func (c *ColdDrink) Packing() Packing {
	return new(Bottle)
}

type Coke struct {
	ColdDrink
}

func (c *Coke) Price() float64 {
	return 30
}
func (c *Coke) Name() string {
	return "Coke"
}

type Pepsi struct {
	ColdDrink
}

func (p *Pepsi) Price() float64 {
	return 35
}
func (p *Pepsi) Name() string {
	return "Pepsi"
}

type Meal struct {
	Items []Item
}

func (m *Meal) AddItem(item Item) {
	m.Items = append(m.Items, item)
}
func (m *Meal) GetCost() float64 {
	var cost float64
	for _, item := range m.Items {
		cost += item.Price()
	}
	return cost
}
func (m *Meal) ShowItems() {
	for _, item := range m.Items {
		fmt.Println("Item : ", item.Name())
		fmt.Print(", Packing : ", item.Packing().Pack())
		fmt.Println(", Price : ", item.Price())
	}
}

type MealBuilder struct {
}

func (m *MealBuilder) prepareVegMeal() *Meal {
	meal := new(Meal)
	meal.AddItem(new(VegBugger))
	meal.AddItem(new(Coke))
	return meal
}
func (m *MealBuilder) prepareNonVegMeal() *Meal {
	meal := new(Meal)
	meal.AddItem(new(ChickenBugger))
	meal.AddItem(new(Pepsi))
	return meal
}
func builderTest() {
	mealBuilder := new(MealBuilder)
	vegMeal := mealBuilder.prepareVegMeal()
	fmt.Println("vegMeal")
	vegMeal.ShowItems()
	fmt.Println("Total Cost: " ,vegMeal.GetCost())
	nonVegMeal := mealBuilder.prepareVegMeal()
	fmt.Println("nonVegMeal")
	nonVegMeal.ShowItems()
	fmt.Println("Total Cost: " ,nonVegMeal.GetCost())
}
