package main

import "fmt"

func main(){
  fmt.Print("Введите длинну (Футы): ")
  var input_feet float64
  fmt.Scanf("%f", &input_feet)

  output_meters := 0.3048 * input_feet
  fmt.Println("Длинна в метрах:" ,output_meters)
}
