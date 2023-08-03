package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tc := range testcases {
		rev, _ := Reverse(tc.in)
		if rev != tc.want {
			t.Errorf("Reverse: %v, want %v", rev, tc.want)
		}
	}
}

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, revError := Reverse(orig)
		if revError != nil {
			return
		}
		doubleRev, doubleRevError := Reverse(rev)
		if doubleRevError != nil {
			return
		}
		if orig != doubleRev {
			t.Errorf("Before: %v, after: %v", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produces invalid UTF-8 string %v", rev)
		}
	})
}
