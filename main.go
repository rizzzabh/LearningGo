package main 
 
import ("fmt")

func variadic (greeting string , values ...interface {}){
	for _ , v := range values {
		fmt.Println(greeting , v) 
	}
}
func sum (nums ...int) int {
	s := 0 
	for _ , value := range nums {
		s += value 
	}
	return s 
}
func main(){
	slices := []int {1 , 2 , 3 , 4}
	totSum := sum (slices...)

	fmt.Println(totSum)
	variadic("Hello" , 10 , 3.15 , "Rishabh")
}
