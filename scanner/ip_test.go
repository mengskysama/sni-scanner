package scanner

import (
	"fmt"
	"testing"
)

func Test_inetA2n(t *testing.T) {
	n, err := inetA2n("192.168.1.1")
	fmt.Println(n)
	if err != nil {
		t.Fatal(err)
	}
	if n != 3232235777 {
		t.Fatal("err val")
	}
}

func Test_inetN2a(t *testing.T) {
	ip := inetN2a(3232235777)
	if ip.String() != "192.168.1.1" {
		t.Fatal("err val")
	}
}
