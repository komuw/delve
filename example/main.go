package main

import (
	"fmt"
	"strconv"
	"time"
)

// Distance is length in meters
type Distance float64

func (d Distance) ToCm() float64 {
	return float64(d) * 100
}
func (d *Distance) ToFeet() float64 {
	return float64(*d) * 3.28084
}

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
break example/main.go:63
whatis d
whatis f
*/
func main() {
	d := Distance(4.8)
	f := Fire{
		HH:   Health{ID: 67, Date: time.Now()},
		Age:  45,
		Name: "Komu",
		Dist: d}

	// TODO: fix,
	// If it is not called; then delve is not able to find it.
	// eg: the `ToFeet` method of the type `Distance` is not found when you do `whatis d`
	f.Hello(89)
	f.MethodTwo()
	_ = d.ToCm()

	fmt.Println(f)
	fmt.Println("hey")
}
