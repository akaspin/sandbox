package b

// changes

import "fmt"
import "./pkg/c"

func RepOne () {
    fmt.Printf("b.RepOne\n")
    c.RepOne()
    c.RepTwo()
}
