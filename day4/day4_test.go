package main

import "testing"

// byr (Birth Year) - four digits; at least 1920 and at most 2002.

func Test_isPropertyValid_byr_between_1920_and_2002_should_be_true(t *testing.T) {
	property1 := "byr:1920"
	property2 := "byr:2002"
	property3 := "byr:1925"
	property4 := "byr:2000"

	if !isPropertyValid(property1) {
		t.Errorf("Property: %s was expected to be true", property1)
	}
	if !isPropertyValid(property2) {
		t.Errorf("Property: %s was expected to be true", property2)
	}
	if !isPropertyValid(property3) {
		t.Errorf("Property: %s was expected to be true", property3)
	}
	if !isPropertyValid(property4) {
		t.Errorf("Property: %s was expected to be true", property4)
	}
}

func Test_isPropertyValid_byr_outside_1920_and_2002_should_be_false(t *testing.T) {
	property1 := "byr:1919"
	property2 := "byr:2003"
	property3 := "byr:1900"
	property4 := "byr:2020"

	if isPropertyValid(property1) {
		t.Errorf("Property: %s was expected to be false", property1)
	}
	if isPropertyValid(property2) {
		t.Errorf("Property: %s was expected to be false", property2)
	}
	if isPropertyValid(property3) {
		t.Errorf("Property: %s was expected to be false", property3)
	}
	if isPropertyValid(property4) {
		t.Errorf("Property: %s was expected to be false", property4)
	}
}

// iyr (Issue Year) - four digits; at least 2010 and at most 2020.

func Test_isPropertyValid_iyr_between_2010_and_2020_should_be_true(t *testing.T) {
	property1 := "iyr:2010"
	property2 := "iyr:2020"
	property3 := "iyr:2011"
	property4 := "iyr:2015"

	if !isPropertyValid(property1) {
		t.Errorf("Property: %s was expected to be true", property1)
	}
	if !isPropertyValid(property2) {
		t.Errorf("Property: %s was expected to be true", property2)
	}
	if !isPropertyValid(property3) {
		t.Errorf("Property: %s was expected to be true", property3)
	}
	if !isPropertyValid(property4) {
		t.Errorf("Property: %s was expected to be true", property4)
	}
}

func Test_isPropertyValid_iyr_outside_2010_and_2020_should_be_false(t *testing.T) {
	property1 := "iyr:2009"
	property2 := "iyr:2021"
	property3 := "iyr:2030"
	property4 := "iyr:1900"

	if isPropertyValid(property1) {
		t.Errorf("Property: %s was expected to be false", property1)
	}
	if isPropertyValid(property2) {
		t.Errorf("Property: %s was expected to be false", property2)
	}
	if isPropertyValid(property3) {
		t.Errorf("Property: %s was expected to be false", property3)
	}
	if isPropertyValid(property4) {
		t.Errorf("Property: %s was expected to be false", property4)
	}
}

// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.

func Test_isPropertyValid_eyr_between_2020_and_2030_should_be_true(t *testing.T) {
	property1 := "eyr:2020"
	property2 := "eyr:2030"
	property3 := "eyr:2021"
	property4 := "eyr:2025"

	if !isPropertyValid(property1) {
		t.Errorf("Property: %s was expected to be true", property1)
	}
	if !isPropertyValid(property2) {
		t.Errorf("Property: %s was expected to be true", property2)
	}
	if !isPropertyValid(property3) {
		t.Errorf("Property: %s was expected to be true", property3)
	}
	if !isPropertyValid(property4) {
		t.Errorf("Property: %s was expected to be true", property4)
	}
}

func Test_isPropertyValid_eyr_outside_2020_and_2030_should_be_false(t *testing.T) {
	property1 := "eyr:2019"
	property2 := "eyr:2031"
	property3 := "eyr:2000"
	property4 := "eyr:2045"

	if isPropertyValid(property1) {
		t.Errorf("Property: %s was expected to be false", property1)
	}
	if isPropertyValid(property2) {
		t.Errorf("Property: %s was expected to be false", property2)
	}
	if isPropertyValid(property3) {
		t.Errorf("Property: %s was expected to be false", property3)
	}
	if isPropertyValid(property4) {
		t.Errorf("Property: %s was expected to be false", property4)
	}
}

/*
	hgt (Height) - a number followed by either cm or in:
	If cm, the number must be at least 150 and at most 193.
	If in, the number must be at least 59 and at most 76.
*/

func Test_isPropertyValid_hgt_in_cm_between_150_and_193_should_be_true(t *testing.T) {
	property1 := "hgt:150cm"
	property2 := "hgt:193cm"
	property3 := "hgt:151cm"
	property4 := "hgt:192cm"

	if !isPropertyValid(property1) {
		t.Errorf("Property: %s was expected to be true", property1)
	}
	if !isPropertyValid(property2) {
		t.Errorf("Property: %s was expected to be true", property2)
	}
	if !isPropertyValid(property3) {
		t.Errorf("Property: %s was expected to be true", property3)
	}
	if !isPropertyValid(property4) {
		t.Errorf("Property: %s was expected to be true", property4)
	}
}

func Test_isPropertyValid_hgt_in_inches_between_59_and_76_should_be_true(t *testing.T) {
	property1 := "hgt:59in"
	property2 := "hgt:76in"
	property3 := "hgt:70in"
	property4 := "hgt:63in"

	if !isPropertyValid(property1) {
		t.Errorf("Property: %s was expected to be true", property1)
	}
	if !isPropertyValid(property2) {
		t.Errorf("Property: %s was expected to be true", property2)
	}
	if !isPropertyValid(property3) {
		t.Errorf("Property: %s was expected to be true", property3)
	}
	if !isPropertyValid(property4) {
		t.Errorf("Property: %s was expected to be true", property4)
	}
}

func Test_isPropertyValid_hgt_in_cm_outside_150_and_193_should_be_false(t *testing.T) {
	property1 := "hgt:149cm"
	property2 := "hgt:194cm"
	property3 := "hgt:100cm"
	property4 := "hgt:200cm"

	if isPropertyValid(property1) {
		t.Errorf("Property: %s was expected to be false", property1)
	}
	if isPropertyValid(property2) {
		t.Errorf("Property: %s was expected to be false", property2)
	}
	if isPropertyValid(property3) {
		t.Errorf("Property: %s was expected to be false", property3)
	}
	if isPropertyValid(property4) {
		t.Errorf("Property: %s was expected to be false", property4)
	}
}

func Test_isPropertyValid_hgt_in_inches_outside_59_and_76_should_be_false(t *testing.T) {
	property1 := "hgt:58in"
	property2 := "hgt:77in"
	property3 := "hgt:23in"
	property4 := "hgt:99in"

	if isPropertyValid(property1) {
		t.Errorf("Property: %s was expected to be false", property1)
	}
	if isPropertyValid(property2) {
		t.Errorf("Property: %s was expected to be false", property2)
	}
	if isPropertyValid(property3) {
		t.Errorf("Property: %s was expected to be false", property3)
	}
	if isPropertyValid(property4) {
		t.Errorf("Property: %s was expected to be false", property4)
	}
}

// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.

func Test_isPropertyValid_hcl_should_start_with_sharp_symbol_and_have_6_chars_0_9_or_a_f(t *testing.T) {
	property1 := "hcl:#123456"
	property2 := "hcl:#abcdef"
	property3 := "hcl:#123abc"
	property4 := "hcl:#fed321"

	if !isPropertyValid(property1) {
		t.Errorf("Property: %s was expected to be true", property1)
	}
	if !isPropertyValid(property2) {
		t.Errorf("Property: %s was expected to be true", property2)
	}
	if !isPropertyValid(property3) {
		t.Errorf("Property: %s was expected to be true", property3)
	}
	if !isPropertyValid(property4) {
		t.Errorf("Property: %s was expected to be true", property4)
	}
}

func Test_isPropertyValid_hcl_without_starting_with_sharp_symbol_should_be_false(t *testing.T) {
	property1 := "hgt:58ig"
	property2 := "hgt:77ix"
	property3 := "hgt:23iy"
	property4 := "hgt:99ip"

	if isPropertyValid(property1) {
		t.Errorf("Property: %s was expected to be false", property1)
	}
	if isPropertyValid(property2) {
		t.Errorf("Property: %s was expected to be false", property2)
	}
	if isPropertyValid(property3) {
		t.Errorf("Property: %s was expected to be false", property3)
	}
	if isPropertyValid(property4) {
		t.Errorf("Property: %s was expected to be false", property4)
	}
}

func Test_isPropertyValid_hcl_with_more_or_less_than_6_chars_after_sharp_symbol_should_be_false(t *testing.T) {
	property1 := "hcl:#1234561"
	property2 := "hcl:#abcde"
	property3 := "hcl:#123"
	property4 := "hcl:#fed32134"

	if isPropertyValid(property1) {
		t.Errorf("Property: %s was expected to be false", property1)
	}
	if isPropertyValid(property2) {
		t.Errorf("Property: %s was expected to be false", property2)
	}
	if isPropertyValid(property3) {
		t.Errorf("Property: %s was expected to be false", property3)
	}
	if isPropertyValid(property4) {
		t.Errorf("Property: %s was expected to be false", property4)
	}
}

func Test_isPropertyValid_hcl_with_letter_higher_than_f_should_be_false(t *testing.T) {
	property1 := "hcl:#12345p"
	property2 := "hcl:#abcdeq"
	property3 := "hcl:#123abt"
	property4 := "hcl:#fed32r"

	if isPropertyValid(property1) {
		t.Errorf("Property: %s was expected to be false", property1)
	}
	if isPropertyValid(property2) {
		t.Errorf("Property: %s was expected to be false", property2)
	}
	if isPropertyValid(property3) {
		t.Errorf("Property: %s was expected to be false", property3)
	}
	if isPropertyValid(property4) {
		t.Errorf("Property: %s was expected to be false", property4)
	}
}

// pid (Passport ID) - a nine-digit number, including leading zeroes.

func Test_isPropertyValid_pid_should_be_9_digit_including_zeros(t *testing.T) {
	property1 := "pid:000000000"
	property2 := "pid:012345678"
	property3 := "pid:123456789"
	property4 := "pid:000000543"

	if !isPropertyValid(property1) {
		t.Errorf("Property: %s was expected to be true", property1)
	}
	if !isPropertyValid(property2) {
		t.Errorf("Property: %s was expected to be true", property2)
	}
	if !isPropertyValid(property3) {
		t.Errorf("Property: %s was expected to be true", property3)
	}
	if !isPropertyValid(property4) {
		t.Errorf("Property: %s was expected to be true", property4)
	}
}

func Test_isPropertyValid_pid_with_more_than_9_digits_should_be_false(t *testing.T) {
	property1 := "pid:0000000000"
	property2 := "pid:0123456782"
	property3 := "pid:0123456789"
	property4 := "pid:0000005473"

	if isPropertyValid(property1) {
		t.Errorf("Property: %s was expected to be false", property1)
	}
	if isPropertyValid(property2) {
		t.Errorf("Property: %s was expected to be false", property2)
	}
	if isPropertyValid(property3) {
		t.Errorf("Property: %s was expected to be false", property3)
	}
	if isPropertyValid(property4) {
		t.Errorf("Property: %s was expected to be false", property4)
	}
}

func Test_isPropertyValid_pid_with_letters_should_be_false(t *testing.T) {
	property1 := "pid:00000000a"
	property2 := "pid:01234567b"
	property3 := "pid:01234567r"
	property4 := "pid:00000054b"

	if isPropertyValid(property1) {
		t.Errorf("Property: %s was expected to be false", property1)
	}
	if isPropertyValid(property2) {
		t.Errorf("Property: %s was expected to be false", property2)
	}
	if isPropertyValid(property3) {
		t.Errorf("Property: %s was expected to be false", property3)
	}
	if isPropertyValid(property4) {
		t.Errorf("Property: %s was expected to be false", property4)
	}
}

func Test_countValidPassports_Puzzle2_invalid_passports(t *testing.T) {
	inputPassports := "eyr:1972 cid:100\nhcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926\n\niyr:2019\nhcl:#602927 eyr:1967 hgt:170cm\necl:grn pid:012533040 byr:1946\n\nhcl:dab227 iyr:2012\necl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277\n\nhgt:59cm ecl:zzz\neyr:2038 hcl:74454a iyr:2023\npid:3556412378 byr:2007"
	expectedCount := 0

	_, actualCount := countValidPassports(inputPassports)

	if actualCount != expectedCount {
		t.Errorf("Actual count was %d and was expected %d", actualCount, expectedCount)
	}
}

func Test_countValidPassports_Puzzle2_valid_passports_should_be_4(t *testing.T) {
	inputPassports := "pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980\nhcl:#623a2f\n\neyr:2029 ecl:blu cid:129 byr:1989\niyr:2014 pid:896056539 hcl:#a97842 hgt:165cm\n\nhcl:#888785\nhgt:164cm byr:2001 iyr:2015 cid:88\npid:545766238 ecl:hzl\neyr:2022\n\niyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719"
	expectedCount := 4

	_, actualCount := countValidPassports(inputPassports)

	if actualCount != expectedCount {
		t.Errorf("Actual count was %d and was expected %d", actualCount, expectedCount)
	}
}
