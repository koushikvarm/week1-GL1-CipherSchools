package main
import "fmt"

func main() {
	// i:=10
	// j:=&i
	var i int
	i=10
	var j *int // declaring..... no value is arrranged to j 
	//j:=new(int) // declaring+ assignment
	j=&i
	*j=100
	name := new(string)
	*name = "golang"
	fmt.Println(*name)
	
}
