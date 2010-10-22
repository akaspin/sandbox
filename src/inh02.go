// works
package main

import ( 
    "log"
)

type Handler interface {
    Handle() // Calls Do()
    Do()
}

// Base handler
type B struct {
    self Handler
    Data string  // Some variable data
}
// Fake constructor
func (b *B) Init(data string) {
    b.Data = data
}
func (b *B) Handle() {
    b.self.Do()
}
func (b *B) Do() {
    log.Print("B: ", b.Data)
}

// Child
type C struct {
    B
}
// Need set self in each child
func (c *C) Init(data string) {
    c.B.Init(data)  // Call base fake constructor
    c.self = c      // set instance
}
func (c *C) Do() {
    log.Print("C1: ", c.Data)
}
// Child
type C1 struct {
    C
}
// Need set self in each child
func (c *C1) Init(data string) {
    c.C.Init(data)  // Call base fake constructor
    c.self = c      // set instance
}
func (c *C1) Do() {
    log.Print("C2: ", c.Data)
}

func main () {
    h := &C{}
    h.Init("data1")
    h.Handle()
    
    h1 := &C1{}
    h1.Init("data1")
    h1.Handle()
}
