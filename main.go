package main

import (
	"awesomeProject/stream"
	"context"
	"fmt"
	"log/slog"
	"strconv"
)

var globalVariable string = "global"
var ctx = context.Background()

func main() {

	x := 10
	y := 4
	z := 3

	expr := x%y == z-1 && x/y > z-2
	fmt.Println(expr)

	array := []int{0, 2}
	for i, v := range array {
		fmt.Printf("%v %v", i, v)
	}

	mp := map[string]int{"a": 1}
	for k, v := range mp {
		fmt.Printf("%v %v\n", k, v)
	}
	/* 	var p1 = Product{
	   		Id:   uuid.New(),
	   		Name: "new product",
	   		qtd:  0,
	   	}
	   	p1.AddProductSize()
	   	fmt.Println(p1) */

	callTypes()
	fmt.Println(globalVariable)

	staticTypes("hello")
	//scanAnOutPut("Rodrigo", 33, 5.2)

	streamTests()
}

func streamTests() {
	//list := Stream[int]{[]int{1, 2, 3}}.ToList()
	//fmt.Println(list)

	doubleIt := func(i int) int { return i * 2 }
	floatIt := func(i int) float64 { return float64(i * 5) }
	suffixIt := func(i float64) string {
		return strconv.FormatFloat(i, 'f', -1, 32) + " transformed"
	}
	filterIt := func(i int) bool { return i%2 == 0 }

	fmt.Println("iniciando print de stream")
	result := stream.From([]int{1, 2, 3, 4}).
		Filter(filterIt).
		Map(doubleIt).
		Map(floatIt).
		Peek(func(f float64) { fmt.Printf("parei pra olhar%.2f\n", f) }).
		Map(suffixIt).
		ToList()
	fmt.Println(result)
	fmt.Println("higher ", stream.From([]int{8, 1, 2, 3}).
		Reduce(0, stream.Max[int]))

	fmt.Println("max ", stream.From([]float64{8, 1, 2, 3}).
		Reduce(3.0, stream.Max[float64]))

	fmt.Println("add ", stream.From([]int{8, 1, 2, 3}).
		Reduce(0, stream.Add[int]))
}

type Person struct {
	Name  string
	Age   int
	Money float64
}

func scanAnOutPut(name string, age int, money float64) {
	fmt.Println("you want to change the person's age?")
	count, err := fmt.Scanf("%d", &age)
	if err != nil {
		slog.Log(ctx, slog.LevelError, "count of arguments %d", count)
		fmt.Println(err.Error())
		return
	}
	p := Person{name, age, money}
	fmt.Println(p)
}

func staticTypes(a string) {
	fmt.Println(a)
}

func callTypes() {
	var greeting string = "hello with builtin println"
	println(greeting)
	array := []int{1, 2, 3}
	fmt.Println(array)
	mapp := map[int]string{1: "a"}
	mapp[2] = "2"
	fmt.Println(mapp)
}
