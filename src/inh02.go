package main

import (
    "log"
)

// Tester interface
type Doer interface {
    Init(msg string) 
    do1()
    do2()
}

// Root
type Root struct {
    msg string
    self Doer
}
func (this *Root) Init(msg string ) {
    this.msg = msg
}
func (this *Root) do1() {
    log.Print("Root:", this.msg)
}
func (this *Root) do2() {
    this.self.do1()
}

// Inherit1
type In1 struct {
    Root
}
func (this *In1) Init(msg string ) {
    this.self = this
    this.Root.Init(msg)
}
func (this *In1) do1() {
    log.Print("In1:", this.msg)
}

// Inherit2
type In2 struct {
    In1
}
func (this *In2) Init(msg string ) {
    this.self = this
    this.Root.Init(msg)
}
func (this *In2) do1() {
    log.Print("In2:", this.msg)
}

func main () {
    in1 := In1{}
    in1.Init("inDo1")

    in2 := In2{}
    in2.Init("inDo2")
    
    in1.do2()
    in2.do2()
}