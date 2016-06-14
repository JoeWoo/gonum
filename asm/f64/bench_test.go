// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package f64

import (
	"runtime"
	"testing"
)

var (
	a = float64(2)
	x = make([]float64, 1000000)
	y = make([]float64, 1000000)
	z = make([]float64, 1000000)
)

func init() {
	for n := range x {
		x[n] = float64(n)
		y[n] = float64(n)
	}
}

func benchaxpyu(t *testing.B, n int, f func(a float64, x, y []float64)) {
	x, y := x[:n], y[:n]

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		f(a, x, y)
	}
	t.StopTimer()
	runtime.GC()
}

func naiveaxpyu(a float64, x, y []float64) {
	for i, v := range x {
		y[i] += a * v
	}
}

func BenchmarkF64AxpyUnitary1(t *testing.B)     { benchaxpyu(t, 1, AxpyUnitary) }
func BenchmarkF64AxpyUnitary2(t *testing.B)     { benchaxpyu(t, 2, AxpyUnitary) }
func BenchmarkF64AxpyUnitary3(t *testing.B)     { benchaxpyu(t, 3, AxpyUnitary) }
func BenchmarkF64AxpyUnitary4(t *testing.B)     { benchaxpyu(t, 4, AxpyUnitary) }
func BenchmarkF64AxpyUnitary5(t *testing.B)     { benchaxpyu(t, 5, AxpyUnitary) }
func BenchmarkF64AxpyUnitary10(t *testing.B)    { benchaxpyu(t, 10, AxpyUnitary) }
func BenchmarkF64AxpyUnitary100(t *testing.B)   { benchaxpyu(t, 100, AxpyUnitary) }
func BenchmarkF64AxpyUnitary1000(t *testing.B)  { benchaxpyu(t, 1000, AxpyUnitary) }
func BenchmarkF64AxpyUnitary5000(t *testing.B)  { benchaxpyu(t, 5000, AxpyUnitary) }
func BenchmarkF64AxpyUnitary10000(t *testing.B) { benchaxpyu(t, 10000, AxpyUnitary) }
func BenchmarkF64AxpyUnitary50000(t *testing.B) { benchaxpyu(t, 50000, AxpyUnitary) }

func BenchmarkLF64AxpyUnitary1(t *testing.B)     { benchaxpyu(t, 1, naiveaxpyu) }
func BenchmarkLF64AxpyUnitary2(t *testing.B)     { benchaxpyu(t, 2, naiveaxpyu) }
func BenchmarkLF64AxpyUnitary3(t *testing.B)     { benchaxpyu(t, 3, naiveaxpyu) }
func BenchmarkLF64AxpyUnitary4(t *testing.B)     { benchaxpyu(t, 4, naiveaxpyu) }
func BenchmarkLF64AxpyUnitary5(t *testing.B)     { benchaxpyu(t, 5, naiveaxpyu) }
func BenchmarkLF64AxpyUnitary10(t *testing.B)    { benchaxpyu(t, 10, naiveaxpyu) }
func BenchmarkLF64AxpyUnitary100(t *testing.B)   { benchaxpyu(t, 100, naiveaxpyu) }
func BenchmarkLF64AxpyUnitary1000(t *testing.B)  { benchaxpyu(t, 1000, naiveaxpyu) }
func BenchmarkLF64AxpyUnitary5000(t *testing.B)  { benchaxpyu(t, 5000, naiveaxpyu) }
func BenchmarkLF64AxpyUnitary10000(t *testing.B) { benchaxpyu(t, 10000, naiveaxpyu) }
func BenchmarkLF64AxpyUnitary50000(t *testing.B) { benchaxpyu(t, 50000, naiveaxpyu) }

func benchaxpyut(t *testing.B, n int, f func(d []float64, a float64, x, y []float64)) {
	x, y, z := x[:n], y[:n], z[:n]

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		f(z, a, x, y)
	}
	t.StopTimer()
	runtime.GC()
}

func naiveaxpyut(d []float64, a float64, x, y []float64) {
	for i, v := range x {
		d[i] = y[i] + a*v
	}
}

func BenchmarkF64AxpyUnitaryTo1(t *testing.B)     { benchaxpyut(t, 1, AxpyUnitaryTo) }
func BenchmarkF64AxpyUnitaryTo2(t *testing.B)     { benchaxpyut(t, 2, AxpyUnitaryTo) }
func BenchmarkF64AxpyUnitaryTo3(t *testing.B)     { benchaxpyut(t, 3, AxpyUnitaryTo) }
func BenchmarkF64AxpyUnitaryTo4(t *testing.B)     { benchaxpyut(t, 4, AxpyUnitaryTo) }
func BenchmarkF64AxpyUnitaryTo5(t *testing.B)     { benchaxpyut(t, 5, AxpyUnitaryTo) }
func BenchmarkF64AxpyUnitaryTo10(t *testing.B)    { benchaxpyut(t, 10, AxpyUnitaryTo) }
func BenchmarkF64AxpyUnitaryTo100(t *testing.B)   { benchaxpyut(t, 100, AxpyUnitaryTo) }
func BenchmarkF64AxpyUnitaryTo1000(t *testing.B)  { benchaxpyut(t, 1000, AxpyUnitaryTo) }
func BenchmarkF64AxpyUnitaryTo5000(t *testing.B)  { benchaxpyut(t, 5000, AxpyUnitaryTo) }
func BenchmarkF64AxpyUnitaryTo10000(t *testing.B) { benchaxpyut(t, 10000, AxpyUnitaryTo) }
func BenchmarkF64AxpyUnitaryTo50000(t *testing.B) { benchaxpyut(t, 50000, AxpyUnitaryTo) }

func BenchmarkLF64AxpyUnitaryTo1(t *testing.B)     { benchaxpyut(t, 1, naiveaxpyut) }
func BenchmarkLF64AxpyUnitaryTo2(t *testing.B)     { benchaxpyut(t, 2, naiveaxpyut) }
func BenchmarkLF64AxpyUnitaryTo3(t *testing.B)     { benchaxpyut(t, 3, naiveaxpyut) }
func BenchmarkLF64AxpyUnitaryTo4(t *testing.B)     { benchaxpyut(t, 4, naiveaxpyut) }
func BenchmarkLF64AxpyUnitaryTo5(t *testing.B)     { benchaxpyut(t, 5, naiveaxpyut) }
func BenchmarkLF64AxpyUnitaryTo10(t *testing.B)    { benchaxpyut(t, 10, naiveaxpyut) }
func BenchmarkLF64AxpyUnitaryTo100(t *testing.B)   { benchaxpyut(t, 100, naiveaxpyut) }
func BenchmarkLF64AxpyUnitaryTo1000(t *testing.B)  { benchaxpyut(t, 1000, naiveaxpyut) }
func BenchmarkLF64AxpyUnitaryTo5000(t *testing.B)  { benchaxpyut(t, 5000, naiveaxpyut) }
func BenchmarkLF64AxpyUnitaryTo10000(t *testing.B) { benchaxpyut(t, 10000, naiveaxpyut) }
func BenchmarkLF64AxpyUnitaryTo50000(t *testing.B) { benchaxpyut(t, 50000, naiveaxpyut) }

func benchaxpyinc(t *testing.B, ln, t_inc int, f func(alpha float64, x, y []float64, n, incX, incY, ix, iy uintptr)) {
	n, inc := uintptr(ln), uintptr(t_inc)
	var idx int
	if t_inc < 0 {
		idx = (-ln + 1) * t_inc
	}

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		f(1, x, y, n, inc, inc, uintptr(idx), uintptr(idx))
	}
	t.StopTimer()
	runtime.GC()
}

func naiveaxpyinc(alpha float64, x, y []float64, n, incX, incY, ix, iy uintptr) {
	for i := 0; i < int(n); i++ {
		y[iy] += alpha * x[ix]
		ix += incX
		iy += incY
	}
}

func BenchmarkF64AxpyIncN1Inc1(b *testing.B) { benchaxpyinc(b, 1, 1, AxpyInc) }

func BenchmarkF64AxpyIncN2Inc1(b *testing.B)  { benchaxpyinc(b, 2, 1, AxpyInc) }
func BenchmarkF64AxpyIncN2Inc2(b *testing.B)  { benchaxpyinc(b, 2, 2, AxpyInc) }
func BenchmarkF64AxpyIncN2Inc4(b *testing.B)  { benchaxpyinc(b, 2, 4, AxpyInc) }
func BenchmarkF64AxpyIncN2Inc10(b *testing.B) { benchaxpyinc(b, 2, 10, AxpyInc) }

func BenchmarkF64AxpyIncN3Inc1(b *testing.B)  { benchaxpyinc(b, 3, 1, AxpyInc) }
func BenchmarkF64AxpyIncN3Inc2(b *testing.B)  { benchaxpyinc(b, 3, 2, AxpyInc) }
func BenchmarkF64AxpyIncN3Inc4(b *testing.B)  { benchaxpyinc(b, 3, 4, AxpyInc) }
func BenchmarkF64AxpyIncN3Inc10(b *testing.B) { benchaxpyinc(b, 3, 10, AxpyInc) }

func BenchmarkF64AxpyIncN4Inc1(b *testing.B)  { benchaxpyinc(b, 4, 1, AxpyInc) }
func BenchmarkF64AxpyIncN4Inc2(b *testing.B)  { benchaxpyinc(b, 4, 2, AxpyInc) }
func BenchmarkF64AxpyIncN4Inc4(b *testing.B)  { benchaxpyinc(b, 4, 4, AxpyInc) }
func BenchmarkF64AxpyIncN4Inc10(b *testing.B) { benchaxpyinc(b, 4, 10, AxpyInc) }

func BenchmarkF64AxpyIncN10Inc1(b *testing.B)  { benchaxpyinc(b, 10, 1, AxpyInc) }
func BenchmarkF64AxpyIncN10Inc2(b *testing.B)  { benchaxpyinc(b, 10, 2, AxpyInc) }
func BenchmarkF64AxpyIncN10Inc4(b *testing.B)  { benchaxpyinc(b, 10, 4, AxpyInc) }
func BenchmarkF64AxpyIncN10Inc10(b *testing.B) { benchaxpyinc(b, 10, 10, AxpyInc) }

func BenchmarkF64AxpyIncN1000Inc1(b *testing.B)  { benchaxpyinc(b, 1000, 1, AxpyInc) }
func BenchmarkF64AxpyIncN1000Inc2(b *testing.B)  { benchaxpyinc(b, 1000, 2, AxpyInc) }
func BenchmarkF64AxpyIncN1000Inc4(b *testing.B)  { benchaxpyinc(b, 1000, 4, AxpyInc) }
func BenchmarkF64AxpyIncN1000Inc10(b *testing.B) { benchaxpyinc(b, 1000, 10, AxpyInc) }

func BenchmarkF64AxpyIncN100000Inc1(b *testing.B)  { benchaxpyinc(b, 100000, 1, AxpyInc) }
func BenchmarkF64AxpyIncN100000Inc2(b *testing.B)  { benchaxpyinc(b, 100000, 2, AxpyInc) }
func BenchmarkF64AxpyIncN100000Inc4(b *testing.B)  { benchaxpyinc(b, 100000, 4, AxpyInc) }
func BenchmarkF64AxpyIncN100000Inc10(b *testing.B) { benchaxpyinc(b, 100000, 10, AxpyInc) }

func BenchmarkF64AxpyIncN100000IncM1(b *testing.B)  { benchaxpyinc(b, 100000, -1, AxpyInc) }
func BenchmarkF64AxpyIncN100000IncM2(b *testing.B)  { benchaxpyinc(b, 100000, -2, AxpyInc) }
func BenchmarkF64AxpyIncN100000IncM4(b *testing.B)  { benchaxpyinc(b, 100000, -4, AxpyInc) }
func BenchmarkF64AxpyIncN100000IncM10(b *testing.B) { benchaxpyinc(b, 100000, -10, AxpyInc) }

func BenchmarkLF64AxpyIncN1Inc1(b *testing.B) { benchaxpyinc(b, 1, 1, naiveaxpyinc) }

func BenchmarkLF64AxpyIncN2Inc1(b *testing.B)  { benchaxpyinc(b, 2, 1, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN2Inc2(b *testing.B)  { benchaxpyinc(b, 2, 2, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN2Inc4(b *testing.B)  { benchaxpyinc(b, 2, 4, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN2Inc10(b *testing.B) { benchaxpyinc(b, 2, 10, naiveaxpyinc) }

func BenchmarkLF64AxpyIncN3Inc1(b *testing.B)  { benchaxpyinc(b, 3, 1, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN3Inc2(b *testing.B)  { benchaxpyinc(b, 3, 2, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN3Inc4(b *testing.B)  { benchaxpyinc(b, 3, 4, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN3Inc10(b *testing.B) { benchaxpyinc(b, 3, 10, naiveaxpyinc) }

func BenchmarkLF64AxpyIncN4Inc1(b *testing.B)  { benchaxpyinc(b, 4, 1, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN4Inc2(b *testing.B)  { benchaxpyinc(b, 4, 2, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN4Inc4(b *testing.B)  { benchaxpyinc(b, 4, 4, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN4Inc10(b *testing.B) { benchaxpyinc(b, 4, 10, naiveaxpyinc) }

func BenchmarkLF64AxpyIncN10Inc1(b *testing.B)  { benchaxpyinc(b, 10, 1, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN10Inc2(b *testing.B)  { benchaxpyinc(b, 10, 2, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN10Inc4(b *testing.B)  { benchaxpyinc(b, 10, 4, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN10Inc10(b *testing.B) { benchaxpyinc(b, 10, 10, naiveaxpyinc) }

func BenchmarkLF64AxpyIncN1000Inc1(b *testing.B)  { benchaxpyinc(b, 1000, 1, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN1000Inc2(b *testing.B)  { benchaxpyinc(b, 1000, 2, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN1000Inc4(b *testing.B)  { benchaxpyinc(b, 1000, 4, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN1000Inc10(b *testing.B) { benchaxpyinc(b, 1000, 10, naiveaxpyinc) }

func BenchmarkLF64AxpyIncN100000Inc1(b *testing.B)  { benchaxpyinc(b, 100000, 1, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN100000Inc2(b *testing.B)  { benchaxpyinc(b, 100000, 2, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN100000Inc4(b *testing.B)  { benchaxpyinc(b, 100000, 4, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN100000Inc10(b *testing.B) { benchaxpyinc(b, 100000, 10, naiveaxpyinc) }

func BenchmarkLF64AxpyIncN100000IncM1(b *testing.B)  { benchaxpyinc(b, 100000, -1, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN100000IncM2(b *testing.B)  { benchaxpyinc(b, 100000, -2, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN100000IncM4(b *testing.B)  { benchaxpyinc(b, 100000, -4, naiveaxpyinc) }
func BenchmarkLF64AxpyIncN100000IncM10(b *testing.B) { benchaxpyinc(b, 100000, -10, naiveaxpyinc) }

func benchaxpyincto(t *testing.B, ln, t_inc int, f func(dst []float64, incDst, idst uintptr, alpha float64, x, y []float64, n, incX, incY, ix, iy uintptr)) {
	n, inc := uintptr(ln), uintptr(t_inc)
	var idx int
	if t_inc < 0 {
		idx = (-ln + 1) * t_inc
	}

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		f(z, inc, uintptr(idx), 1, x, y, n, inc, inc, uintptr(idx), uintptr(idx))
	}
	t.StopTimer()
	runtime.GC()
}

func naiveaxpyincto(dst []float64, incDst, idst uintptr, alpha float64, x, y []float64, n, incX, incY, ix, iy uintptr) {
	for i := 0; i < int(n); i++ {
		dst[idst] = alpha*x[ix] + y[iy]
		ix += incX
		iy += incY
		idst += incDst
	}
}

func BenchmarkF64AxpyIncToN1Inc1(b *testing.B) { benchaxpyincto(b, 1, 1, AxpyIncTo) }

func BenchmarkF64AxpyIncToN2Inc1(b *testing.B)  { benchaxpyincto(b, 2, 1, AxpyIncTo) }
func BenchmarkF64AxpyIncToN2Inc2(b *testing.B)  { benchaxpyincto(b, 2, 2, AxpyIncTo) }
func BenchmarkF64AxpyIncToN2Inc4(b *testing.B)  { benchaxpyincto(b, 2, 4, AxpyIncTo) }
func BenchmarkF64AxpyIncToN2Inc10(b *testing.B) { benchaxpyincto(b, 2, 10, AxpyIncTo) }

func BenchmarkF64AxpyIncToN3Inc1(b *testing.B)  { benchaxpyincto(b, 3, 1, AxpyIncTo) }
func BenchmarkF64AxpyIncToN3Inc2(b *testing.B)  { benchaxpyincto(b, 3, 2, AxpyIncTo) }
func BenchmarkF64AxpyIncToN3Inc4(b *testing.B)  { benchaxpyincto(b, 3, 4, AxpyIncTo) }
func BenchmarkF64AxpyIncToN3Inc10(b *testing.B) { benchaxpyincto(b, 3, 10, AxpyIncTo) }

func BenchmarkF64AxpyIncToN4Inc1(b *testing.B)  { benchaxpyincto(b, 4, 1, AxpyIncTo) }
func BenchmarkF64AxpyIncToN4Inc2(b *testing.B)  { benchaxpyincto(b, 4, 2, AxpyIncTo) }
func BenchmarkF64AxpyIncToN4Inc4(b *testing.B)  { benchaxpyincto(b, 4, 4, AxpyIncTo) }
func BenchmarkF64AxpyIncToN4Inc10(b *testing.B) { benchaxpyincto(b, 4, 10, AxpyIncTo) }

func BenchmarkF64AxpyIncToN10Inc1(b *testing.B)  { benchaxpyincto(b, 10, 1, AxpyIncTo) }
func BenchmarkF64AxpyIncToN10Inc2(b *testing.B)  { benchaxpyincto(b, 10, 2, AxpyIncTo) }
func BenchmarkF64AxpyIncToN10Inc4(b *testing.B)  { benchaxpyincto(b, 10, 4, AxpyIncTo) }
func BenchmarkF64AxpyIncToN10Inc10(b *testing.B) { benchaxpyincto(b, 10, 10, AxpyIncTo) }

func BenchmarkF64AxpyIncToN1000Inc1(b *testing.B)  { benchaxpyincto(b, 1000, 1, AxpyIncTo) }
func BenchmarkF64AxpyIncToN1000Inc2(b *testing.B)  { benchaxpyincto(b, 1000, 2, AxpyIncTo) }
func BenchmarkF64AxpyIncToN1000Inc4(b *testing.B)  { benchaxpyincto(b, 1000, 4, AxpyIncTo) }
func BenchmarkF64AxpyIncToN1000Inc10(b *testing.B) { benchaxpyincto(b, 1000, 10, AxpyIncTo) }

func BenchmarkF64AxpyIncToN100000Inc1(b *testing.B)  { benchaxpyincto(b, 100000, 1, AxpyIncTo) }
func BenchmarkF64AxpyIncToN100000Inc2(b *testing.B)  { benchaxpyincto(b, 100000, 2, AxpyIncTo) }
func BenchmarkF64AxpyIncToN100000Inc4(b *testing.B)  { benchaxpyincto(b, 100000, 4, AxpyIncTo) }
func BenchmarkF64AxpyIncToN100000Inc10(b *testing.B) { benchaxpyincto(b, 100000, 10, AxpyIncTo) }

func BenchmarkF64AxpyIncToN100000IncM1(b *testing.B)  { benchaxpyincto(b, 100000, -1, AxpyIncTo) }
func BenchmarkF64AxpyIncToN100000IncM2(b *testing.B)  { benchaxpyincto(b, 100000, -2, AxpyIncTo) }
func BenchmarkF64AxpyIncToN100000IncM4(b *testing.B)  { benchaxpyincto(b, 100000, -4, AxpyIncTo) }
func BenchmarkF64AxpyIncToN100000IncM10(b *testing.B) { benchaxpyincto(b, 100000, -10, AxpyIncTo) }

func BenchmarkLF64AxpyIncToN1Inc1(b *testing.B) { benchaxpyincto(b, 1, 1, naiveaxpyincto) }

func BenchmarkLF64AxpyIncToN2Inc1(b *testing.B)  { benchaxpyincto(b, 2, 1, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN2Inc2(b *testing.B)  { benchaxpyincto(b, 2, 2, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN2Inc4(b *testing.B)  { benchaxpyincto(b, 2, 4, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN2Inc10(b *testing.B) { benchaxpyincto(b, 2, 10, naiveaxpyincto) }

func BenchmarkLF64AxpyIncToN3Inc1(b *testing.B)  { benchaxpyincto(b, 3, 1, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN3Inc2(b *testing.B)  { benchaxpyincto(b, 3, 2, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN3Inc4(b *testing.B)  { benchaxpyincto(b, 3, 4, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN3Inc10(b *testing.B) { benchaxpyincto(b, 3, 10, naiveaxpyincto) }

func BenchmarkLF64AxpyIncToN4Inc1(b *testing.B)  { benchaxpyincto(b, 4, 1, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN4Inc2(b *testing.B)  { benchaxpyincto(b, 4, 2, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN4Inc4(b *testing.B)  { benchaxpyincto(b, 4, 4, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN4Inc10(b *testing.B) { benchaxpyincto(b, 4, 10, naiveaxpyincto) }

func BenchmarkLF64AxpyIncToN10Inc1(b *testing.B)  { benchaxpyincto(b, 10, 1, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN10Inc2(b *testing.B)  { benchaxpyincto(b, 10, 2, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN10Inc4(b *testing.B)  { benchaxpyincto(b, 10, 4, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN10Inc10(b *testing.B) { benchaxpyincto(b, 10, 10, naiveaxpyincto) }

func BenchmarkLF64AxpyIncToN1000Inc1(b *testing.B)  { benchaxpyincto(b, 1000, 1, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN1000Inc2(b *testing.B)  { benchaxpyincto(b, 1000, 2, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN1000Inc4(b *testing.B)  { benchaxpyincto(b, 1000, 4, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN1000Inc10(b *testing.B) { benchaxpyincto(b, 1000, 10, naiveaxpyincto) }

func BenchmarkLF64AxpyIncToN100000Inc1(b *testing.B)  { benchaxpyincto(b, 100000, 1, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN100000Inc2(b *testing.B)  { benchaxpyincto(b, 100000, 2, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN100000Inc4(b *testing.B)  { benchaxpyincto(b, 100000, 4, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN100000Inc10(b *testing.B) { benchaxpyincto(b, 100000, 10, naiveaxpyincto) }

func BenchmarkLF64AxpyIncToN100000IncM1(b *testing.B)  { benchaxpyincto(b, 100000, -1, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN100000IncM2(b *testing.B)  { benchaxpyincto(b, 100000, -2, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN100000IncM4(b *testing.B)  { benchaxpyincto(b, 100000, -4, naiveaxpyincto) }
func BenchmarkLF64AxpyIncToN100000IncM10(b *testing.B) { benchaxpyincto(b, 100000, -10, naiveaxpyincto) }
