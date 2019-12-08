package main
import "fmt"

func main(){
  var num[3] int
  for i:=0;i<len(num);i++{
   num[i]=i+100
  }
  for i:=0;i<len(num);i++{
    fmt.Println(num[i])
  }
}