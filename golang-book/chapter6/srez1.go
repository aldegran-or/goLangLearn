package main
import "fmt"
func main() {
  arr := [5]int{0,1,2,3,4,}

  fmt.Println(arr[0])
  fmt.Println(arr[1])
  fmt.Println(arr[2])
  fmt.Println(arr[3])
  fmt.Println(arr[4])

  x := arr[2:] // [2,3,4]
  fmt.Println(len(arr))
  fmt.Println(x)
}
