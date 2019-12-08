package main
import ( 
      "fmt"
      "strconv"
)

func itoA(a int)string{
 return strconv.Itoa(a)+"hello"
}

func main(){
  m:=itoA(4)
  fmt.Println(m)
}