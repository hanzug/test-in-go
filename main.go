package main

import "fmt"

type T struct {
	value int
}

func (p *T) Modify() {
	p.value = 10
}

func main() {
	t := T{value: 5}
	t.Modify() // 内部调用 t 的地址 &t.Method()
	fmt.Print(t)
}
