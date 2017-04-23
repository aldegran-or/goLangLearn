package main

import (
//	"fmt"
)
func main(){
    var i = 0
    for i = 1; i < 11; i++ {
         if i%2 == 1 {
            continue
         }
        println(i)
    }
}
