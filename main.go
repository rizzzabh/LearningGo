package main 

import "fmt"

type Animal interface {
	eat()
	haveSex()
}
type Person struct {
	name string 
	age int 
	language string
}

func (p Person) speak() {
	 fmt.Printf("I speak %s language" , p.language) 
}

func (p Person) eat(){
	fmt.Println("I only eat")
}
func (p Person) haveSex(){
	fmt.Println("I do have sex")
}

func main(){
	p1 := Person{name : "Rishabh" , age : 19 , language : "Hindi"}
	p1.eat()
	var prasann Animal
	prasann = p1 
	_ , ok := prasann.(Person)
	if ok {
		fmt.Println("Ok ho gaya")
	}else{
		fmt.Println("nahi kar raha hai")
	}
}