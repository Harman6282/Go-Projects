package main

import "fmt"

func main(){
	todos := Todos{}
	todos.add("Learn React native")
	todos.add("Learn devops")
	fmt.Printf("%+v\n\n", todos)
	todos.delete(0)
	fmt.Printf("%+v", todos)
}