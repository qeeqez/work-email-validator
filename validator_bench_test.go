package workemailvalidator_test

import (
	"testing"

	workemailvalidator "github.com/qeeqez/work-email-validator"
)

func BenchmarkIsDisposableDomain(b *testing.B) {
	for b.Loop() {
		workemailvalidator.IsDisposableDomain("temp-mail.com")
	}
}

func BenchmarkIsFreeDomain(b *testing.B) {
	for b.Loop() {
		workemailvalidator.IsFreeDomain("gmail.com")
	}
}

func BenchmarkIsBusinessDomain(b *testing.B) {
	for b.Loop() {
		workemailvalidator.IsBusinessDomain("example.com")
	}
}
