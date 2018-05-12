# StringConcat

about dotdotdot
https://golang.org/ref/spec#Passing_arguments_to_..._parameters

add this line to record the number of allocation
```
b.ReportAllocs()
```

check the result
```
$ go test -bench . ./StringConcat
goos: windows
goarch: amd64
BenchmarkStringConcat1-4         2000000              1074 ns/op             262 B/op          8 allocs/op
BenchmarkStringConcat2-4         1000000              1447 ns/op             445 B/op          9 allocs/op
BenchmarkStringConcat3-4         1000000              1174 ns/op             262 B/op          8 allocs/op
BenchmarkStringConcat4-4         3000000               499 ns/op             128 B/op          2 allocs/op
PASS
```