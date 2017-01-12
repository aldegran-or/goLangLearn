package main

import "fmt"

func main(){
  for i := 1; i <= 10; i++ {
    switch i {
    case 1: fmt.Println(i, "один")
    case 2: fmt.Println(i, "два")
    case 3: fmt.Println(i, "три")
    case 4: fmt.Println(i, "четыре")
    case 5: fmt.Println(i, "пять")
    case 6: fmt.Println(i, "шесть")
    case 7: fmt.Println(i, "семь")
    case 8: fmt.Println(i, "восемь")
    case 9: fmt.Println(i, "девять")
    case 10: fmt.Println(i, "десять")
    }
  }
}
