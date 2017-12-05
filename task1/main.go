package main

import (
	"fmt"
	"time"
)

type point struct {
	x     int
	y     int
	value int
}
type fout struct {
	n1, n2 int
	t      time.Time
	p      *point
}

func f(in1 int, args ...interface{}) (o fout) {
	// setting default values:
	o.n1 = in1
	o.n2 = 10
	o.t = time.Now()
	o.p = nil
	// parsing arguments
	l := len(args)
	if l > 0 {
		o.n2 = args[0].(int)
	}
	if l > 1 {
		o.t = args[1].(time.Time)
	}
	if l > 2 {
		o.p = args[2].(*point)
	}
	return
}
func main() {
	tmp, _ := time.Parse("2006-Jan-02", "2013-Feb-03")
	//fmt.Println(f())
	// Not allowed: not enough arguments in call to f
	fmt.Println(f(77).n1)
	// Output: 77
	fmt.Println(f(77, 321, tmp, &point{10, 20, 255}).n1)
	// Output: 77
	fmt.Println(f(77).n2)
	// Output: 10 -- default value
	fmt.Println(f(77, 321, tmp, &point{10, 20, 255}).n2)
	// Output: 321
	fmt.Println(f(77, 321, tmp, &point{10, 20, 255}).t)
	// Output: 2013-02-03 00:00:00 +0000 UTC
	fmt.Println(f(77, 321).t)
	// Output: now
}
