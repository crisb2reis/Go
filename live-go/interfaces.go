package main

import (
	"fmt"
)

type veiculo interface {
	start() string
}

type Car struct {
	CarName string `json: "car"`
	CarYear int    `json: "year"`
}

func (c Car) start() string {
	return "Startou!!!"
}

func (c Car) drive() {
	fmt.Println(c.CarName, "andou")
}

func exemplo(car veiculo) {

}

func main() {

	car := Car{
		CarName: "BMW",
		CarYear: 2020,
	}
	exemplo(car)

}
