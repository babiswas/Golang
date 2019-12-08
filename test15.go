package main
import "fmt"

func multiply(a,b int) int{
   mul:=a*b
   return mul
}

func main(){
   m,n:=2,3
   res:=multiply(m,n)
   fmt.Println(res)
}

