package main

import (
    "fmt"
    "time"
    "http"
    "log"
    "runtime"
)

// Ordinal
func ordinalHandler (w http.ResponseWriter, r *http.Request)  {
    log.Printf("Ordinal ", r.URL.Path[1:])
    time.Sleep(10000000000) 
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

// Object handler
type ObjHandler struct {
    res http.ResponseWriter
    req *http.Request
    chunk string
}

func (self *ObjHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    self.res = w
    self.req = r
    self.get()
}

func (self *ObjHandler) get() {
    log.Print("Obj ", self)
    http.Error(self.res, "Not implemented", 405)
}

// oped
type OpHandler struct {
    ObjHandler
}

func (self *OpHandler) get() {
    log.Printf("Op ", self)
    time.Sleep(2000000000) 
    fmt.Fprintf(self.res, "obj %v", self)
}

func main() {
    runtime.GOMAXPROCS(10)

    objHandler := new(ObjHandler)
    opHandler := new(OpHandler)
    http.Handle("/obj/", objHandler)
    http.Handle("/op/", opHandler)
    
    //var obj2Handler ObjHandler 
    //obj2Handler = ObjHandler{}
    http.Handle("/obj2/", &ObjHandler{})
    http.Handle("/op2/", &OpHandler{})
    
    http.HandleFunc("/ordinal/", ordinalHandler)
    http.ListenAndServe(":8080", nil)
}