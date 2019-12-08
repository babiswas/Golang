package main
import "fmt"

func main(){
  str1:="hello"
  str2:=""
  for _,ch:=range str1{
      str2=str2+string(ch)
  }
  fmt.Println(str2)
}