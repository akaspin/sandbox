package main

import (
   "log"
   "reflect"
)

type Doer interface {
    Init(msg string)
    do()
}

type Root struct {
    msg string
}
func (self *Root) Init(msg string) {
    self.msg = msg
    log.Print(reflect.Typeof(self))
    self.do()
}
func (self *Root) do() {
    log.Print("Root:", self.msg)
}

// Inherit
type In1 struct {
    Root
}
func (self *In1) do() {
    log.Print("In:", self.msg)
}

func Handle(h Doer, msg string ) {
   h.Init(msg)
}

func main () {
    root := Root{}
    in := In1{} 
    
    Handle(&root, "root")
    Handle(&in, "one")
}