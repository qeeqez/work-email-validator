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
