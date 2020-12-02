package main

import "testing"

func TestIsValidPasswordPuzzle1WithValidPasswords(t *testing.T) {
	if !IsValidPasswordPuzzle1("abcde", "a", 1, 3) {
		t.Error("Fail: 1-3 a: abcde")
	}
	if !IsValidPasswordPuzzle1("ccccccccc", "c", 2, 9) {
		t.Error("Fail: 2-9 c: ccccccccc")
	}
}

func TestIsIsValidPasswordPuzzle1WithInvalidPassword(t *testing.T) {
	if IsValidPasswordPuzzle1("cdefg", "b", 1, 3) {
		t.Error("Fail: 1-3 b: cdefg")
	}
}

func TestIsValidPasswordPuzzle2WithValidPasswords(t *testing.T) {
	if !IsValidPasswordPuzzle2("abcde", "a"[0], 1, 3) {
		t.Error("Fail: 1-3 a: abcde")
	}
}

func TestIsValidPasswordPuzzle2WithInvalidPassword(t *testing.T) {
	if IsValidPasswordPuzzle2("cdefg", "b"[0], 1, 3) {
		t.Error("Fail: 1-3 b: cdefg")
	}

	if IsValidPasswordPuzzle2("ccccccccc", "c"[0], 2, 9) {
		t.Error("Fail: 2-9 c: ccccccccc")
	}
}
