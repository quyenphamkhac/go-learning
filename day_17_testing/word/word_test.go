package word

import "testing"

func TestPalindrome(t *testing.T) {
	input1 := "detartrated"
	if !IsPalindrome(input1) {
		t.Errorf(`IsPalindrome("%q") = true`, input1)
	}
	input2 := "kayak"
	if !IsPalindrome(input2) {
		t.Errorf(`IsPalindrome("%q") = true`, input2)
	}
}

func TestNonPalindrome(t *testing.T) {
	input := "hello"
	if IsPalindrome(input) {
		t.Errorf(`IsPalindrome("%q") = false`, input)
	}
}

func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") = true`)
	}
}

func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = true`, input)
	}
}

func TestIsPalindrome(t *testing.T) {
	// table-driven testing
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false},
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v but go %v", test.input, test.want, got)
		}
	}
}
