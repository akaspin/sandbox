	package main
	
	import ( 
	    "log"
	)
	
	type Handler interface {
	    Handle()
	}
	
	type Doer interface {
	    Do(*Handler)    // Calls Serv
	    Serv()  // Service function
	}
	
	// Handler. Keeps data and runs Doer.Do() 
	type Handle struct {
	    Data string
	    D Doer
	}
	func (h *Handle) Handle(){
	    h.D.Do(h)
	}
	
	// Doer one. Not keeps any data. 
	// Calls other self methods only through Handle.D 
	type Do1 struct {}
	func (d *Do1) Do(h *Handler) {
	    log.Print("Do1", h.Data)
	    h.D.Serv()
	}
	func (d *Do1) Serv() {
	    log.Print("Serv1")
	}
	
	type Do2 struct { Do1 }
	func (d *Do2) Serv() {
	    log.Print("Serv2")
	}
	
	func main () {
	    h1 := &Handle{"data1", &Do1{}}
	    h2 := &Handle{"data2", &Do2{}}
	
	    h1.Handle()
	    h2.Handle()
	}
	
	/* not works
	../src/main.go:22: cannot use h (type *Handle) as type *Handler in function argument
	../src/main.go:29: h.Data undefined (type *Handler has no field or method Data)
	../src/main.go:30: h.D undefined (type *Handler has no field or method D)
	*/
	
	

