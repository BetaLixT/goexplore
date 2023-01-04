package main

import (
	"testing"
)

var num = 1000

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mapRoute("12queries/QueryTask")
	}
}

func BenchmarkSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		switchRoute("12queries/QueryTask")
	}
}
