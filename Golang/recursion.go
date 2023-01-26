// recursion --> a function calling itself untill base condition is satisfied.
package main
import "fmt"

func main() {
	// fmt.Println(fib(4))
	// fmt.Println(fact(4))
	rec(5)
	
}
func rec(num int) {
	if num <=0 {
		return
	}
	rec(num-1)
	rec(num-2)
	fmt.Println(num-1)
	// if(num%2==0){
	// 	rec(num-1)
	// 	fmt.Println(num-1)
	// }else{
	// 	rec(num-1)
	// 	fmt.Println(num-1)
	// }
	
}
func fib(number int) int {
	if number ==0|| number ==1 {
		return number
	}

	result:=fib(number-1)+fib(number-2)
	return result
}


func fact(number int) int {
	if number==1 || number ==0 {
		return 1
	}
	if number < 0{
		fmt.Println("Invalid number")
		return -1
	}
	result:=number*fact(number-1)
	return result
}

