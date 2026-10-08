/*
exercise:

	benchmark
	Compare string concatenation in a loop via += and via strings.Builder; save the results in a note
*/
package benchmark

import (
	"strings"
	"testing"
)

func BenchmarkStringConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := ""

		for j := 0; j < 100; j++ {
			result += "hello"
		}

		_ = result
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var builder strings.Builder

		for j := 0; j < 100; j++ {
			builder.WriteString("hello")
		}

		_ = builder.String()
	}
}
