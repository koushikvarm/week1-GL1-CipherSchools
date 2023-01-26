package main
import"fmt"

func main() {
	// fmt.Print("Enter a Number: ")
	// var input int
	// fmt.Scanln(&input)

	// if input%2 == 0 {
	// 	fmt.Println("Number is Even")
	// } else{
	// 	fmt.Println("Number is odd")
	// }
	// if x:=10; x%2 == 0 {
	// 	fmt.Println("Number is Even")

	// }else{
	// 	fmt.Println("Number is Even")
	// }
	data := 1000
	switch data {
	case 10:
		data = 100
		fmt.Println(data)
	case 100:
		data = 101
		fmt.Println(data)
	case 1000:
		data = 102
		fmt.Println(data)
	case 10000:
		data = 103
		fmt.Println(data)
	}
	// fallthrough --> execute to next case also
}
