package main
import "fmt"

func main(){
  var num [3]int
  for i:=0;i<len(num);i++{
    num[i]=200+i
  }
  sum:=0
  for _,n:=range num{
      sum=sum+n
  }
  fmt.Println(sum)
}