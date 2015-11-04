package main

import "fmt"
// go  Dom't have while do  Only For !
func main() {
	// i := 0
	// for {
	// 	if i == 5000 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// 	i++
	// }	

	slice := []int{10,20,30,40}
	for _, v := range slice{
		fmt.Println(v)
	}

}