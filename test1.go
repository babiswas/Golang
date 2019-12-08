package main
import "fmt"

func main(){
  fmt.Println("Hello World")
  fun1(1,2,3)
}

func fun1(num ... int) {
  sum:=0
  for _,i:=range num{
    sum+=i
  }
  fmt.Println(sum)
}
