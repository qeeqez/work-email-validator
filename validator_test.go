package workemailvalidator_test

import (
	"testing"

	workemailvalidator "github.com/qeeqez/work-email-validator"
)

func TestIsDisposableDomain(t *testing.T) {
	t.Parallel()

	tests := []struct {
		domain   string
		expected bool
	}{
		{"temp-mail.com", true},
		{"10minutemail.com", true},
		{"guerrillamail.com", true},
		{"gmail.com", false},
		{"example.com", false},
		{"TEMP-MAIL.COM", true},
		{"  temp-mail.org  ", true},
	}

	for _, testCase := range tests {
		t.Run(testCase.domain, func(t *testing.T) {
			t.Parallel()

			result := workemailvalidator.IsDisposableDomain(testCase.domain)
			if result != testCase.expected {
				t.Errorf("IsDisposableDomain(%q) = %v, want %v", testCase.domain, result, testCase.expected)
			}
		})
	}
}

func TestIsFreeDomain(t *testing.T) {
	t.Parallel()

	tests := []struct {
		domain   string
		expected bool
	}{
		{"gmail.com", true},
		{"outlook.com", true},
		{"yahoo.com", true},
		{"hotmail.com", true},
		{"icloud.com", true},
		{"protonmail.com", true},
		{"example.com", false},
		{"temp-mail.com", false},
		{"GMAIL.COM", true},
		{"  outlook.com  ", true},
	}

	for _, testCase := range tests {
		t.Run(testCase.domain, func(t *testing.T) {
			t.Parallel()

			result := workemailvalidator.IsFreeDomain(testCase.domain)
			if result != testCase.expected {
				t.Errorf("IsFreeDomain(%q) = %v, want %v", testCase.domain, result, testCase.expected)
			}
		})
	}
}

func TestIsDisposableOrFreeDomain(t *testing.T) {
	t.Parallel()

	tests := []struct {
		domain   string
		expected bool
	}{
		{"gmail.com", true},
		{"temp-mail.com", true},
		{"example.com", false},
		{"mycompany.com", false},
	}

	for _, testCase := range tests {
		t.Run(testCase.domain, func(t *testing.T) {
			t.Parallel()

			result := workemailvalidator.IsDisposableOrFreeDomain(testCase.domain)
			if result != testCase.expected {
				t.Errorf("IsDisposableOrFreeDomain(%q) = %v, want %v", testCase.domain, result, testCase.expected)
			}
		})
	}
}

func TestIsBusinessDomain(t *testing.T) {
	t.Parallel()

	tests := []struct {
		domain   string
		expected bool
	}{
		{"example.com", true},
		{"mycompany.com", true},
		{"gmail.com", false},
		{"temp-mail.com", false},
		{"outlook.com", false},
	}

	for _, testCase := range tests {
		t.Run(testCase.domain, func(t *testing.T) {
			t.Parallel()

			result := workemailvalidator.IsBusinessDomain(testCase.domain)
			if result != testCase.expected {
				t.Errorf("IsBusinessDomain(%q) = %v, want %v", testCase.domain, result, testCase.expected)
			}
		})
	}
}
