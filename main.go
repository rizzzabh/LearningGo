package main 
 
import ("fmt")

func main(){

	var num int 
	fmt.Scan(&num)
	for i := 1 ; i <= num ; i++ {
		isPrime := true
		for j := 2 ; j*j <= i ; j++ {
			if i % j == 0 {
				isPrime = false 
				break 
			}
		}
		if isPrime {
			fmt.Print(i , " ") 
		}
	}
	fmt.Println()
}