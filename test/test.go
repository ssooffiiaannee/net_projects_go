package test

import (
	"net"
	// "fmt"
	"testing"
)

// func TestListener(){
// 	fmt.Println("hello test !")
// }

func TestListener(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = listener.Close() }()
	t.Logf("bound to %q", listener.Addr())
}