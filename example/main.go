package main

import (
	"fmt"
	"net/http"
	"os"
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

type Cooler struct{}

func (d Cooler) Ala() float64 {
	return 3.28084
}

type Fire struct {
	HH        Health
	Age       int64
	Name      string
	Dist      Distance
	Day       time.Time
	privField int64
}

func (f Fire) Hello(i int) string {
	return strconv.Itoa(i)
}

func (f *Fire) MethodTwo() int64 {
	return 78
}

func (f Fire) privMethod() string {
	return "A private Method"
}

type MyInter interface {
	Read(p []byte) (n int, err error)
}

type fn func(x int, y string) uint64

func (f fn) MethodOnFunc() int64 {
	return 3465
}

/*
go mod tidy
make install
go build -x -gcflags="all=-N -l" -ldflags='all=-linkshared' -o example/example example/main.go
/go/bin/dlv exec example/example
# dlv debug example/main.go # does not work. # we need to update the default debug command to include `-linkshared`

rm -rf example/example && \
go mod tidy && \
make install && \
go build -x -gcflags="all=-N -l" -ldflags='all=-linkshared' -o example/example example/main.go && \
/go/bin/dlv exec example/example

break example/main.go:124
whatis d
whatis f     // f is a value Type struct
whatis hReq  // hReq is a pointer Type struct
*/
func main() {
	d := Distance(4.8)
	f := Fire{ // f is a value Type struct
		HH:        Health{ID: 67, Date: time.Now()},
		Age:       45,
		Name:      "Komu",
		Dist:      d,
		privField: 4}

	// TODO: fix,
	// If it is not called; then delve is not able to find it.
	// eg: the `ToFeet` method of the type `Distance` is not found when you do `whatis d`
	// f.Hello(89)
	// f.MethodTwo()
	// f.privMethod()
	_ = d.ToCm()
	cool := &Cooler{}

	myFile, _ := os.Create("/tmp/cool.txt")
	var inter MyInter = myFile
	var interTwo MyInter = nil

	var MyFn fn = func(x int, y string) uint64 {
		return 56
	}

	hReq, err := http.NewRequest("GET", "https://google.com", nil) // hReq is a pointer Type struct
	if err != nil {
		panic(fmt.Sprintf("http.NewRequest err: %v", err))
	}

	// whatis d
	// whatis Distance
	// whatis f
	// whatis cool
	// whatis myFile
	// whatis inter
	// whatis MyInter
	// whatis interTwo
	// whatis MyFn
	// whatis fn
	// whatis err
	// whatis hReq
	fmt.Println(f)
	fmt.Println(hReq)
	fmt.Println("hey", cool, inter, interTwo, MyFn)
}
