package main

import (
	"testing"
)

var db, _ = dbConn()

func BenchmarkCompressionTreversal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetPersonalVolumes(nil, db)
	}
}

func BenchmarkTreversalWithoutCompression(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetPersonalVolumesWithoutCompression(nil, db)
	}
}
