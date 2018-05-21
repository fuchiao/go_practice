package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"testing"
)

// subtests
func TestFoo(t *testing.T) {
	// setUp
	fmt.Println("SetUp")
	a := 1
	t.Run("Add1", func(t *testing.T) {
		if a+1 != 2 {
			t.Fatal("Failed")
		}
	})
	t.Run("Add2", func(t *testing.T) {
		if a+2 != 3 {
			t.Fatal("Failed")
		}
	})
	t.Run("Add3", func(t *testing.T) {
		if a+3 != 4 {
			t.Fatal("Failed")
		}
	})
	fmt.Println("TearDown")
	// tearDown
}

// table driven tests
func TestAdd1(t *testing.T) {
	cases := []struct{ A, B, Expected int }{
		{1, 2, 3},
		{2, 3, 5},
		{3, 4, 7},
		{4, 5, 9},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d+%d", tc.A, tc.B), func(t *testing.T) {
			actual := tc.A + tc.B
			if actual != tc.Expected {
				t.Fatal("Failed")
			}
		})
	}
}

// naming cases
func TestAdd2(t *testing.T) {
	cases := []struct {
		Name           string
		A, B, Expected int
	}{
		{"one plus two", 1, 2, 3},
		{"two plus three", 2, 3, 5},
		{"three plus four", 3, 4, 7},
		{"four plus five", 4, 5, 9},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			actual := tc.A + tc.B
			if actual != tc.Expected {
				t.Fatal("Failed")
			}
		})
	}
}

// create cleanup function
func testTempFile(t *testing.T) (string, func()) {
	t.Helper() // hide the info of the function
	tf, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	t.Fatal("test")
	tf.Close()
	return tf.Name(), func() { os.Remove(tf.Name()) }
}
func TestSomeThing(t *testing.T) {
	_, tfclose := testTempFile(t)
	defer tfclose()
}

// not have to mock net.conn, create a real connection
func testConn(t *testing.T, msg []byte) net.Conn {
	t.Helper()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal("Listen() Failure")
	}
	go func() {
		defer ln.Close()
		server, err := ln.Accept()
		if err != nil {
			t.Fatal("Server Accept() Failure")
		}
		buf := make([]byte, 512)
		n, err := server.Read(buf)
		if err != nil {
			t.Fatal("Server Read() Failure")
		}
		fmt.Printf("Server received: %s\n", string(buf[:n]))
		if !bytes.Equal(msg, buf[:n]) {
			t.Fatal("Failed")
		}
	}()
	client, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatal("Client Dial() Failure")
	}
	return client
}

func TestConn(t *testing.T) {
	msg := []byte("HelloWorld")
	cli := testConn(t, msg)
	cli.Write(msg)
	fmt.Printf("Client send: %s\n", string(msg))
}
