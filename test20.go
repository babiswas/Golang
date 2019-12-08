package main
import "fmt"

func main(){
  s:=make([]int,3)
  fmt.Println(s)
  s[0]=1
  s[1]=2
  s[2]=3
  fmt.Println(s)
  s=append(s,10,11,12)
  fmt.Println(s)
  fmt.Println("After slicing")
  fmt.Println(s[3:6])
  x:=s
  fmt.Println(x)
}