package main

import (
	"fmt"
	"testing"
	"time"
)

// Testing only by no exception for now
// * TODO: Add more complex checks
func Test_f1(t *testing.T) {
	fmt.Println(f(77))
}
func Test_f2(t *testing.T) {
	fmt.Println(f(77, 321))
}
func Test_f3(t *testing.T) {
	tmp, _ := time.Parse("2006-Jan-02", "2013-Feb-03")
	fmt.Println(f(77, 321, tmp))
}
func Test_f4(t *testing.T) {
	tmp, _ := time.Parse("2006-Jan-02", "2013-Feb-03")
	fmt.Println(f(77, 321, tmp, &point{10, 50, 255}))
}
