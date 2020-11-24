package pattern

import "fmt"


/*
模式名称：桥接模式
目的（What）：将抽象部分与实现部分分离，使它们都可以独立的变化
解决的问题（Why）：在有多种可能会变化的情况下，用继承会造成类爆炸问题，扩展起来不灵活。当实现系统可能有多个角度分类（多个抽象类），每一种角度会有不同的变化。
解决方案（How）：把每个角度的分类分离出来，让它们独立变化，减少它们之间的耦合
解决效果：
优点：
抽象和实现的分离
优秀的扩展能力
实现细节对客户透明
缺点：
桥接模式的引入会增加系统的理解与设计难度，由于聚合关联关系（桥接）建立在抽象层，要求开发者针对抽象进行设计与编程

*/
type car interface {
	name() string
	run()
}

type MercedesBenz struct {
}

func (m MercedesBenz) name() string {
	return "Mercedes Benz"
}
func (m MercedesBenz) run() {
	fmt.Println("benz running")
}

type Audi struct {
}

func (a Audi) name() string {
	return "Audi"
}
func (a Audi) run() {
	fmt.Println("audi running")
}

type shiftMode interface {
	up()
	down()
}
type Manual struct {
	car car
}

func (m Manual) up() {
	fmt.Printf("Manual up the %s car\n", m.car.name())
	m.car.run()
}
func (m Manual) down() {
	fmt.Printf("Manual down the %s car\n", m.car.name())
	m.car.run()
}
func getNewManualShift(c car) *Manual {
	return &Manual{car: c}
}

type Automatic struct {
	car car
}

func (a Automatic) up() {
	fmt.Printf("Automatic up the %s car\n", a.car.name())
	a.car.run()
}
func (a Automatic) down() {
	fmt.Printf("Automatic down the %s car\n", a.car.name())
	a.car.run()
}
func getNewAutoShift(c car) *Automatic {
	return &Automatic{car: c}
}

func TestBridge() {
	ma := getNewAutoShift(&MercedesBenz{})
	mm := getNewManualShift(&MercedesBenz{})
	aa := getNewAutoShift(&Audi{})
	am := getNewManualShift(&Audi{})

	ma.up()
	ma.down()
	println("+++++++++++++++++++++++++")
	mm.up()
	mm.down()
	println("+++++++++++++++++++++++++")
	aa.up()
	aa.down()
	println("+++++++++++++++++++++++++")
	am.up()
	am.down()
	println("+++++++++++++++++++++++++")

}
