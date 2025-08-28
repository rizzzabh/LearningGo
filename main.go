package main 
 
import ("fmt")


func main(){
	slice := make([]int , 2 , 3)
	
	slice = append(slice , 4 , 5) 

	fmt.Print(slice)

}