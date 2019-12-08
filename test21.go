package main
import "fmt"

func main(){
  s:=[]int{1,2,3,4}
  fmt.Println(s)
  x:=make([]int,len(s))
  copy(x,s[0:])
  fmt.Println(x)
  x[0]=500
  fmt.Println(x)
  fmt.Println(s)
}