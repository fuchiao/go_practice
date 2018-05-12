# Issue 14813

https://github.com/golang/go/issues/14813

> popcnt is a leaf function, so the compiler can inline it.
> Because the function is inlined, the compiler can see it has no side effects, so the call is eliminated.

### when the function call is eliminated
```
$ go test -bench . ./issue14813
goos: windows
goarch: amd64
BenchmarkPopcnt-4       2000000000               0.42 ns/op
```
### put the return value into exported variable to prevent the issue
```
$ go test -bench . ./issue14813
goos: windows
goarch: amd64
BenchmarkPopcnt-4       1000000000               2.60 ns/op
```

### exported type should have comment
https://github.com/golang/go/wiki/CodeReviewComments