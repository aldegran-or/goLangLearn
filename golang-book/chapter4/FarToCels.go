package main

import "fmt"

func main(){
  fmt.Print("Введите температуру (Фарингейт): ")
  var input_farTemp float64
  fmt.Scanf("%f", &input_farTemp)

  output_celTemp := (input_farTemp - 32) * 5 / 9
  fmt.Println("Температура в Цельсиях:" ,output_celTemp)
}
