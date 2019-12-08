package main
import "fmt"

func main(){
  var arr[3] int
  arr[0]=22
  arr[1]=23
  arr[2]=33
  for i:=0;i<len(arr);i++{
   fmt.Println(arr[i])
  }
}