package workemailvalidator

import "testing"

func BenchmarkIsDisposableDomain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsDisposableDomain("temp-mail.com")
	}
}

func BenchmarkIsFreeDomain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsFreeDomain("gmail.com")
	}
}

func BenchmarkIsBusinessDomain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsBusinessDomain("example.com")
	}
}
