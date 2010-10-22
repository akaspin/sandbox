// Not works
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
    Data string  // Some variable data
}
func (b *B) Handle() {
    b.Do()
}
func (b *B) Do() {
    log.Print("B: ", b.Data)
}

// Child
type C struct {
    B
}
func (c *C) Do() {
    log.Print("C1: ", c.Data)
}

func main () {
    h := &C{B{"data"}} // Why not &C{"data"}?
    h.Handle()
}
