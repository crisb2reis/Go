// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Car struct {
// 	CarName string `json: "car"`
// 	CarYear int    `json: "year"`
// }

// func (c Car) drive() {
// 	fmt.Println(c.CarName, "andou")
// }

// func main() {

// 	j := []byte(`{"car":"BMW", "year":2020}`)

// 	var car Car

// 	json.Unmarshal(j, &car)

// 	fmt.Println(car.CarName)

// }
