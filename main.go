package main 
 
import ("fmt")

func variadic (greeting string , values ...interface {}){
	for _ , v := range values {
		fmt.Println(greeting , v) 
	}
}
func main(){
	variadic("Hello" , 10 , 3.15 , "Rishabh")
}
