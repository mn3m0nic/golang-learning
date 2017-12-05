package main
import (
  "testing"
)
// Testing parse function with correct value
func Test_parse_passed_value_with_number (t *testing.T)  {
  var t1 tests
  t1 = parse("passed: 100")
  if t1.passed!=100 {
    t.Errorf("test failed with %d, expected %d\n",t1.passed,100)
  }
}
// Testing parse function with incorrect value with text data
func Test_parse_passed_with_incorrect_data (t *testing.T)  {
  var t1 tests
  t1 = parse("passed: ZZZ")
  if t1.passed!=0 {
    t.Errorf("test failed with %d, expected %d\n",t1.passed,0)
  }
}
// Testing Parse function with correct value for "failed" regex
func Test_parse_test_failed_with_number (t *testing.T)  {
  var t1 tests
  t1 = parse("Tests failed: 321")
  if t1.failed!=321 {
    t.Errorf("test failed with %d, expected %d\n",t1.failed,321)
  }
}
// Testing Parse function with incorrect value for "failed" regex
func Test_parse_test_failed_with_incorrect_data (t *testing.T)  {
  var t1 tests
  t1 = parse("Tests failed: ZZZ")
  if t1.failed!=0 {
    t.Errorf("test failed with %d, expected %d\n",t1.failed,0)
  }
}
// Testing Parse function with correct value for "failed" regex
func Test_parse_test_failed_which_should_be_ignored (t *testing.T)  {
  var t1 tests
  t1 = parse("Snapshot dependency builds failed: 1")
  if t1.failed!=0 {
    t.Errorf("test failed with %d, expected %d\n",t1.failed,0)
  }
}

// Test Parse function that Muted and Ignored counted in ignored
func Test_parse_Muted_are_with_Ignored (t *testing.T)  {
  var t1 tests
  const expected=30 
  t1 = parse("muted: 10, ignored: 20")
  if t1.ignored!=expected {
    t.Errorf("test 3 failed with %d, expected %d\n",t1.ignored,expected)
  }
}
