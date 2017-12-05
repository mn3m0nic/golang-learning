### Current version

v0.02

### Program output


Simple program output:

```
./main
77
77
10
321
2013-02-03 00:00:00 +0000 UTC
2017-12-05 13:55:52.448419561 +0300 +03
```


## Tests 

### Testing execution with different parameters

```
=== RUN   Test_f1
{77 10 {63648073262 257812446 0x5dd520} <nil>}
--- PASS: Test_f1 (0.00s)
=== RUN   Test_f2
{77 321 {63648073262 257863757 0x5dd520} <nil>}
--- PASS: Test_f2 (0.00s)
=== RUN   Test_f3
{77 321 {63495446400 0 <nil>} <nil>}
--- PASS: Test_f3 (0.00s)
=== RUN   Test_f4
{77 321 {63495446400 0 <nil>} 0xc42000aba0}
--- PASS: Test_f4 (0.00s)
```

#### Testing default and custom

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
PASS
ok  	_/task1	0.002s
```

### Coverage

```
{77 10 {63648073262 458836925 0x5df5c0} <nil>}
{77 321 {63648073262 458893908 0x5df5c0} <nil>}
{77 321 {63495446400 0 <nil>} <nil>}
{77 321 {63495446400 0 <nil>} 0xc42000ab60}
PASS
coverage: 63.2% of statements
ok  	_/task1	0.002s
```
