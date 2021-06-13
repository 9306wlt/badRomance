package method

import (
	"fmt"
	"math"
	"testing"
)

/*
内嵌类型的方法和继承
当一个匿名类型被内嵌在结构体中时，匿名类型的可见方法也同样被内嵌，这在效果上等同于
外层类型继承了这些 方法：将父类型放在子类型中来实现。

内嵌将一个已存在类型的字段和方法注入到另一个类型里：匿名字段上的方法"晋升"成为了外层
类型的方法。当然类型可以有只作用于本身实例而不作用于内嵌“父”类型上的方法
可以覆写方法（像字段一样）：和内嵌类型方法具有同样名字的外层类型的方法会覆写内嵌类型对应的方法。
 */

type Point struct {
	x,y float64
}

func(p *Point)Abs()float64{
	return math.Abs(p.x*p.x + p.y*p.y)
}

func(n *NamedPoint)Abs() float64{
	return n.Point.Abs() * 100
}

type NamedPoint struct {
	Point
	name string
}

func Test_method3(t *testing.T){
	n := &NamedPoint{Point{3,4},"Pythagoras"}
	fmt.Println(n.Abs())
}