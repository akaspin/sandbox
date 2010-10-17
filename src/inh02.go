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
func (self Root) Run(msg string) {
    self.msg = msg
    self.self.do1()
}
func (self *Root) do1() {
    log.Print("Root:", self.msg)
}
func (self *Root) do2() {
    self.do1()
}

// Inherit
type In1 struct {
    Root
}
func (self *In1) do1() {
    log.Print("In1:", self.msg)
}

func main () {
    in1 := In1{}
    in1.self = &in1
    in1.Run("inDo1")
    in1.do2()

}