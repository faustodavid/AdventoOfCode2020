package main

import "testing"

func Test_GetRow_Valid_Boarding_Returns_44(t *testing.T) {
	/*
		Start by considering the whole range, rows 0 through 127.
		F means to take the lower half, keeping rows 0 through 63.
		B means to take the upper half, keeping rows 32 through 63.
		F means to take the lower half, keeping rows 32 through 47.
		B means to take the upper half, keeping rows 40 through 47.
		B keeps rows 44 through 47.
		F keeps rows 44 through 45.
		The final F keeps the lower of the two, row 44.
	*/

	boardingId := "FBFBBFFRLR"
	expectedRow := 44

	actualColumn := GetRow(boardingId)

	if actualColumn != expectedRow {
		t.Errorf("Expected column: %d Actual: %d", expectedRow, actualColumn)
	}
}

func Test_GetColumn_Valid_Boarding_Returns_5(t *testing.T) {
	/*
		Start by considering the whole range, columns 0 through 7.
		R means to take the upper half, keeping columns 4 through 7.
		L means to take the lower half, keeping columns 4 through 5.
		The final R keeps the upper of the two, column 5.
	*/

	boardingId := "FBFBBFFRLR"
	expectedColumn := 5

	actualColumn := GetColumn(boardingId)

	if actualColumn != expectedColumn {
		t.Errorf("Expected column: %d Actual: %d", expectedColumn, actualColumn)
	}
}

func Test_GetSeatId(t *testing.T) {
	/*
	   	BFFFBBFRRR: row 70, column 7, seat ID 567.
	      FFFBBBFRRR: row 14, column 7, seat ID 119.
	      BBFFBBFRLL: row 102, column 4, seat ID 820.
	*/

	boardingId1 := "BFFFBBFRRR"
	expectedBoardingPass1 := 567

	boardingId2 := "FFFBBBFRRR"
	expectedBoardingPass2 := 119

	boardingId3 := "BBFFBBFRLL"
	expectedBoardingPass3 := 820

	actualBoardingPass1 := GetSeatId(boardingId1)
	actualBoardingPass2 := GetSeatId(boardingId2)
	actualBoardingPass3 := GetSeatId(boardingId3)

	if actualBoardingPass1 != expectedBoardingPass1 {
		t.Errorf("Expected boarding pass: %d Actual boading pass: %d", expectedBoardingPass1, actualBoardingPass1)
	}

	if actualBoardingPass2 != expectedBoardingPass2 {
		t.Errorf("Expected boarding pass: %d Actual boading pass: %d", actualBoardingPass2, expectedBoardingPass2)
	}

	if actualBoardingPass3 != expectedBoardingPass3 {
		t.Errorf("Expected boarding pass: %d Actual boading pass: %d", actualBoardingPass3, expectedBoardingPass3)
	}
}
