package main

import (
    "log"
)

// Tester interface
type Tester interface {
    Run(msg string)
}

// Root
type Root struct {
    msg string
}
func (self *Root) Run(msg string) {
    self.msg = msg
    self.do()
}
func (self *Root) do() {
    log.Print("Root:", self.msg)
}

// Inherit
type In1 struct {
    *Root
}
func (self *In1) do() {
    log.Print("In:", self.msg)
}

func main () {
    root := new(Root)
    in1 := &In1{&Root{}} 
    
    in1.Run("one")
    root.Run("root")
}