package main

import (
	"fmt"
	"strconv"
	"time"
)

type Distance float64

type Health struct {
	ID   uint64
	Date time.Time
}

type Fire struct {
	HH   Health
	Age  int64
	Name string
	Dist Distance
	Day  time.Time
}

func (f Fire) Hello(i int) string {
	return strconv.Itoa(i)
}

func (f *Fire) MethodTwo() int64 {
	return 78
}

/*
go mod tidy
make install
/Users/komuw/go/bin/dlv debug example/main.go
break example/main.go:48
whatis f
*/
func main() {
	f := Fire{
		HH:   Health{ID: 67, Date: time.Now()},
		Age:  45,
		Name: "Komu",
		Dist: 4.8}
	f.Hello(89) // If it is not called; then delve is not able to find it.
	f.MethodTwo()

	fmt.Println(f)
	fmt.Println("hey")
}
