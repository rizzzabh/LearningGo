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

type MyError struct {
	Code int 
	Msg string 
}

func (e MyError) Error() string {
	return fmt.Sprintf("Error code : %f and message : %s" , e.Code , e.Msg)
}

func risky() error {
	return MyError{Code: 404 , Msg: "Not Found"}
}
func main(){
	result , err := divide(10 , 2)
	risk := risky() 

	if risk != nil {
		fmt.Println("Oops , " , risk)
	}
	if err != nil {
		fmt.Println("Error: " , err) 
	}else {
		fmt.Println("Result: " , result)
	}
}