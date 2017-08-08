package chop

import (
	"testing"
	"fmt"
)

func testChop(t *testing.T, chop Chopper, test, expected string) {
	value := chop.Chop(test)
	if value != expected {
		t.Fatal(expected, "!=", value)
	}
}

func TestChopper_Chop(t *testing.T) {
	testChop(t, Chopper{Slices: 3, Sep: '/', SliceSize: 3}, "abcdefxyz123", "abc/def/xyz123")
	testChop(t, Chopper{Slices: 2, Sep: '/', SliceSize: 3}, "abcdefxyz123", "abc/defxyz123")
	testChop(t, Chopper{Slices: 1, Sep: '/', SliceSize: 3}, "abcdefxyz123", "abcdefxyz123")

	testChop(t, Chopper{Slices: 3, Sep: '/', SliceSize: 2}, "abcdefxyz123", "ab/cd/efxyz123")

	// overflow
	testChop(t, Chopper{Slices: 100, Sep: '/', SliceSize: 3}, "abcdefxyz123", "abc/def/xyz/123")
	testChop(t, Chopper{Slices: 3, Sep: '/', SliceSize: 10000}, "abcdefxyz123", "abcdefxyz123")
	testChop(t, Chopper{Slices: 1000, Sep: '/', SliceSize: 10}, "abcdefxyz123", "abcdefxyz1/23")
}

func ExampleChopper_Chop() {
	ch := Chopper{Slices: 3, Sep: '/', SliceSize: 3}
	fmt.Println(ch.Chop("abcdefxyz123"))
	// Output: abc/def/xyz123
}
