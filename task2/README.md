
## Input file

```
Tests passed: 64
Tests failed: 2 (1 new), passed: 173, ignored: 6, muted: 3; process exited with code 1
Tests failed: 4 (3 new), passed: 55; process exited with code 1
Tests failed: 5 (1 new), passed: 125; process exited with code 1
Tests failed: 1 (1 new), passed: 310, ignored: 13; process exited with code 1
Tests passed: 311, ignored: 13
Tests failed: 1 (1 new), passed: 0; process exited with code 1
Tests failed: 3 (2 new), passed: 174, ignored: 7, muted: 1; process exited with code 1
Tests passed: 2, failed: 173, muted: 6, ignored: 3; process exited with code 1
Tests failed: 15, passed: 151; process exited with code 1
Success
Process exited with code 1
Canceled
Snapshot dependency builds failed: 1
Canceled (Process exited with code 137)
```

### Output file

```
{0 64 0}
{2 173 9}
{4 55 0}
{5 125 0}
{1 310 13}
{0 311 13}
{1 0 0}
{3 174 8}
{173 2 9}
{15 151 0}
{0 0 0}
{0 0 0}
{0 0 0}
{0 0 0}
{0 0 0}
```


### Tests


```
go test -v
=== RUN   Test_parse1
--- PASS: Test_parse1 (0.00s)
=== RUN   Test_parse2
--- PASS: Test_parse2 (0.00s)
=== RUN   Test_parse3
--- PASS: Test_parse3 (0.00s)
PASS
ok  	_/home/mn3m/epam/bench/golang/tasks2-1	0.001s
```
