package main

import "fmt"


func main() {

	var name string
    var university string
    var age string
    var hobby string

	fmt.Println("Write your name : ")
	fmt.Scan(&name)

    fmt.Println("Write your university : ")
	fmt.Scan(&university)

    fmt.Println("Write your age : ")
	fmt.Scan(&age)

    fmt.Println("Write your hobby : ")
	fmt.Scan(&hobby)

    v := fmt.Sprintf("Hello my name is %s from %s my age %s and my hobby %s",name,university,age,hobby)

    fmt.Println(v)
}

