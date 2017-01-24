package main

import "fmt"
func main(){
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 {
			if i % 5 == 0 {
				fmt.Println(i, "Это число делится на 5 и 3")
				} else { fmt.Println(i, "Это число делится на 3")
					}
			
		} else if i % 5 == 0 {
			fmt.Println(i, "Это число делится на 5")
		} else { fmt.Println(i, "Это число не делится ни на 3 ни на 5")
			} 
	}
}
