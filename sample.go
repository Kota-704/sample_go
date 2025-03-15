package main

import "fmt"

type User struct {
	name string
}

func (u User) cal(weight, height float64) (result float64) {
	result = weight / height / height * 10000
	return
}

func main () {
	user001 := User {"Kota"}
	fmt.Println(user001.name, user001.cal(60, 171))
}
