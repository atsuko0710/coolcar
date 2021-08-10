package main

import (
	"context"
	"fmt"
)

type ParamKey struct {}

func main() {
	c := context.WithValue(context.Background(), ParamKey{}, "abc")
	mainTask(c)
}

func mainTask(c context.Context) {
	fmt.Print(c.Value(ParamKey{}))
	smallTask(context.Background(), "task1")
	smallTask(c, "task2")
}

func smallTask(c context.Context, name string)  {
	fmt.Print(c.Value(ParamKey{}))
	fmt.Printf("%s started\n", name)
}