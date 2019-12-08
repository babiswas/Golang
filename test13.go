package main
import "fmt"

func giggler(a,b int) (int,int){
   return a+100,b+100
}

func main(){
   m,n:=giggler(3,4)
   fmt.Println(m)
   fmt.Println(n)
}
