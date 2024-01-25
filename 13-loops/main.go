package main

import "fmt"

func main() {
	fmt.Println("Loops")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)
	
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Println(index, day)
	}

	rougeValue := 0

	for rougeValue < 10 {

		if rougeValue == 2 {
			goto lco
		} 

		if rougeValue == 5 {
			rougeValue++
			continue
		} else if rougeValue == 8 {
			break
		}

		fmt.Println(rougeValue)
		rougeValue++
	}

lco:
	fmt.Println("LCO")


}