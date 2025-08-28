package main 
 
import ("fmt")


func main(){
	slice := make([]int , 2 , 3)

	slice = append(slice , 4 , 5) 

	fmt.Print(slice)
	arr := [6]int{1 ,2 ,3 , 4 , 5 , 6}

	slice2 := arr[1 : 4]

	fmt.Println(slice2)

	for i , v := range arr {
		fmt.Println(i , v)
	}
	
}
