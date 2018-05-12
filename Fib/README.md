# Fib

run benchmarks
```
$ go test -bench . ./Fib
```
to get more accuracy
```
$ go test -bench . ./Fib -benchtime=10s
$ go test -bench . ./Fib -count=10
```
compare the results
```
# install benchstat (make sure $GOPATH/bin is included in $PATH)
$ go get -u rsc.io/benchstat

# comment case 2 in fib()
$ go test -bench . ./Fib -count=10|tee old.txt
# uncomment case 2 in fib()
$ go test -bench . ./Fib -count=10|tee new.txt
$ benchstat old.txt new.txt
name   old time/op  new time/op  delta
Fib20  71.2µs ± 1%  49.1µs ± 1%  -31.01%  (p=0.000 n=9+9)
```

if we want to avoid setup costs, we can try
* b.ResetTimer()
* b.StopTimer()
* b.StartTimer()

if we want to record the number of allocation
*  b.ReportAllocs()