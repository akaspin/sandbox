package main

import (
    "log"
)

// Tester interface
type Tester interface {
    Run(msg string) 
    do1()
    do2()
}

// Root
type Root struct {
    self Tester
    msg string
}
func (this Root) Run(msg string) {
    this.msg = msg
    this.self.do1()
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
func (this *In1) do1() {
    log.Print("In1:", this.msg)
}

// Inherit2
type In2 struct {
    Root
}
func (this *In2) do1() {
    log.Print("In2:", this.msg)
}
func (this *In2) do2() {
    this.self.do1()
}


func main () {
    in1 := new(In1)
    in1.self = in1
    in1.Run("inDo1")
    in1.do2()

    in2 := new(In2)
    in2.self = in2
    in2.Run("inDo2")
    in2.do2()
    
    

}