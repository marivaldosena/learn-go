package strutils

import "testing"

const word = "test"

func TestContains(t *testing.T) {
	result := Contains(word, "es")

	if result != true {
		t.Error("Contains expected to be true, but got ", result)
	}
}

func TestCount(t *testing.T) {
	result := Count(word, "t")

	if result != 2 {
		t.Error("Count expected to be 2, but got ", result)
	}
}

func TestBeginsWith(t *testing.T) {
	result := BeginsWith(word, "te")

	if result != true {
		t.Error("BeginsWith expected to be true, but got ", result)
	}
}

func TestEndsWith(t *testing.T) {
	result := EndsWith(word, "st")

	if result != true {
		t.Error("EndsWith expected to be true, but got ", result)
	}
}

func TestIndex(t *testing.T) {
	result := Index(word, "e")

	if result != 1 {
		t.Error("Index expected to be true, but got ", result)
	}
}

func TestConcat(t *testing.T) {
	result := Concat([]string{"a", "b"}, "-")

	if result != "a-b" {
		t.Error("Concat expected to be `a-b`, but got ", result)
	}
}

func TestRepeat(t *testing.T) {
	result := Repeat("a", 5)

	if result != "aaaaa" {
		t.Error("Repeat expected to be `aaaaa`, but got ", result)
	}
}

func TestUpper(t *testing.T) {
	result := Upper(word)

	if result != "TEST" {
		t.Error("Upper expected to be `TEST`, but got ", result)
	}
}
