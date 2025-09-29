package main
import (
	"fmt"
	
)

func main() {
	myToy := "teddy bear"
	fmt.Println(myToy)

	toyAddress := &myToy
	fmt.Println(toyAddress)
	fmt.Println(*toyAddress)
}