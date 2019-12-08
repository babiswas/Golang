package main
import "fmt"

type person struct{
  name string
  age int
}

func (p person)get_name()string{
   return p.name
}

func (p person)get_age()int{
   return p.age
}

func (p person)get_name_age()(string,int){
   return p.name,p.age
}

func main(){
 p:=person{"Ajit",32}
 fmt.Println(p.get_name())
 fmt.Println(p.get_age())
 name,age:=p.get_name_age()
 fmt.Println(name)
 fmt.Println(age)
}