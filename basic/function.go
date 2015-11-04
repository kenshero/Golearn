package main
import "fmt"

func main() {
		sum,flag := add(10,9)
		fmt.Println(sum , flag)
}

func add(n1 int,n2 int) (int,int) {
	return n1 + n2 , 20
}