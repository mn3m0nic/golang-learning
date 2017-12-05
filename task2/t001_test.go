package main
import (
  "testing"
)
func Test_parse1 (t *testing.T)  {
  var t1 tests
  t1 = parse("passed: 100")
  if t1.passed!=100 {
    t.Errorf("test 1 failed with %d, expected %d\n",t1.passed,100)
  }
}
func Test_parse2 (t *testing.T)  {
  var t1 tests
  t1 = parse("Tests failed: 321")
  if t1.failed!=321 {
    t.Errorf("test 2 failed with %d, expected %d\n",t1.failed,321)
  }
}
func Test_parse3 (t *testing.T)  {
  var t1 tests
  t1 = parse("muted: 10, ignored: 20")
  if t1.ignored!=30 {
    t.Errorf("test 3 failed with %d, expected %d\n",t1.ignored,30)
  }
}

