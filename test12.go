package main
import "fmt"

func Sum(a int,b int){
   sum:=0
   sum=a+b
   fmt.Println(sum)
}

func main(){
   Sum(2,3)
}