package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

/* How to run this demo:
A. Get the code and drop inside a docker container.
git clone git@github.com:go-delve/delve.git
cd delve
git checkout issues/2249-A
docker-compose run app

B. build delve and debug the example application
rm -rf example/example /go/bin/dlv && \
go mod tidy && \
make install && \
go build -x -gcflags="all=-N -l" -ldflags='all=-linkshared' -o example/example example/main.go && \
/go/bin/dlv exec example/example

C. set breakpoint and execute various `whatis` commands
break example/main.go:86
(dlv) continue
(dlv) whatis -v f
(dlv) whatis -v hReq

Note:
1. This is just a demo full of bad half-baked code.
2. This has only been tested on linux 64bit
3. For now, it depends on `-buildmode=shared`
4. There's a proposal by rsc to remove -buildmode=shared https://github.com/golang/go/issues/47788
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

	myChan := make(chan int64)
	bufmyChan := make(chan Distance, 2)
	bufmyChan <- d

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
	// whatis myChan
	// whatis bufmyChan
	fmt.Println(f)
	fmt.Println(hReq)
	fmt.Println("hey", cool, inter, interTwo, MyFn, myChan, bufmyChan)
}

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
