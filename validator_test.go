package workemailvalidator

import "testing"

func TestIsDisposableDomain(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.domain, func(t *testing.T) {
			result := IsDisposableDomain(tt.domain)
			if result != tt.expected {
				t.Errorf("IsDisposableDomain(%q) = %v, want %v", tt.domain, result, tt.expected)
			}
		})
	}
}

func TestIsFreeDomain(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.domain, func(t *testing.T) {
			result := IsFreeDomain(tt.domain)
			if result != tt.expected {
				t.Errorf("IsFreeDomain(%q) = %v, want %v", tt.domain, result, tt.expected)
			}
		})
	}
}

func TestIsDisposableOrFreeDomain(t *testing.T) {
	tests := []struct {
		domain   string
		expected bool
	}{
		{"gmail.com", true},
		{"temp-mail.com", true},
		{"example.com", false},
		{"mycompany.com", false},
	}

	for _, tt := range tests {
		t.Run(tt.domain, func(t *testing.T) {
			result := IsDisposableOrFreeDomain(tt.domain)
			if result != tt.expected {
				t.Errorf("IsDisposableOrFreeDomain(%q) = %v, want %v", tt.domain, result, tt.expected)
			}
		})
	}
}

func TestIsBusinessDomain(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.domain, func(t *testing.T) {
			result := IsBusinessDomain(tt.domain)
			if result != tt.expected {
				t.Errorf("IsBusinessDomain(%q) = %v, want %v", tt.domain, result, tt.expected)
			}
		})
	}
}

func TestNormalizeDomain(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"GMAIL.COM", "gmail.com"},
		{"  outlook.com  ", "outlook.com"},
		{"Yahoo.Com", "yahoo.com"},
		{"example.com", "example.com"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := normalizeDomain(tt.input)
			if result != tt.expected {
				t.Errorf("normalizeDomain(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func BenchmarkIsDisposableDomain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsDisposableDomain("temp-mail.com")
	}
}

func BenchmarkIsFreeDomain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsFreeDomain("gmail.com")
	}
}

func BenchmarkIsBusinessDomain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsBusinessDomain("example.com")
	}
}
