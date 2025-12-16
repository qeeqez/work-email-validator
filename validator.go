// Package workemailvalidator provides email validation utilities to determine
// whether an email address is from a disposable, free, or business domain.
package workemailvalidator

import (
	_ "embed"
	"strings"
)

// normalize prepares the domain for lookup: trims spaces, converts to ASCII (for IDN), and lowercases.
func normalize(domain string) string {
	domain = strings.TrimSpace(domain)
	domain = domainToASCII(domain)
	return strings.ToLower(domain)
}

// contains checks if the domain or any of its parents exist in the map.
// It expects the domain to be already normalized (lowercase, trimmed).
func contains(domain string, domainMap map[string]struct{}) bool {
	if _, ok := domainMap[domain]; ok {
		return true
	}

	for i := range len(domain) {
		if domain[i] == '.' {
			if _, ok := domainMap[domain[i+1:]]; ok {
				return true
			}
		}
	}

	return false
}

// isValidDomainSyntax checks structure. Assumes domain is trimmed.
func isValidDomainSyntax(domain string) bool {
	// Min 4 chars: "a.go" (1 label + 1 dot + 2 TLD)
	if len(domain) < 4 || len(domain) > 253 {
		return false
	}

	lastDotIndex := strings.LastIndexByte(domain, '.')

	// Must have a dot, not at start/end
	if lastDotIndex <= 0 || lastDotIndex >= len(domain)-1 {
		return false
	}

	const TLDMinLength = 2
	if len(domain[lastDotIndex+1:]) < TLDMinLength {
		return false
	}

	// Check for invalid chars (spaces, controls)
	for i := range len(domain) {
		if domain[i] <= ' ' || domain[i] == 127 {
			return false
		}
	}

	return true
}

// IsDisposableDomain checks if the given domain is a disposable/temporary email domain.
func IsDisposableDomain(domain string) bool {
	return contains(normalize(domain), disposableDomains)
}

// IsFreeDomain checks if the given domain is a free email provider domain.
func IsFreeDomain(domain string) bool {
	return contains(normalize(domain), freeDomains)
}

// IsDisposableOrFreeDomain checks if the given domain is either disposable or free.
func IsDisposableOrFreeDomain(domain string) bool {
	normalized := normalize(domain)
	return contains(normalized, disposableDomains) || contains(normalized, freeDomains)
}

// IsBusinessDomain checks if the given domain is neither disposable nor free.
func IsBusinessDomain(domain string) bool {
	normalized := normalize(domain)

	if !isValidDomainSyntax(normalized) {
		return false
	}

	return !contains(normalized, disposableDomains) && !contains(normalized, freeDomains)
}

// IsWorkEmail checks if the given email address is from a business domain.
func IsWorkEmail(email string) bool {
	atIndex := strings.LastIndexByte(email, '@')

	if atIndex <= 0 || atIndex >= len(email)-1 {
		return false
	}

	domain := email[atIndex+1:]

	return IsBusinessDomain(domain)
}
