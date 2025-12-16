package workemailvalidator_test

import (
	"strings"
	"testing"

	workemailvalidator "github.com/qeeqez/work-email-validator"
)

// FuzzIsDisposableDomain tests IsDisposableDomain with random inputs.
func FuzzIsDisposableDomain(f *testing.F) {
	// Seed with interesting test cases
	seeds := []string{
		"temp-mail.com",
		"TEMP-MAIL.COM",
		"  temp-mail.com  ",
		"sub.temp-mail.com",
		"gmail.com",
		"example.com",
		"",
		".",
		"..",
		"...",
		"a",
		"a.b",
		"a.b.c",
		strings.Repeat("a", 1000),
		"domain.com.",
		".domain.com",
		"domain..com",
		"@",
		"domain@com",
		"\x00domain.com",
		"domain.com\x00",
		"münchen.de",
		"😀.com",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, domain string) {
		// Should never panic
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("IsDisposableDomain panicked with input %q: %v", domain, r)
			}
		}()

		result := workemailvalidator.IsDisposableDomain(domain)

		// Result should be deterministic
		result2 := workemailvalidator.IsDisposableDomain(domain)
		if result != result2 {
			t.Errorf("IsDisposableDomain not deterministic for %q: %v != %v", domain, result, result2)
		}

		// Case insensitive - upper and lower should give same result
		if domain != "" {
			upperResult := workemailvalidator.IsDisposableDomain(strings.ToUpper(domain))

			lowerResult := workemailvalidator.IsDisposableDomain(strings.ToLower(domain))
			if upperResult != lowerResult {
				t.Errorf("Case sensitivity issue for %q: upper=%v, lower=%v", domain, upperResult, lowerResult)
			}
		}
	})
}

// FuzzIsFreeDomain tests IsFreeDomain with random inputs.
func FuzzIsFreeDomain(f *testing.F) {
	seeds := []string{
		"gmail.com",
		"GMAIL.COM",
		"  gmail.com  ",
		"mail.gmail.com",
		"outlook.com",
		"yahoo.com",
		"example.com",
		"",
		".",
		strings.Repeat("x", 500),
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, domain string) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("IsFreeDomain panicked with input %q: %v", domain, r)
			}
		}()

		result := workemailvalidator.IsFreeDomain(domain)

		// Deterministic check
		result2 := workemailvalidator.IsFreeDomain(domain)
		if result != result2 {
			t.Errorf("IsFreeDomain not deterministic for %q", domain)
		}

		// Case insensitivity check
		if domain != "" {
			upperResult := workemailvalidator.IsFreeDomain(strings.ToUpper(domain))

			lowerResult := workemailvalidator.IsFreeDomain(strings.ToLower(domain))
			if upperResult != lowerResult {
				t.Errorf("Case sensitivity issue for %q", domain)
			}
		}
	})
}

// FuzzIsBusinessDomain tests IsBusinessDomain with random inputs.
func FuzzIsBusinessDomain(f *testing.F) {
	seeds := []string{
		"example.com",
		"mycompany.com",
		"gmail.com",
		"temp-mail.com",
		"",
		"x",
		strings.Repeat("business", 100),
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, domain string) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("IsBusinessDomain panicked with input %q: %v", domain, r)
			}
		}()

		business := workemailvalidator.IsBusinessDomain(domain)
		disposableOrFree := workemailvalidator.IsDisposableOrFreeDomain(domain)

		// Business domain should be opposite of disposable or free, except for empty domains
		trimmed := strings.TrimSpace(domain)
		if trimmed != "" {
			if business == disposableOrFree {
				t.Errorf("Inconsistency for %q: business=%v, disposableOrFree=%v (should be opposite)",
					domain, business, disposableOrFree)
			}
		} else {
			// Empty domains should return false for business
			if business {
				t.Errorf("IsBusinessDomain returned true for empty domain %q", domain)
			}
		}

		// Deterministic check
		business2 := workemailvalidator.IsBusinessDomain(domain)
		if business != business2 {
			t.Errorf("IsBusinessDomain not deterministic for %q", domain)
		}
	})
}

// FuzzIsDisposableOrFreeDomain tests the combined function.
func FuzzIsDisposableOrFreeDomain(f *testing.F) {
	seeds := []string{
		"gmail.com",
		"temp-mail.com",
		"example.com",
		"",
		"test",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, domain string) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("IsDisposableOrFreeDomain panicked with input %q: %v", domain, r)
			}
		}()

		disposableOrFree := workemailvalidator.IsDisposableOrFreeDomain(domain)
		disposable := workemailvalidator.IsDisposableDomain(domain)
		free := workemailvalidator.IsFreeDomain(domain)

		// Should be true if either disposable or free
		expected := disposable || free
		if disposableOrFree != expected {
			t.Errorf("Inconsistency for %q: IsDisposableOrFreeDomain=%v, but disposable=%v, free=%v",
				domain, disposableOrFree, disposable, free)
		}
	})
}

// FuzzIsWorkEmail tests IsWorkEmail with random inputs
//
//nolint:cyclop
func FuzzIsWorkEmail(f *testing.F) {
	seeds := []string{
		"user@example.com",
		"user@gmail.com",
		"user@temp-mail.com",
		"@",
		"@domain.com",
		"user@",
		"",
		"no-at-sign",
		"multiple@@at.com",
		"user@domain@com",
		strings.Repeat("a", 100) + "@example.com",
		"user@" + strings.Repeat("sub.", 20) + "example.com",
		"名前@example.com", //nolint:gosmopolitan
		"user@münchen.de",
		"user+tag@example.com",
		"first.last@example.com",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, email string) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("IsWorkEmail panicked with input %q: %v", email, r)
			}
		}()

		result := workemailvalidator.IsWorkEmail(email)

		// Deterministic check
		result2 := workemailvalidator.IsWorkEmail(email)
		if result != result2 {
			t.Errorf("IsWorkEmail not deterministic for %q", email)
		}

		// If email is valid work email, extract domain and verify consistency
		if result && strings.Contains(email, "@") {
			atIndex := strings.LastIndexByte(email, '@')
			if atIndex > 0 && atIndex < len(email)-1 {
				domain := email[atIndex+1:]

				business := workemailvalidator.IsBusinessDomain(domain)
				if !business {
					t.Errorf("IsWorkEmail(%q)=true but domain %q is not business", email, domain)
				}
			}
		}

		// If email is not work email due to invalid format, check edge cases
		if !result && strings.Contains(email, "@") {
			atIndex := strings.LastIndexByte(email, '@')
			if atIndex > 0 && atIndex < len(email)-1 {
				// Valid format, so must be non-business domain
				domain := email[atIndex+1:]

				business := workemailvalidator.IsBusinessDomain(domain)
				if business {
					t.Errorf("IsWorkEmail(%q)=false but domain %q is business", email, domain)
				}
			}
		}
	})
}

// FuzzConsistency ensures all functions remain consistent with each other.
func FuzzConsistency(f *testing.F) {
	seeds := []string{
		"gmail.com",
		"temp-mail.com",
		"example.com",
		"mycompany.com",
		"",
		"test",
		"sub.gmail.com",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, domain string) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Consistency test panicked with input %q: %v", domain, r)
			}
		}()

		disposable := workemailvalidator.IsDisposableDomain(domain)
		free := workemailvalidator.IsFreeDomain(domain)
		disposableOrFree := workemailvalidator.IsDisposableOrFreeDomain(domain)
		business := workemailvalidator.IsBusinessDomain(domain)

		// disposableOrFree should equal disposable OR free
		if disposableOrFree != (disposable || free) {
			t.Errorf("IsDisposableOrFreeDomain inconsistent for %q", domain)
		}

		// business should be opposite of disposableOrFree, except for empty/invalid domains
		// where both can be false
		trimmed := strings.TrimSpace(domain)
		if trimmed != "" {
			if business == disposableOrFree {
				t.Errorf("IsBusinessDomain inconsistent for %q", domain)
			}
		} else {
			// Empty/whitespace domains should return false for business
			if business {
				t.Errorf("IsBusinessDomain returned true for empty domain %q", domain)
			}
		}

		// A domain should not be both disposable and free (data integrity check)
		if disposable && free {
			t.Logf("Warning: %q marked as both disposable and free", domain)
		}
	})
}
