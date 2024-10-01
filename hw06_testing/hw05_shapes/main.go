package main

import (
	"errors"
	"fmt"
	"math"
)

type Shape interface {
	Square() float32
}

type Circle struct {
	radius float32
}

func (c *Circle) Radius() float32 {
	return c.radius
}

func (c *Circle) SetRadius(radius float32) {
	c.radius = radius
}

func NewCircle(radius float32) *Circle {
	circle := Circle{
		radius: radius,
	}
	return &circle
}

type Rectangle struct {
	width  float32
	height float32
}

func (r *Rectangle) Width() float32 {
	return r.width
}

func (r *Rectangle) Height() float32 {
	return r.height
}

func (r *Rectangle) SetWidth(width float32) {
	r.width = width
}

func (r *Rectangle) SetHeight(height float32) {
	r.height = height
}

func NewRectangle(width, height float32) *Rectangle {
	rectangle := Rectangle{
		width:  width,
		height: height,
	}
	return &rectangle
}

type Triangle struct {
	base   float32
	height float32
}

func (t *Triangle) Base() float32 {
	return t.base
}

func (t *Triangle) Height() float32 {
	return t.height
}

func (t *Triangle) SetBase(base float32) {
	t.base = base
}

func (t *Triangle) SetHeight(height float32) {
	t.height = height
}

func NewTriangle(base, height float32) *Triangle {
	triangle := Triangle{
		base:   base,
		height: height,
	}
	return &triangle
}

type Cat struct{}

func NewCat() *Cat {
	cat := Cat{}
	return &cat
}

func (c *Circle) Square() float32 {
	return math.Pi * c.Radius() * c.Radius()
}

func (r *Rectangle) Square() float32 {
	return r.Height() * r.Width()
}

func (t *Triangle) Square() float32 {
	return t.base * t.height / 2
}

func calculateArea(s any) (float32, error) {
	var ok bool
	var shape Shape
	shape, ok = s.(Shape)

	if !ok {
		err := errors.New("переданный объект не является фигурой")
		return 0, err
	}
	square := shape.Square()

	return square, nil
}

func printArea(s any) {
	shape, err := calculateArea(s)
	if err == nil {
		fmt.Println("Площадь ", shape)
	} else {
		fmt.Println("Ошибка: ", err)
	}
}

func main() {
	fmt.Println("Круг: радиус 5")
	circle := NewCircle(5)
	printArea(circle)

	fmt.Println("Прямоугольник: ширина 10, высота 5")
	rectangle := NewRectangle(10, 5)
	printArea(rectangle)

	fmt.Println("Треугольник: основание 8, высота 6")
	triangle := NewTriangle(8, 6)
	printArea(triangle)

	fmt.Println("Кот")
	cat := NewCat()
	printArea(cat)
}
