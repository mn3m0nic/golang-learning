### Current version

v0.03

### Changelog

- [X] Add support not dirrect order, for example: f(10, &point{…}),
- [X] Add test for 1 case of not dirrect order

### Program output

Simple program output:

```
./main
77
77
10
321
2013-02-03 00:00:00 +0000 UTC
2017-12-05 19:58:28.639815712 +0300 +03
{77 321 {63495446400 0 <nil>} 0xc42000a340}
2013-02-03 00:00:00 +0000 UTC
{10 10 {63648089908 639972070 0x50ed80} 0xc42000a3e0}
```

With function calls:

```
(...)
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
        // --- added: ---
	fmt.Println(f(10, &point{1, 2, 3}))
	// Output:
	// {10 10 {63648089695 53667753 0x50ed80} 0xc42000a480}
```


# Tests

### Execution


```
=== RUN   Test_f1
{77 10 {63648089913 150554742 0x5dd540} <nil>}
--- PASS: Test_f1 (0.00s)
=== RUN   Test_f2
{77 321 {63648089913 150595389 0x5dd540} <nil>}
--- PASS: Test_f2 (0.00s)
=== RUN   Test_f3
{77 321 {63495446400 0 <nil>} <nil>}
--- PASS: Test_f3 (0.00s)
=== RUN   Test_f4
{77 321 {63495446400 0 <nil>} 0xc42000aba0}
--- PASS: Test_f4 (0.00s)
```

### Test function calls with different arguments


```
=== RUN   Test_n1_single_value
--- PASS: Test_n1_single_value (0.00s)
=== RUN   Test_n2_default_value
--- PASS: Test_n2_default_value (0.00s)
=== RUN   Test_n2_custom_value
--- PASS: Test_n2_custom_value (0.00s)
=== RUN   Test_t_default_value_is_NOW
--- PASS: Test_t_default_value_is_NOW (0.00s)
=== RUN   Test_t_custom_value
--- PASS: Test_t_custom_value (0.00s)
=== RUN   Test_p_default_value_is_nil
--- PASS: Test_p_default_value_is_nil (0.00s)
=== RUN   Test_different_order_1
--- PASS: Test_different_order_1 (0.00s)
PASS
ok  	_/task1	0.002s
```


```
PASS
coverage: 47.6% of statements
ok  	_/task1	0.001s
```

