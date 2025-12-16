package workemailvalidator_test

import (
	"testing"

	workemailvalidator "github.com/qeeqez/work-email-validator"
)

func TestIsWorkEmail(t *testing.T) {
	t.Parallel()

	tests := []struct {
		email    string
		expected bool
	}{
		{"user@gmail.com", false},
		{"user@outlook.com", false},
		{"user@temp-mail.com", false},
		{"user@mycompany.com", true},
		{"contact@example.com", true}, // example.com is not in free/disposable lists
		{"invalid-email", false},
		{"user@sub.gmail.com", false},
		{"user@sub.temp-mail.com", false},
		{"user@corp.google.com", true}, // google.com is not in free list
	}

	for _, testCase := range tests {
		t.Run(testCase.email, func(t *testing.T) {
			t.Parallel()

			result := workemailvalidator.IsWorkEmail(testCase.email)
			if result != testCase.expected {
				t.Errorf("IsWorkEmail(%q) = %v, want %v", testCase.email, result, testCase.expected)
			}
		})
	}
}
