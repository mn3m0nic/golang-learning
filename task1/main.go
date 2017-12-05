package main
/* 1) Write a function that takes a one required parameter int and three optional: 
first - int, default value 10
second - time.Time, default time.Now()
third - pointer to any struct, default - nil
The function should return all passed parameters.
Write unit tests for this function.
*/
import ( 
 // "os"
  "fmt"
 // "strconv"
  "time"
)
type point struct {
  x int
  y int
  value int
}


func f(in1 int, args ...interface{}) (n1,n2 int, t time.Time, p *point)   {
  n1=in1
  n2=10
  t=time.Now()
  p=nil
  l:=len(args)
  if l>0 {
    n2=args[0].(int)
  }
  if l>1 {
    t=args[1].(time.Time)
  }
  if l>2 {
    p=args[2].(*point)
  }
   return
}

func main() {
  fmt.Println(f(77))
  fmt.Println(f(77,321))
  tmp, _ := time.Parse("2006-Jan-02", "2013-Feb-03")
  fmt.Println(f(77,321, tmp ))
  fmt.Println(f(77,321, tmp, &point{10,20,255}))
}
