package workemailvalidator_test

import (
	"strings"
	"testing"

	workemailvalidator "github.com/qeeqez/work-email-validator"
)

// Benchmark best case - exact match in map.
func BenchmarkIsDisposableDomain_ExactMatch(b *testing.B) {
	for b.Loop() {
		workemailvalidator.IsDisposableDomain("temp-mail.com")
	}
}

// Benchmark worst case - needs to scan entire domain for dots.
func BenchmarkIsDisposableDomain_NoMatch(b *testing.B) {
	for b.Loop() {
		workemailvalidator.IsDisposableDomain("very.long.subdomain.that.is.not.disposable.example.com")
	}
}

// Benchmark subdomain match.
func BenchmarkIsDisposableDomain_SubdomainMatch(b *testing.B) {
	for b.Loop() {
		workemailvalidator.IsDisposableDomain("sub.sub.sub.temp-mail.com")
	}
}

// Benchmark with whitespace (requires trimming).
func BenchmarkIsDisposableDomain_WithWhitespace(b *testing.B) {
	for b.Loop() {
		workemailvalidator.IsDisposableDomain("  temp-mail.com  ")
	}
}

// Benchmark with case conversion.
func BenchmarkIsDisposableDomain_UpperCase(b *testing.B) {
	for b.Loop() {
		workemailvalidator.IsDisposableDomain("TEMP-MAIL.COM")
	}
}

// Benchmark free domain best case.
func BenchmarkIsFreeDomain_ExactMatch(b *testing.B) {
	for b.Loop() {
		workemailvalidator.IsFreeDomain("gmail.com")
	}
}

// Benchmark free domain worst case.
func BenchmarkIsFreeDomain_NoMatch(b *testing.B) {
	for b.Loop() {
		workemailvalidator.IsFreeDomain("a.very.long.business.domain.example.com")
	}
}

// Benchmark free domain subdomain.
func BenchmarkIsFreeDomain_SubdomainMatch(b *testing.B) {
	for b.Loop() {
		workemailvalidator.IsFreeDomain("mail.gmail.com")
	}
}

// Benchmark business domain (checks both maps).
func BenchmarkIsBusinessDomain_True(b *testing.B) {
	for b.Loop() {
		workemailvalidator.IsBusinessDomain("example.com")
	}
}

// Benchmark business domain false (disposable).
func BenchmarkIsBusinessDomain_FalseDisposable(b *testing.B) {
	for b.Loop() {
		workemailvalidator.IsBusinessDomain("temp-mail.com")
	}
}

// Benchmark business domain false (free).
func BenchmarkIsBusinessDomain_FalseFree(b *testing.B) {
	for b.Loop() {
		workemailvalidator.IsBusinessDomain("gmail.com")
	}
}

// Benchmark combined check.
func BenchmarkIsDisposableOrFreeDomain_Disposable(b *testing.B) {
	for b.Loop() {
		workemailvalidator.IsDisposableOrFreeDomain("temp-mail.com")
	}
}

func BenchmarkIsDisposableOrFreeDomain_Free(b *testing.B) {
	for b.Loop() {
		workemailvalidator.IsDisposableOrFreeDomain("gmail.com")
	}
}

func BenchmarkIsDisposableOrFreeDomain_Neither(b *testing.B) {
	for b.Loop() {
		workemailvalidator.IsDisposableOrFreeDomain("example.com")
	}
}

// Benchmark work email validation.
func BenchmarkIsWorkEmail_Valid(b *testing.B) {
	for b.Loop() {
		workemailvalidator.IsWorkEmail("user@example.com")
	}
}

func BenchmarkIsWorkEmail_InvalidFree(b *testing.B) {
	for b.Loop() {
		workemailvalidator.IsWorkEmail("user@gmail.com")
	}
}

func BenchmarkIsWorkEmail_InvalidDisposable(b *testing.B) {
	for b.Loop() {
		workemailvalidator.IsWorkEmail("user@temp-mail.com")
	}
}

func BenchmarkIsWorkEmail_InvalidFormat(b *testing.B) {
	for b.Loop() {
		workemailvalidator.IsWorkEmail("not-an-email")
	}
}

// Benchmark with very long email.
func BenchmarkIsWorkEmail_LongEmail(b *testing.B) {
	email := strings.Repeat("a", 100) + "@" + strings.Repeat("subdomain.", 10) + "example.com"

	b.ResetTimer()

	for b.Loop() {
		workemailvalidator.IsWorkEmail(email)
	}
}

// Benchmark with subdomain of free provider.
func BenchmarkIsWorkEmail_SubdomainFree(b *testing.B) {
	for b.Loop() {
		workemailvalidator.IsWorkEmail("user@mail.gmail.com")
	}
}

// Benchmark with subdomain of business.
func BenchmarkIsWorkEmail_SubdomainBusiness(b *testing.B) {
	for b.Loop() {
		workemailvalidator.IsWorkEmail("user@api.mycompany.com")
	}
}

// Benchmark parallel execution patterns.
func BenchmarkIsWorkEmail_Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			workemailvalidator.IsWorkEmail("user@example.com")
		}
	})
}

func BenchmarkIsDisposableDomain_Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			workemailvalidator.IsDisposableDomain("temp-mail.com")
		}
	})
}

func BenchmarkIsFreeDomain_Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			workemailvalidator.IsFreeDomain("gmail.com")
		}
	})
}

// Benchmark mixed workload (realistic usage).
func BenchmarkMixedWorkload(b *testing.B) {
	domains := []string{
		"user@example.com",
		"user@gmail.com",
		"user@temp-mail.com",
		"admin@mycompany.com",
		"test@outlook.com",
	}

	b.ResetTimer()

	for b.Loop() {
		for _, domain := range domains {
			workemailvalidator.IsWorkEmail(domain)
		}
	}
}

// Benchmark to measure allocation overhead.
func BenchmarkIsWorkEmail_Allocations(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		workemailvalidator.IsWorkEmail("user@example.com")
	}
}

func BenchmarkIsDisposableDomain_Allocations(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		workemailvalidator.IsDisposableDomain("temp-mail.com")
	}
}

func BenchmarkIsFreeDomain_Allocations(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		workemailvalidator.IsFreeDomain("gmail.com")
	}
}
