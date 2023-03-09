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
*/
// main bonus

var sum int
var sumPtr *int = &sum

func BenchmarkGetMainBonusBronze(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getMainBonus(ctx, db, "user1", 2, 2, sumPtr)
	}
	sum = 0
}

func BenchmarkGetMainBonusBronzePro(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getMainBonus(ctx, db, "user1", 2, 3, sumPtr)
	}
	sum = 0
}

func BenchmarkGetMainBonusSilver(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getMainBonus(ctx, db, "user1", 2, 4, sumPtr)
	}
	sum = 0
}

func BenchmarkGetMainBonusSilverPro(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getMainBonus(ctx, db, "user1", 2, 5, sumPtr)
	}
	sum = 0
}

func BenchmarkGetMainBonusGold(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getMainBonus(ctx, db, "user1", 2, 6, sumPtr)
	}
	sum = 0
}

func BenchmarkGetMainBonusGoldPro(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getMainBonus(ctx, db, "user1", 2, 7, sumPtr)
	}
	sum = 0
}

func BenchmarkGetMainBonusPlatinum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getMainBonus(ctx, db, "user1", 2, 8, sumPtr)
	}
	sum = 0
}

func BenchmarkGetMainBonusPlatinimPro(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getMainBonus(ctx, db, "user1", 2, 9, sumPtr)
	}
	sum = 0
}

func BenchmarkGetMainBonusDiamond(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getMainBonus(ctx, db, "user1", 2, 10, sumPtr)
	}
	sum = 0
}
