package main 

import "fmt"

type car struct {
	make string 
	model string 
}
type truck struct {
	car 
	bigWheel int 
}
func main(){
	// const name = "Rishabh" 
	// const surName = "Mishra" 

	// msg := fmt.Sprintf("Hi %v %v, welcome to my Club" , name , surName) 

	// fmt.Println(msg)

	truckDriver := truck{} 

	truckDriver.model = "AmitabhBachhan" 

	fmt.Println(truckDriver.model)
}