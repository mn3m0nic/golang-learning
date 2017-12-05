package main

import (
	//"fmt"
	"testing"
	"time"
)

// Testing n1 single value
func Test_n1_single_value(t *testing.T) {
	n1_res := f(77).n1
	if n1_res != 77 {
		t.Errorf("Test failed with %d, expected %d", n1_res, 77)
	}
}

// Testing default value N2 should be 10
func Test_n2_default_value(t *testing.T) {
	n2_res := f(77).n2
	if n2_res != 10 {
		t.Errorf("Test failed with %d, expected %d", n2_res, 10)
	}
}

// Testing custom n2 is set and returned
func Test_n2_custom_value(t *testing.T) {
	n2_res := f(77, 123).n2
	if n2_res != 123 {
		t.Errorf("Test failed with %d, expected %d", n2_res, 123)
	}
}

// Testing custom t default value is NOW
// (with acceptable shift 1000ms)
func Test_t_default_value_is_NOW(t *testing.T) {
	t_res := f(77).t
	t_tmp := time.Now()
	t_diff := t_tmp.Sub(t_res)
	if t_diff > 1000 {
		t.Errorf("Test failed with %d, expected %d", t_diff, 200)
	}
}

// Testing t (time.Time) custom value
func Test_t_custom_value(t *testing.T) {
	tmp, _ := time.Parse("2006-Jan-02", "2013-Feb-03")
	t_res := f(77, 123, tmp).t
	if t_res != tmp {
		t.Errorf("Test failed with %d, expected %d", t_res, tmp)
	}
}

// Testing default p (pointer) value is nil
func Test_p_default_value_is_nil(t *testing.T) {
	p_res := f(77).p
	if p_res != nil {
		t.Errorf("Test failed with %d, expected nil", p_res)
	}
}


func Test_different_order_1(t *testing.T) {
	p_res := f(11, &point{1,2,3}).n2
	if p_res != 10 {
		t.Errorf("Test failed with %d, expected %d", p_res, 10)
	}
}
