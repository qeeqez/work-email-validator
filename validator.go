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

func loadDomains(data string) map[string]struct{} {
	domains := make(map[string]struct{})

	for line := range strings.Lines(data) {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		domains[strings.ToLower(line)] = struct{}{}
	}

	return domains
}

func contains(domain string, domainMap map[string]struct{}) bool {
	domain = strings.TrimSpace(strings.ToLower(domain))

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

// IsDisposableDomain checks if the given domain is a disposable/temporary email domain.
func IsDisposableDomain(domain string) bool {
	return contains(domain, disposableDomains)
}

// IsFreeDomain checks if the given domain is a free email provider domain.
func IsFreeDomain(domain string) bool {
	return contains(domain, freeDomains)
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
	atIndex := strings.LastIndexByte(email, '@')

	if atIndex <= 0 || atIndex >= len(email)-1 {
		return false
	}

	domain := email[atIndex+1:]

	return IsBusinessDomain(domain)
}
