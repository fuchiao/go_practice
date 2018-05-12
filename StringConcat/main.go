package main

import (
	"bytes"
	"fmt"
	"time"
)

func concat1(name string) string {
	s := "hello, "
	s += name
	s += ". "
	s += time.Now().String()
	return s
}

func concat2(name string) string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "hello, %s. %v", name, time.Now().String())
	return b.String()
}

func concat3(name string) string {
	return fmt.Sprintf("hello, %s. %v", name, time.Now().String())
}

func concat4(name string) string {
	b := make([]byte, 0, 40)
	b = append(b, "hello, "...)
	b = append(b, name...)
	b = append(b, ". "...)
	b = time.Now().AppendFormat(b, time.UnixDate)
	return string(b)
}

func main() {
	fmt.Println(concat4("zack"))
}
