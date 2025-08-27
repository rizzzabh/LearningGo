package main 
 
import ("errors"  
	"fmt")

type error interface {
	Error() string 
}

func divide (x int , y int) (int , error ){
	if (y == 0){
		return 0 , errors.New("Divide by zero error")
	}

	return x / y , nil 
}
func main(){
	result , err := divide(10 , 2)

	if err != nil {
		fmt.Println("Error: " , err) 
	}else {
		fmt.Println("Result: " , result)
	}
}