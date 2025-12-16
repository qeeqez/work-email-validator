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
	disposableDomains map[string]bool
	freeDomains       map[string]bool
)

func init() {
	disposableDomains = loadDomains(disposableDomainsData)
	freeDomains = loadDomains(freeDomainsData)
}

func loadDomains(data string) map[string]bool {
	domains := make(map[string]bool)
	lines := strings.Split(data, "\n")
	for _, line := range lines {
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
	for i := 0; i < len(parts)-1; i++ {
		parent := strings.Join(parts[i:], ".")
		if domainMap[parent] {
			return true
		}
	}
	return false
}

func IsDisposableDomain(domain string) bool {
	return isDomainOrParentInMap(domain, disposableDomains)
}

func IsFreeDomain(domain string) bool {
	return isDomainOrParentInMap(domain, freeDomains)
}

func IsDisposableOrFreeDomain(domain string) bool {
	return IsDisposableDomain(domain) || IsFreeDomain(domain)
}

func IsBusinessDomain(domain string) bool {
	return !IsDisposableOrFreeDomain(domain)
}

func IsWorkEmail(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	domain := parts[1]
	return IsBusinessDomain(domain)
}
