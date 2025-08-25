package main 

import "fmt"

func main(){
	const name = "Rishabh" 
	const surName = "Mishra" 

	msg := fmt.Sprintf("Hi %v %v, welcome to my Club" , name , surName) 

	fmt.Println(msg)
}