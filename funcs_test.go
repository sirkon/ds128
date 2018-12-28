package ds128

import "testing"

func TestMul128by64(t *testing.T) {
	type args struct {
		lo uint64
		hi uint64
		v  uint64
	}
	tests := []struct {
		name   string
		args   args
		wantLo uint64
		wantHi uint64
	}{
		{
			name: "trivial",
			args: args{
				lo: 12,
				hi: 13,
				v:  2,
			},
			wantLo: 24,
			wantHi: 26,
		},
		{
			name: "basic-overflow",
			args: args{
				lo: 0x8000000000000000,
				hi: 0,
				v:  2,
			},
			wantLo: 0,
			wantHi: 1,
		},
		{
			name: "generic-overflow",
			args: args{
				lo: 0x8000000000000000,
				hi: 2,
				v:  2,
			},
			wantLo: 0,
			wantHi: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResLo, gotResHi := Mul64(tt.args.lo, tt.args.hi, tt.args.v)
			pureLo, pureHi := PureMul(tt.args.lo, tt.args.hi, tt.args.v)
			if gotResLo != pureLo {
				t.Errorf("gotResLo = %v, pureLo = %v", gotResLo, pureLo)
			}
			if gotResHi != pureHi {
				t.Errorf("gotResHi = %v, pureHi = %v", gotResHi, pureHi)
			}
			if gotResLo != tt.wantLo {
				t.Errorf("Mul64() gotResLo = %v, want %v", gotResLo, tt.wantLo)
			}
			if gotResHi != tt.wantHi {
				t.Errorf("Mul64() gotResHi = %v, want %v", gotResHi, tt.wantHi)
			}
		})
	}
}

func serviceMul(a, b uint64) (uint64, uint64) {
	const w = 32
	const m = 1<<w - 1
	ahi, bhi, alo, blo := a>>w, b>>w, a&m, b&m
	lo := alo * blo
	mid1 := alo * bhi
	mid2 := ahi * blo
	lo, c1 := addUint128_64(lo, mid1<<w)
	lo, c2 := addUint128_64(lo, mid2<<w)
	hi, _ := addUint128_64(ahi*bhi, mid1>>w+mid2>>w+c1+c2)
	return lo, hi

}

func addUint128_64(a uint64, b uint64) (lo uint64, hi uint64) {
	lo = a + b
	if lo < a {
		hi = 1
	}
	return
}

func PureMul(lo, hi, v uint64) (uint64, uint64) {
	resLo, addHi := serviceMul(lo, v)
	resHi, _ := serviceMul(hi, v)
	return resLo, resHi + addHi
}

func BenchmarkMul64(b *testing.B) {
	for n := 0; n < 1000000; n++ {
		lo := uint64(n)
		if n&1 == 0 {
			lo = ^lo
		}
		for i := 0; i < b.N; i++ {
			Mul64(lo, uint64(n&5), uint64(n&11))
		}
	}
}

func BenchmarkPureMul(b *testing.B) {
	for n := 0; n < 1000000; n++ {
		lo := uint64(n)
		if n&1 == 0 {
			lo = ^lo
		}
		for i := 0; i < b.N; i++ {
			PureMul(lo, uint64(n&5), uint64(n&11))
		}
	}
}

func TestBenchmarkValidity(t *testing.T) {
	for n := 0; n < 1000000; n++ {
		lo := uint64(n)
		if n&1 == 0 {
			lo = ^lo
		}
		hi := uint64(n & 5)
		v := uint64(n & 11)
		lo1, hi1 := Mul64(lo, hi, v)
		lo2, hi2 := PureMul(lo, hi, v)
		if lo1 != lo2 || hi1 != hi2 {
			t.Fatalf("Mul64 != PureMul on (%v, %v, %v): (%v, %v) != (%v, %v)", lo, hi, v, lo1, hi1, lo2, hi2)
		}
	}
}
