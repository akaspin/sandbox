/**  Handler-Doer pattern
 
    Based on idea of dynamically created handler that
    keeps next 
*/ 

package main

import ( 
    "log"
    "strconv"
)

/** Handler is dynamicaly created container. 
    It calls Doer methods and intercept all panics
    Handler keeps next set of things:
    
    * env - static enviroment
    * input - input data that reezes in handler
    * data - dynamic data storage for Doer
    * Doer
    * error handling function
    
    Handler don't process any data. All data must be 
    processed by Doer. 
*/
type Handler interface {
    // Sets static enviroment. Doer and enviroment.
    // Dynamic data handles by doer.
    // Calls only on creation.
    // Can be inherited
    Init(Doer, func(Handler, int), interface{}, interface{})
    
    // Get Doer. Calls only by Handler
    getDoer() Doer
    
    // Get Doer. Calls only by Handler
    getEH() func(Handler, int)
    
    // Get inpit data
    Input() interface{}
    
    // Get dynamic data. Calls only by Doer and Error function
    Data() interface{}
    
    // Get static enviroment
    Env() interface{}

    // Set dynamic data. Calls only by Doer
    SetData(interface{})
    
    // Handle. Calls for process.
    // Must be inherited
    Handle()
}

/** Doer processes all data. */
type Doer interface {
    // Doer methods
    DoOne(h Handler)
    DoTwo(h Handler)
}

// implementations

// Base handler
type Handle struct {
    env interface{}
    data interface{}
    input interface{}
    doer Doer
    err func(Handler, int)
}
func (b *Handle) Init(d Doer, eh func(Handler, int), 
        input interface{}, env interface{}) {
    b.doer = d
    b.err = eh
    b.input = input
    b.env = env
    
}
func (b *Handle) getDoer() Doer { return b.doer }
func (b *Handle) getEH() func(Handler, int) { return b.err }
func (b *Handle) Input() interface{} { return b.input }
func (b *Handle) Data() interface{} { return b.data }
func (b *Handle) Env() interface{} { return b.env }
func (b *Handle) SetData(data interface{} ) {b.data = data}
func (b *Handle) Handle() {
    // Set recover
    defer func() {
        if e := recover(); e != nil {
            b.getEH()(b, e.(int))
	        log.Print("Recovered")
        }
    }()

    log.Printf("%s handle. Iteration %s", b.Env(), b.Input())

    // And go
    b.getDoer().DoOne(b)
    b.getDoer().DoTwo(b)
}

type Do1 struct {}
func (d *Do1) DoOne(h Handler) {
    h.SetData(h.Input().(string) + "-1")
    log.Printf("  Do1-One: %s", h.Data())
}
func (d *Do1) DoTwo(h Handler) {
    h.SetData(h.Data().(string) + "-1")
    if (h.Input().(string) == "1") {
        panic(1)
    }
    log.Printf("  Do1-Two: %s", h.Data())
}
// Inherited
type Do2 struct { Do1 }
func (d *Do2) DoTwo(h Handler) {
    h.SetData(h.Data().(string) + "-2")
    if (h.Input().(string) == "2") {
        panic(2)
    }
    log.Printf("  Do2-Two: %s", h.Data())
}


func main () {
    // recover func
    rec := func(h Handler, err int) {
        log.Printf("Some panic on iteration %d", err)
    }
    // doers
    d1 := &Do1{}
    d2 := &Do2{}
    
    var h1, h2 Handler
    for i := 0; i<2; i++ {
        h1 = &Handle{}
        h1.Init(d1, rec, strconv.Itoa(i), "First")
        h1.Handle()
        
        h2 = &Handle{}
        h2.Init(d2, rec, strconv.Itoa(i), "Second")
        h2.Handle()
    }
    
}
