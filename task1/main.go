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
	for _, arg := range args {
		switch v := arg.(type) {
		case int:
			o.n2 = arg.(int)
		case time.Time:
			o.t = arg.(time.Time)
		case *point:
			o.p = arg.(*point)
		default:
			fmt.Printf("unexpected type %T", v)
		}
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
	fmt.Println(f(77, 321, tmp, &point{10, 20, 255}))
	// Output:
	// {77 321 {63495446400 0 <nil>} 0xc42000a3e0}
	// ---- Part 2 ----
	fmt.Println(f(77, tmp).t)
	// Output:
	// 2013-02-03 00:00:00 +0000 UTC
	fmt.Println(f(10, &point{1, 2, 3}))
	// Output:
	// {10 10 {63648089695 53667753 0x50ed80} 0xc42000a480}
}
