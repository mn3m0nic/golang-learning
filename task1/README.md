### Current version

v0.02

### Program output

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

```
=== RUN   Test_n1_single_value
--- PASS: Test_n1_single_value (0.00s)
=== RUN   Test_n2_default_value
--- PASS: Test_n2_default_value (0.00s)
=== RUN   Test_n2_custom_value
--- PASS: Test_n2_custom_value (0.00s)
=== RUN   Test_t_default_value_is_NOW
--- PASS: Test_t_default_value_is_NOW (0.00s)
=== RUN   Test_f1
{77 10 {63648068206 122436360 0x5dd4e0} <nil>}
--- PASS: Test_f1 (0.00s)
=== RUN   Test_f2
{77 321 {63648068206 122544598 0x5dd4e0} <nil>}
--- PASS: Test_f2 (0.00s)
=== RUN   Test_f3
{77 321 {63495446400 0 <nil>} <nil>}
--- PASS: Test_f3 (0.00s)
=== RUN   Test_f4
{77 321 {63495446400 0 <nil>} 0xc42000abc0}
--- PASS: Test_f4 (0.00s)
PASS
ok  	_/home/mn3m/devel/private/golang-leaning/golang-learning/task1	0.005s
```

### Tests Coverage

```
{77 10 {63648068206 340646866 0x5dd580} <nil>}
{77 321 {63648068206 340670944 0x5dd580} <nil>}
{77 321 {63495446400 0 <nil>} <nil>}
{77 321 {63495446400 0 <nil>} 0xc42000ab80}
PASS
coverage: 63.2% of statements
ok  	_/home/mn3m/devel/private/golang-leaning/golang-learning/task1	0.001s
```
