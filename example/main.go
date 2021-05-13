package main

import (
	"fmt"
	"strconv"
)

type Fire struct {
	Age  int64
	Name string
}

func (f Fire) Hello(i int) string {
	return strconv.Itoa(i)
}

/*
go mod tidy
make install
/Users/komuw/go/bin/dlv debug example/main.go
break example/main.go:26
whatis f
*/
func main() {
	f := Fire{Age: 45, Name: "Komu"}
	fmt.Println(f)
	fmt.Println("hey")
}
