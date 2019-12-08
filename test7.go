package main
import "fmt"

func main(){
  i:=0
  str1:="Hello"
  for{
    fmt.Println(str1)
    i++
    if i==10{
       break
    }
  }
}