//Массивы, срезы, карты
package main
import "fmt"
func main (){
		x := [5]float64{
		98,
		93,
		95,
		91,
		87,
		}
		var total float64
		total = 0
		
		for _, value := range x {
			total = total + value
			}
		fmt.Println( total / float64(len(x)) )
		
}
