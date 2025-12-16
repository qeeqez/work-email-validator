// Package workemailvalidator provides email validation utilities to determine
// whether an email address is from a disposable, free, or business domain.
package workemailvalidator

import (
	_ "embed"
	"strings"
)

//go:embed data/disposable_domains.txt
var disposableDomainsData string

//go:embed data/free_domains.txt
var freeDomainsData string

var (
	disposableDomains = loadDomains(disposableDomainsData)
	freeDomains       = loadDomains(freeDomainsData)
)

func loadDomains(data string) map[string]bool {
	domains := make(map[string]bool)

	for line := range strings.SplitSeq(data, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		domains[strings.ToLower(line)] = true
	}

	return domains
}

func normalizeDomain(domain string) string {
	return strings.ToLower(strings.TrimSpace(domain))
}

func isDomainOrParentInMap(domain string, domainMap map[string]bool) bool {
	domain = normalizeDomain(domain)

	parts := strings.Split(domain, ".")
	for i := range len(parts) - 1 {
		parent := strings.Join(parts[i:], ".")
		if domainMap[parent] {
			return true
		}
	}

	return false
}

// IsDisposableDomain checks if the given domain is a disposable/temporary email domain.
func IsDisposableDomain(domain string) bool {
	return isDomainOrParentInMap(domain, disposableDomains)
}

// IsFreeDomain checks if the given domain is a free email provider domain.
func IsFreeDomain(domain string) bool {
	return isDomainOrParentInMap(domain, freeDomains)
}

// IsDisposableOrFreeDomain checks if the given domain is either disposable or free.
func IsDisposableOrFreeDomain(domain string) bool {
	return IsDisposableDomain(domain) || IsFreeDomain(domain)
}

// IsBusinessDomain checks if the given domain is neither disposable nor free,
// indicating it's likely a business/corporate domain.
func IsBusinessDomain(domain string) bool {
	return !IsDisposableOrFreeDomain(domain)
}

// IsWorkEmail checks if the given email address is from a business domain.
// It returns true if the email is from a domain that is neither disposable nor free.
func IsWorkEmail(email string) bool {
	const expectedParts = 2

	parts := strings.Split(email, "@")
	if len(parts) != expectedParts {
		return false
	}

	domain := parts[1]

	return IsBusinessDomain(domain)
}
