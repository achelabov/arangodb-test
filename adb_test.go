package main

import (
	"testing"
)

var db, _ = dbConn()

func BenchmarkCompressionTreversal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetPersonalVolume(nil, db)
	}
}
