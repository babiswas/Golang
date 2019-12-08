package main
import "fmt"

func Sum(a,b int) int{
    sum:=0
    sum=a+b
    return sum
}

func main(){
   val:=Sum(4,5)
   fmt.Println(val)
}