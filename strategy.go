package main

import "fmt"

type Print interface {
	println(int)
}

type ForLoop struct {
}

type Recursion struct {
}

func (f *ForLoop) println(i int) {

	for j := i; j > 0; j-- {
		println(j)
	}
}

func (r *Recursion) println(i int) {
	if i == 0 {
		return
	}
	println(i)
	r.println(i - 1)
}

func main() {
	var p Print
	fmt.Println("ForLoop")
	p = &ForLoop{}
	p.println(10)

	fmt.Println("Recursion")
	p = &Recursion{}
	p.println(10)
}
