package main
import "fmt"
func main(){
	x := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
		} 
	val := len(x)
	min := x[0]
	for i := 1; i <= val-1; i++ {
			fmt.Println("итерация", i, x[i],min)	
		if x[i] < min { 
			min = x[i]
			}
	}
	fmt.Println("кол-во Элементов массива:",val)
	fmt.Println("Минимальное число:",min)
}
