/*
    GC example. GC not work:
    http://groups.google.com/group/golang-nuts/browse_thread/thread/c5dc1a44c5bff95a
*/

package main

import ( 
    "time"
    "log"
)

type Big struct {
    Data string 
}

func biggy () {
    a := &Big{}
    a.Data = "Very very long string"
}

func main() {
    log.Print("Prepare")
    time.Sleep(10000000000)
    log.Print("Trashing")
    for i := 0; i<100000000; i++ {
        biggy()
    }
    log.Print("All done")
    for {
        time.Sleep(1000000)
    }
}