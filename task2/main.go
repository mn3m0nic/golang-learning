package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type tests struct {
	failed  int64
	passed  int64
	ignored int64
}

func filtern(st, reg string) (number int64) {
	re1 := regexp.MustCompile(reg + ": [0-9]*")
	substring1 := re1.FindString(st)
	substring1number := regexp.MustCompile("[0-9]+")
	snumber := substring1number.FindString(substring1)
	number, _ = strconv.ParseInt(snumber, 10, 64)
	return
}
func parse(st string) (t tests) {
	t.failed = filtern(st, "(Tests|,) failed")
	t.passed = filtern(st, "passed")
	t.ignored = filtern(st, "ignored") + filtern(st, "muted")
	return
}
func main() {
	file, _ := os.Open("in.txt")
	defer file.Close()
	str := bufio.NewScanner(file)
	for str.Scan() {
		fmt.Println(parse(str.Text()))
	}
}
