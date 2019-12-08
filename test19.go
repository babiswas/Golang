package main
import "fmt"

func main(){
  var x=100
  var p=&x
  var q=&p
  **q=300
  fmt.Println(x)
  return
}