package main

import "testing"

func TestIsPasswordValidWithValidPasswords(t *testing.T) {
	if !IsPasswordValid(CorruptedPassword{"abcde", "a", 1, 3}) {
		t.Error("Fail: 1-3 a: abcde")
	}
	if !IsPasswordValid(CorruptedPassword{"ccccccccc", "c", 2, 9}) {
		t.Error("Fail: 2-9 c: ccccccccc")
	}
}

func TestIsPasswordValidWithInvalidPassword(t *testing.T) {
	if IsPasswordValid(CorruptedPassword{"cdefg", "b", 1, 3}) {
		t.Error("Fail: 1-3 b: cdefg")
	}
}
