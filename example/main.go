package main

import (
	"fmt"
	"strconv"
)

type Distance float64

type Fire struct {
	Age  int64
	Name string
	Dist Distance
}

func (f Fire) Hello(i int) string {
	return strconv.Itoa(i)
}

/*
go mod tidy
make install
/Users/komuw/go/bin/dlv debug example/main.go
break example/main.go:29
whatis f
*/
func main() {
	f := Fire{Age: 45, Name: "Komu", Dist: 4.8}
	fmt.Println(f)
	fmt.Println("hey")
}
