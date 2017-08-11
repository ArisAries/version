package version

import (
	"testing"
)

func BenchmarkVersionParser_GetVersion(b *testing.B) {
	for i := 0; i < b.N;  i ++{
		GetGlobalVersionParser().GetVersion("1.1.1")
	}
}

func BenchmarkCompareVersion(b *testing.B) {
	for i := 0; i < b.N;  i ++ {
		CompareVersion("1.1.1.1", "v2.1.1.1")
	}
}

func BenchmarkSetPrefix(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		SetPrefix("v")
	}
}