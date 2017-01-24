//Массивы, срезы, карты
package main
import "fmt"
func main (){
		arr := [10]float64{
		98,
		93,
		95,
		91,
		87,
		78,
		83,
		55,
		71,
		67,
		}
		x := arr[5:]
		fmt.Println(x)
}
