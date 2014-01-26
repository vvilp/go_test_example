package main

import "fmt"

type draw interface {
	draw_line()
	draw_dot()
}
//----------------------------------
type opengl struct {
	draw_type int
}

func (o *opengl) draw_line() {
	fmt.Println("opengl draw_line")
}

func (o *opengl) draw_dot() {
	fmt.Println("opengl draw_dot")
}
//----------------------------------
type directx struct {
	draw_type int
}

func (d *directx) draw_line() {
	fmt.Println("directx draw_line")
}

func (d *directx) draw_dot() {
	fmt.Println("directx draw_dot")
}
//----------------------------------

func draw_func(d draw) {
	fmt.Println("draw func")
	fmt.Println(d)
	d.draw_line()
	d.draw_dot()
}

func draw_func2(d interface{}) {
	fmt.Println("draw func 2")
	d.(draw).draw_line();
}

func main() {
	// o is pointer
	o := new(opengl)
	o.draw_line()

	d := new(directx)
	d.draw_dot()

	oo := opengl{1}

	draw_func(o)
	draw_func(d)

	//need pass the address as a pointer 
	//because : func (o *opengl) draw_line()
	draw_func(&oo)

	draw_func2(o)
}