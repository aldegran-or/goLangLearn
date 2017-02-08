package main
import "fmt"

func f(y int) int{
	return y + 1
}

func main(){
	x := 4
	fmt.Println(f(x))
}
