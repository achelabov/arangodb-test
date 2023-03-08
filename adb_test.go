package main

import (
	"context"
	"testing"
)

var db, _ = dbConn()
var ctx = context.Background()

/*
func BenchmarkCompressionTreversal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetPersonalVolumes(ctx, db)
	}
}
*/

func BenchmarkTreversalFrom16To16Lvl(b *testing.B) {
	for n := 0; n < b.N; n++ {
		traversal(ctx, db, "user1", 16, 16)
	}
}

func BenchmarkTreversalFrom13To16Lvl(b *testing.B) {
	for n := 0; n < b.N; n++ {
		traversal(ctx, db, "user1", 13, 16)
	}
}

func BenchmarkTreversalFrom10To13Lvl(b *testing.B) {
	for n := 0; n < b.N; n++ {
		traversal(ctx, db, "user1", 10, 13)
	}
}
func BenchmarkTreversalFrom7To10Lvl(b *testing.B) {
	for n := 0; n < b.N; n++ {
		traversal(ctx, db, "user1", 7, 10)
	}
}

func BenchmarkTreversalFrom1To16Lvl(b *testing.B) {
	for n := 0; n < b.N; n++ {
		traversal(ctx, db, "user1", 1, 16)
	}
}
