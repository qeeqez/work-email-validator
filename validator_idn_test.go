package workemailvalidator_test

import (
	"testing"

	workemailvalidator "github.com/qeeqez/work-email-validator"
)

// TestInternationalizedDomains tests support for IDN (Internationalized Domain Names).
func TestInternationalizedDomains(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		domain   string
		expected bool
	}{
		// German domains with umlauts
		{"german_umlaut", "münchen.de", true},
		{"german_business", "bücher.com", true},
		{"german_ae", "äpfel.de", true},

		// Cyrillic domains
		{"cyrillic_russian", "москва.рф", true},
		{"cyrillic_ukraine", "київ.ua", true},
		{"cyrillic_example", "пример.com", true},

		// Chinese domains
		{"chinese_simplified", "例子.中国", true},  //nolint:gosmopolitan
		{"chinese_traditional", "範例.台灣", true}, //nolint:gosmopolitan
		{"chinese_mixed", "测试.com", true},      //nolint:gosmopolitan

		// Arabic domains
		{"arabic_example", "مثال.com", true},
		{"arabic_saudi", "السعودية.sa", true},

		// Japanese domains
		{"japanese_hiragana", "にほん.jp", true},
		{"japanese_katakana", "テスト.jp", true},
		{"japanese_kanji", "日本.jp", true}, //nolint:gosmopolitan

		// Greek domains
		{"greek_example", "παράδειγμα.gr", true},

		// Hebrew domains
		{"hebrew_example", "דוגמה.com", true},

		// Mixed script (should still work)
		{"mixed_english_cyrillic", "test-тест.com", true},

		// Emoji domains (these exist!)
		{"emoji_heart", "❤️.com", true},
		{"emoji_smile", "😀.com", true},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			// Test IsBusinessDomain - all these should be business domains
			// (none should be in free/disposable lists)
			result := workemailvalidator.IsBusinessDomain(testCase.domain)
			if result != testCase.expected {
				t.Errorf("IsBusinessDomain(%q) = %v, want %v", testCase.domain, result, testCase.expected)
			}

			// They should NOT be free or disposable
			if workemailvalidator.IsFreeDomain(testCase.domain) {
				t.Errorf("IsFreeDomain(%q) = true, but internationalized domain shouldn't be in free list", testCase.domain)
			}

			if workemailvalidator.IsDisposableDomain(testCase.domain) {
				t.Errorf("IsDisposableDomain(%q) = true, but internationalized domain shouldn't be in disposable list",
					testCase.domain)
			}
		})
	}
}

// TestInternationalizedEmails tests email validation with IDN.
func TestInternationalizedEmails(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		// German
		{"german_umlaut_domain", "user@münchen.de", true},
		{"german_umlaut_local", "müller@example.com", true},

		// Cyrillic
		{"cyrillic_domain", "user@москва.рф", true},
		{"cyrillic_local", "пользователь@example.com", true},

		// Chinese
		{"chinese_domain", "user@例子.中国", true},    //nolint:gosmopolitan
		{"chinese_local", "用户@example.com", true}, //nolint:gosmopolitan

		// Japanese
		{"japanese_domain", "user@日本.jp", true}, //nolint:gosmopolitan
		{"japanese_local", "ユーザー@example.com", true},

		// Mixed
		{"mixed_local_domain", "user@тест.com", true},
		{"both_internationalized", "пользователь@москва.рф", true},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			result := workemailvalidator.IsWorkEmail(testCase.email)
			if result != testCase.expected {
				t.Errorf("IsWorkEmail(%q) = %v, want %v", testCase.email, result, testCase.expected)
			}
		})
	}
}

// TestIDNEquivalence tests that IDN domains are normalized correctly.
func TestIDNEquivalence(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name               string
		domain1            string
		domain2            string
		shouldBeEquivalent bool
	}{
		// German umlaut normalization
		{"german_ae_vs_a", "müller.com", "muller.com", false}, // ü ≠ u
		{"german_same", "münchen.de", "MÜNCHEN.DE", true},

		// Cyrillic normalization
		{"cyrillic_case", "москва.рф", "МОСКВА.РФ", true},

		// Mixed case should normalize the same
		{"mixed_case", "Test.COM", "test.com", true},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			result1 := workemailvalidator.IsBusinessDomain(testCase.domain1)
			result2 := workemailvalidator.IsBusinessDomain(testCase.domain2)

			if testCase.shouldBeEquivalent {
				if result1 != result2 {
					t.Errorf("Domains %q and %q should be equivalent, but got %v != %v",
						testCase.domain1, testCase.domain2, result1, result2)
				}
			}
		})
	}
}

// TestIDNPunycodeConversion tests that Punycode conversion works.
func TestIDNPunycodeConversion(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		unicodeDomain  string
		punycodeDomain string
	}{
		// These should be treated as equivalent
		{"german", "münchen.de", "xn--mnchen-3ya.de"},
		{"russian", "москва.рф", "xn--80adxhks.xn--p1ai"},
		{"chinese", "例子.中国", "xn--fsqu00a.xn--fiqs8s"}, //nolint:gosmopolitan
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			// Both unicode and punycode versions should give same result
			unicodeResult := workemailvalidator.IsBusinessDomain(testCase.unicodeDomain)
			punycodeResult := workemailvalidator.IsBusinessDomain(testCase.punycodeDomain)

			if unicodeResult != punycodeResult {
				t.Errorf("Unicode %q and Punycode %q should give same result, but got %v != %v",
					testCase.unicodeDomain, testCase.punycodeDomain, unicodeResult, punycodeResult)
			}
		})
	}
}

// TestIDNInvalidDomains tests that invalid IDN domains are handled gracefully.
func TestIDNInvalidDomains(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		domain string
	}{
		{"only_emoji", "😀"},
		{"emoji_no_tld", "❤️"},
		{"invalid_unicode", "\x00\x01\x02.com"},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			// These should not panic and should return false
			result := workemailvalidator.IsBusinessDomain(testCase.domain)
			if result {
				t.Errorf("IsBusinessDomain(%q) = true, but should be false for invalid domain", testCase.domain)
			}
		})
	}
}
