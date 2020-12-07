package main

import "testing"

func Test_GroupAnswersByGroups_happy_flow(t *testing.T) {
	answers := "abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb"

	groupedAnswers := GroupAnswersByGroups(answers)

	if groupedAnswers[0] != "abc" {
		t.Errorf("Expected 'abc' actual '%s'", groupedAnswers[0])
	}

	if groupedAnswers[1] != "abc" {
		t.Errorf("Expected 'abc' actual '%s'", groupedAnswers[1])
	}

	if groupedAnswers[2] != "abac" {
		t.Errorf("Expected 'abac' actual '%s'", groupedAnswers[2])
	}
	if groupedAnswers[3] != "aaaa" {
		t.Errorf("Expected 'aaaa' actual '%s'", groupedAnswers[3])
	}

	if groupedAnswers[4] != "b" {
		t.Errorf("Expected 'b' actual '%s'", groupedAnswers[4])
	}
}

func Test_CountUniqueLetters_happy_flow(t *testing.T) {
	answers := "abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb"
	expectedCount := 11

	actualCount := CountUniqueLetters(GroupAnswersByGroups(answers))

	if actualCount != expectedCount {
		t.Errorf("Expected '%d' actual '%d'", expectedCount, actualCount)
	}
}
