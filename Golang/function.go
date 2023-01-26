package main
import "fmt"
func main(){
	w:= new(int)
	name:=new(string)
	result, x:=c(w,name)
	fmt.Println(result)
	fmt.Println(x)
	fmt.Println(*name)
	// r, _:=b(1,2,4,5,5)
	// fmt.Println(r)
}
func a() (int,string) {
	return 122,"koushik"
}
func b(args ...int) (bool,int) {
	for _,v :=range args {
		fmt.Println(v)
	}
	return true,11
	
}
func c(w *int,name *string)(i int, j string) {
	i = 10
	j="varma"
	*w=100
	*name = "star"
	return

}
