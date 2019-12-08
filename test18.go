package main
import "fmt"

func main(){
  x:=100
  p:=&x
  fmt.Println(x)
  *p=200
  fmt.Println(x)
  return
}