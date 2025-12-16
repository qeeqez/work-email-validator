package workemailvalidator

import "testing"

func TestIsWorkEmail(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.email, func(t *testing.T) {
			result := IsWorkEmail(tt.email)
			if result != tt.expected {
				t.Errorf("IsWorkEmail(%q) = %v, want %v", tt.email, result, tt.expected)
			}
		})
	}
}
