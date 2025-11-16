package main

import "fmt"

func main(){
	nums := []int{1,2,3,4,5,6,7,8,9,10}

	input := make(chan int)
	doubled := make(chan int)

	go func ()  {
		defer close(input)
		for _, num := range nums{
			input <- num
		}
	}()

	go func ()  {
		defer close(doubled)
		for num := range input{
			doubled <- num*2
		}
	}()

	fmt.Println("Результаты умножения на 2:")
	for result :=range doubled{
		fmt.Println(result)
	}
}

