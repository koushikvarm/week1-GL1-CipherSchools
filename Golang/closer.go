package main
import "fmt"
func main() {
	
	var number int
	number = 10
	fmt.Println(number)

	var getInt func(int) int
    getInt = func(x int) int {
		fmt.Println("In a 1st function")
		return 20+x
	}
	getInt(1)
	getInt=func(x int) int {
		fmt.Println("In a 2nd function")
		return 10+x
	}
	
	j:= func(x int) int {
		fmt.Println("In a 1st function")
		return 20+x

	}(19)
	fmt.Println(j)
	

}
