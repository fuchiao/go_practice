# Advanced Testing with Go
https://www.youtube.com/watch?v=8hQG7QlcLBk

## Subtests
* TestFoo
```
go test -run 'Foo' ./Subtests -v
go test -run 'Foo/Add2' ./Subtests -v
go test -run '/Add3' ./Subtests -v
```

## table driven tests

## create cleanup function

* Test Helper
https://golang.org/pkg/testing/#T.Helper

## net conn test

