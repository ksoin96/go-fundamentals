package main
import "fmt"

func main(){
     var x int
     fmt.Print("What’s your first number? ")
     fmt.Scan(&x)
     
     var y int
     fmt.Print("What's your second number? ")
     fmt.Scan(&y)

     fmt.Println(x+y)
     
}