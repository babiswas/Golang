package main
import "fmt"
var N int

func showN() int{
  if N==0{
    fmt.Println("N didn't increase")
  }
  N=N+1
  return N
}

func main(){
 for i:=0;i<4;i++{
   m:=showN()
   fmt.Println(m)
 }
}
