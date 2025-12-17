package workemailvalidator_test

import (
	"strings"
	"testing"

	workemailvalidator "github.com/qeeqez/work-email-validator"
)

// testCase represents a generic test case for domain/email validation.
type testCase struct {
	name     string
	input    string
	expected bool
}

// runDomainTests is a DRY helper that runs domain validation tests.
func runDomainTests(t *testing.T, tests []testCase, validatorFunc func(string) bool) {
	t.Helper()

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			result := validatorFunc(testCase.input)
			if result != testCase.expected {
				t.Errorf("got %v, want %v for input %q", result, testCase.expected, testCase.input)
			}
		})
	}
}

// Edge case tests for IsDisposableDomain.
func TestIsDisposableDomain(t *testing.T) {
	t.Parallel()

	tests := []testCase{
		// Known disposable domains
		{"known_disposable", "temp-mail.com", true},
		{"known_disposable_2", "10minutemail.com", true},
		{"known_disposable_3", "guerrillamail.com", true},

		// Case insensitivity
		{"uppercase", "TEMP-MAIL.COM", true},
		{"mixedcase", "TeMp-MaIl.CoM", true},

		// Whitespace handling
		{"leading_whitespace", "  temp-mail.org", true},
		{"trailing_whitespace", "temp-mail.org  ", true},
		{"both_whitespace", "  temp-mail.org  ", true},
		{"tabs_and_spaces", "\t temp-mail.org \t", true},

		// Non-disposable domains
		{"gmail_not_disposable", "gmail.com", false},
		{"business_domain", "example.com", false},
		{"corporate_domain", "mycompany.com", false},

		// Edge cases
		{"empty_string", "", false},
		{"single_char", "a", false},
		{"no_tld", "domain", false},
		{"just_dot", ".", false},
		{"multiple_dots", "...", false},
		{"leading_dot", ".domain.com", false},
		{"trailing_dot", "domain.com.", false},

		// Subdomain tests (if parent is disposable)
		{"subdomain_of_disposable", "sub.temp-mail.com", true},
		{"deep_subdomain_disposable", "a.b.c.temp-mail.com", true},

		// Special characters
		{"hyphen_in_domain", "my-domain.com", false},
		{"numbers_in_domain", "123domain.com", false},
		{"mixed_alphanumeric", "test123-domain456.com", false},

		// Very long inputs
		{"long_domain", strings.Repeat("a", 100) + ".com", false},
		{"long_subdomain", strings.Repeat("sub.", 20) + "domain.com", false},

		// Unicode domains (IDN)
		{"unicode_domain", "münchen.de", false},
		{"emoji_domain", "😀.com", false},
	}

	runDomainTests(t, tests, workemailvalidator.IsDisposableDomain)
}

// Edge case tests for IsFreeDomain.
func TestIsFreeDomain(t *testing.T) {
	t.Parallel()

	tests := []testCase{
		// Known free domains
		{"gmail", "gmail.com", true},
		{"outlook", "outlook.com", true},
		{"yahoo", "yahoo.com", true},
		{"hotmail", "hotmail.com", true},
		{"icloud", "icloud.com", true},
		{"protonmail", "protonmail.com", true},

		// Case insensitivity
		{"uppercase_gmail", "GMAIL.COM", true},
		{"mixedcase_outlook", "OuTlOoK.cOm", true},

		// Whitespace handling
		{"whitespace_yahoo", "  yahoo.com  ", true},
		{"tabs_hotmail", "\t\thotmail.com\t\t", true},

		// Non-free domains
		{"business_domain", "example.com", false},
		{"disposable_not_free", "temp-mail.com", false},
		{"corporate", "mycompany.com", false},

		// Edge cases
		{"empty_string", "", false},
		{"single_letter", "g", false},
		{"incomplete_domain", "gma", false},
		{"typo_gmail", "gmial.com", false},

		// Subdomains (if parent is free)
		{"subdomain_gmail", "mail.gmail.com", true},
		{"subdomain_outlook", "accounts.outlook.com", true},
		{"deep_subdomain_yahoo", "a.b.yahoo.com", true},

		// Edge case subdomains
		{"empty_subdomain", ".gmail.com", true},
		{"numeric_subdomain", "123.gmail.com", true},
	}

	runDomainTests(t, tests, workemailvalidator.IsFreeDomain)
}

// Edge case tests for IsDisposableOrFreeDomain.
func TestIsDisposableOrFreeDomain(t *testing.T) {
	t.Parallel()

	tests := []testCase{
		// Free domains
		{"free_gmail", "gmail.com", true},
		{"free_outlook", "outlook.com", true},

		// Disposable domains
		{"disposable_tempmail", "temp-mail.com", true},
		{"disposable_10min", "10minutemail.com", true},

		// Business domains (neither)
		{"business_example", "example.com", false},
		{"business_corporate", "mycompany.com", false},
		{"business_custom", "custom-business.io", false},

		// Edge cases
		{"empty", "", false},
		{"whitespace_only", "   ", false},
		{"invalid_format", "not a domain", false},

		// Mixed cases
		{"uppercase_free", "GMAIL.COM", true},
		{"whitespace_disposable", "  temp-mail.com  ", true},

		// Subdomains
		{"subdomain_free", "mail.gmail.com", true},
		{"subdomain_disposable", "x.temp-mail.com", true},
		{"subdomain_business", "api.mycompany.com", false},
	}

	runDomainTests(t, tests, workemailvalidator.IsDisposableOrFreeDomain)
}

// Edge case tests for IsBusinessDomain.
func TestIsBusinessDomain(t *testing.T) {
	t.Parallel()

	tests := []testCase{
		// Business domains
		{"simple_business", "example.com", true},
		{"corporate", "mycompany.com", true},
		{"startup", "startupname.io", true},
		{"enterprise", "bigcorp.net", true},

		// Non-business (free)
		{"free_gmail", "gmail.com", false},
		{"free_outlook", "outlook.com", false},

		// Non-business (disposable)
		{"disposable_tempmail", "temp-mail.com", false},
		{"disposable_guerrilla", "guerrillamail.com", false},

		// Edge cases - invalid domain syntax should return false
		{"empty", "", false},
		{"single_char", "x", false},
		{"tld_only", ".com", false},
		{"single_char_tld", "domain.a", false}, // TLD must be at least 2 chars

		// Subdomains of business domains
		{"business_subdomain", "api.mycompany.com", true},
		{"business_deep_sub", "v1.api.mycompany.com", true},

		// Subdomains of free/disposable (should be non-business)
		{"free_subdomain", "custom.gmail.com", false},
		{"disposable_subdomain", "test.temp-mail.com", false},

		// Case and whitespace
		{"uppercase_business", "EXAMPLE.COM", true},
		{"whitespace_business", "  example.com  ", true},
	}

	runDomainTests(t, tests, workemailvalidator.IsBusinessDomain)
}

// Edge case tests for IsWorkEmail.
func TestIsWorkEmail(t *testing.T) {
	t.Parallel()

	tests := []testCase{
		// Valid work emails
		{"valid_work", "user@mycompany.com", true},
		{"valid_work_2", "contact@example.com", true},
		{"valid_work_subdomain", "admin@mail.company.com", true},

		// Non-work (free)
		{"free_gmail", "user@gmail.com", false},
		{"free_outlook", "user@outlook.com", false},
		{"free_yahoo", "user@yahoo.com", false},

		// Non-work (disposable)
		{"disposable", "user@temp-mail.com", false},
		{"disposable_2", "test@10minutemail.com", false},

		// Invalid email formats
		{"no_at_sign", "invalid-email", false},
		{"empty_string", "", false},
		{"only_at", "@", false},
		{"at_at_start", "@domain.com", false},
		{"at_at_end", "user@", false},
		// Multiple @ signs - domain extraction uses LastIndexByte
		// "user@@domain.com" -> "domain.com" (valid and business)
		// "user@domain@com" -> "com" (invalid - TLD only, fails validation)
		{"multiple_at", "user@@domain.com", true},
		{"multiple_at_2", "user@domain@com", false},

		// Edge cases with @ symbol
		{"just_domain", "domain.com", false},
		{"no_local_part", "@domain.com", false},
		{"no_domain_part", "user@", false},

		// Whitespace
		{"whitespace_domain", "user@  example.com  ", true},
		{"whitespace_full", "  user@example.com  ", true},

		// Case insensitivity
		{"uppercase_domain", "user@EXAMPLE.COM", true},
		{"uppercase_free", "user@GMAIL.COM", false},

		// Special characters in local part (should still validate domain)
		{"plus_addressing", "user+tag@example.com", true},
		{"dots_in_local", "first.last@example.com", true},
		{"underscore", "user_name@example.com", true},
		{"hyphen", "user-name@example.com", true},

		// Subdomain edge cases
		{"subdomain_business", "user@mail.mycompany.com", true},
		{"subdomain_free", "user@mail.gmail.com", false},
		{"subdomain_disposable", "user@x.temp-mail.com", false},

		// Multiple @ signs - uses LastIndexByte, so extracts domain after last @
		// "user@host@company.com" extracts "company.com" which is valid and business
		{"email_with_at_in_local", "user@host@company.com", true},

		// Very long emails
		{"long_local_part", strings.Repeat("a", 64) + "@example.com", true},
		{"long_domain", "user@" + strings.Repeat("sub.", 10) + "example.com", true},

		// Unicode
		{"unicode_local", "名前@example.com", true}, //nolint:gosmopolitan
		{"unicode_domain", "user@münchen.de", true},
	}

	runDomainTests(t, tests, workemailvalidator.IsWorkEmail)
}

// Test subdomain hierarchy handling.
func TestSubdomainHierarchy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		domain     string
		disposable bool
		free       bool
		business   bool
	}{
		{"root_disposable", "temp-mail.com", true, false, false},
		{"sub_disposable", "test.temp-mail.com", true, false, false},
		{"deep_sub_disposable", "a.b.c.temp-mail.com", true, false, false},

		{"root_free", "gmail.com", false, true, false},
		{"sub_free", "mail.gmail.com", false, true, false},
		{"deep_sub_free", "x.y.z.gmail.com", false, true, false},

		{"root_business", "example.com", false, false, true},
		{"sub_business", "api.example.com", false, false, true},
		{"deep_sub_business", "v2.api.example.com", false, false, true},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			disposable := workemailvalidator.IsDisposableDomain(testCase.domain)
			free := workemailvalidator.IsFreeDomain(testCase.domain)
			business := workemailvalidator.IsBusinessDomain(testCase.domain)

			if disposable != testCase.disposable {
				t.Errorf("IsDisposableDomain(%q) = %v, want %v", testCase.domain, disposable, testCase.disposable)
			}

			if free != testCase.free {
				t.Errorf("IsFreeDomain(%q) = %v, want %v", testCase.domain, free, testCase.free)
			}

			if business != testCase.business {
				t.Errorf("IsBusinessDomain(%q) = %v, want %v", testCase.domain, business, testCase.business)
			}
		})
	}
}

// Test consistency between functions.
func TestFunctionConsistency(t *testing.T) {
	t.Parallel()

	domains := []string{
		"gmail.com",
		"temp-mail.com",
		"example.com",
		"outlook.com",
		"mycompany.com",
		"",
		"invalid",
	}

	for _, domain := range domains {
		t.Run(domain, func(t *testing.T) {
			t.Parallel()

			disposable := workemailvalidator.IsDisposableDomain(domain)
			free := workemailvalidator.IsFreeDomain(domain)
			disposableOrFree := workemailvalidator.IsDisposableOrFreeDomain(domain)
			business := workemailvalidator.IsBusinessDomain(domain)

			// IsDisposableOrFreeDomain should be true if either disposable or free
			expectedDisposableOrFree := disposable || free
			if disposableOrFree != expectedDisposableOrFree {
				t.Errorf("IsDisposableOrFreeDomain(%q) = %v, but disposable=%v, free=%v",
					domain, disposableOrFree, disposable, free)
			}

			// IsBusinessDomain should be opposite of IsDisposableOrFreeDomain
			// BUT only for VALID domains. Invalid domains (empty, invalid syntax) return false for both

			// Check if domain would pass basic validation
			normalized := strings.ToLower(strings.TrimSpace(domain))
			isLikelyValid := len(normalized) >= 4 && strings.Contains(normalized, ".")

			if isLikelyValid {
				// For valid-looking domains, business should be opposite of disposable/free
				if business == disposableOrFree {
					t.Errorf("IsBusinessDomain(%q) = %v, but IsDisposableOrFreeDomain = %v (should be opposite for valid domains)",
						domain, business, disposableOrFree)
				}
			} else {
				// For invalid domains, business should be false
				// disposableOrFree can be false (if not in lists) or could theoretically be true (if matched)
				if business {
					t.Errorf("IsBusinessDomain(%q) = true, but domain appears invalid", domain)
				}
			}

			// A domain cannot be both disposable and free
			// (this is a logical constraint, not enforced by code but should be true for data)
			if disposable && free {
				t.Logf("Warning: %q is marked as both disposable and free", domain)
			}
		})
	}
}
