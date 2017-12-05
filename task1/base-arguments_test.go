package main
import (
  //"fmt"
  "testing"
  "time"
)
// Testing n1 single value
func Test_n1_single_value (t *testing.T) {
  n1_res:=f(77).n1
  if n1_res!=77 {
    t.Errorf("Test failed with %d, expected %d",n1_res, 77)
  }
}
// Testing n1 is required and can't be skipped
//func Test_n1_required (t *testing.T) {
//  n1_res,err:=f().n1
//  if err==nil {
//    t.Errorf("First argument is not required which not match task")
//  }
//}


// Testing default value N2 should be 10
func Test_n2_default_value (t *testing.T) {
  n2_res:=f(77).n2
  if n2_res!=10 {
    t.Errorf("Test failed with %d, expected %d",n2_res, 10)
  }
}


// Testing custom n2 is set and returned
func Test_n2_custom_value (t *testing.T) {
  n2_res:=f(77,123).n2
  if n2_res!=123 {
    t.Errorf("Test failed with %d, expected %d",n2_res,123)
  }
}

// Testing custom t default value is NOW 
// (with acceptable shift 200ms)
func Test_t_default_value_is_NOW (t *testing.T) {
  t_res:=f(77).t
  t_tmp:=time.Now()
  t_diff:=t_tmp.Sub(t_res)
  if t_diff>200 {
    t.Errorf("Test failed with %d, expected %d",t_diff,200)
  }
}

