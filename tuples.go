package golang

import "github.com/keepitlight/golang/tuples"

// Tuple creates a new Pair.
//
// 创建二元组
func Tuple[A, B any](a A, b B) *tuples.Pair[A, B] {
	return &tuples.Pair[A, B]{A: a, B: b}
}

// Tuple3 creates a new Triplet.
//
// 创建三元组
func Tuple3[A, B, C any](a A, b B, c C) *tuples.Triplet[A, B, C] {
	return &tuples.Triplet[A, B, C]{A: a, B: b, C: c}
}

// Tuple4 creates a new Quad(quadruplet).
//
// 创建四元组
func Tuple4[A, B, C, D any](a A, b B, c C, d D) *tuples.Quad[A, B, C, D] {
	return &tuples.Quad[A, B, C, D]{A: a, B: b, C: c, D: d}
}

// Tuple5 creates a new Quint(quintuplet).
//
// 创建五元组
func Tuple5[A, B, C, D, E any](a A, b B, c C, d D, e E) *tuples.Quint[A, B, C, D, E] {
	return &tuples.Quint[A, B, C, D, E]{A: a, B: b, C: c, D: d, E: e}
}

// Tuple6 creates a new Sextet(sextuplet).
//
// 创建六元组
func Tuple6[A, B, C, D, E, F any](a A, b B, c C, d D, e E, f F) *tuples.Sextet[A, B, C, D, E, F] {
	return &tuples.Sextet[A, B, C, D, E, F]{A: a, B: b, C: c, D: d, E: e, F: f}
}

// Tuple7 creates a new Sept(septuplet).
//
// 创建七元组
func Tuple7[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, f F, g G) *tuples.Sept[A, B, C, D, E, F, G] {
	return &tuples.Sept[A, B, C, D, E, F, G]{A: a, B: b, C: c, D: d, E: e, F: f, G: g}
}

// Tuple8 creates a new Oct(octuplet).
//
// 创建八元组
func Tuple8[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, g G, h H) *tuples.Oct[A, B, C, D, E, F, G, H] {
	return &tuples.Oct[A, B, C, D, E, F, G, H]{A: a, B: b, C: c, D: d, E: e, F: f, G: g, H: h}
}

// Tuple9 creates a new Nonet(nonuplet).
//
// 创建九元组
func Tuple9[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I) *tuples.Nonet[A, B, C, D, E, F, G, H, I] {
	return &tuples.Nonet[A, B, C, D, E, F, G, H, I]{A: a, B: b, C: c, D: d, E: e, F: f, G: g, H: h, I: i}
}

// T3 creates a new Triplet.
//
// 创建三元组
func T3[A, B, C any](a A, b B, c C) *tuples.Triplet[A, B, C] {
	return &tuples.Triplet[A, B, C]{A: a, B: b, C: c}
}

// T4 creates a new Quad(quadruplet).
//
// 创建四元组
func T4[A, B, C, D any](a A, b B, c C, d D) *tuples.Quad[A, B, C, D] {
	return &tuples.Quad[A, B, C, D]{A: a, B: b, C: c, D: d}
}

// T5 creates a new Quint(quintuplet).
//
// 创建五元组
func T5[A, B, C, D, E any](a A, b B, c C, d D, e E) *tuples.Quint[A, B, C, D, E] {
	return &tuples.Quint[A, B, C, D, E]{A: a, B: b, C: c, D: d, E: e}
}

// T6 creates a new Sextet(sextuplet).
//
// 创建六元组
func T6[A, B, C, D, E, F any](a A, b B, c C, d D, e E, f F) *tuples.Sextet[A, B, C, D, E, F] {
	return &tuples.Sextet[A, B, C, D, E, F]{A: a, B: b, C: c, D: d, E: e, F: f}
}

// T7 creates a new Sept(septuplet).
//
// 创建七元组
func T7[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, f F, g G) *tuples.Sept[A, B, C, D, E, F, G] {
	return &tuples.Sept[A, B, C, D, E, F, G]{A: a, B: b, C: c, D: d, E: e, F: f, G: g}
}

// T8 creates a new Oct(octuplet).
//
// 创建八元组
func T8[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, g G, h H) *tuples.Oct[A, B, C, D, E, F, G, H] {
	return &tuples.Oct[A, B, C, D, E, F, G, H]{A: a, B: b, C: c, D: d, E: e, F: f, G: g, H: h}
}

// T9 creates a new Nonet(nonuplet).
//
// 创建九元组
func T9[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I) *tuples.Nonet[A, B, C, D, E, F, G, H, I] {
	return &tuples.Nonet[A, B, C, D, E, F, G, H, I]{A: a, B: b, C: c, D: d, E: e, F: f, G: g, H: h, I: i}
}

// Pair creates a pair.
//
// 创建二元组
func Pair[A, B any](a A, b B) *tuples.Pair[A, B] {
	return &tuples.Pair[A, B]{A: a, B: b}
}

// Triplet creates a new triplet.
//
// 创建三元组
func Triplet[A, B, C any](a A, b B, c C) *tuples.Triplet[A, B, C] {
	return &tuples.Triplet[A, B, C]{A: a, B: b, C: c}
}

// Quad creates a new quadruplet.
//
// 创建四元组
func Quad[A, B, C, D any](a A, b B, c C, d D) *tuples.Quad[A, B, C, D] {
	return &tuples.Quad[A, B, C, D]{A: a, B: b, C: c, D: d}
}

// Quint creates a new quintuplet.
//
// 创建五元组
func Quint[A, B, C, D, E any](a A, b B, c C, d D, e E) *tuples.Quint[A, B, C, D, E] {
	return &tuples.Quint[A, B, C, D, E]{A: a, B: b, C: c, D: d, E: e}
}

// Sextet creates a new sextuplet.
//
// 创建六元组
func Sextet[A, B, C, D, E, F any](a A, b B, c C, d D, e E, f F) *tuples.Sextet[A, B, C, D, E, F] {
	return &tuples.Sextet[A, B, C, D, E, F]{A: a, B: b, C: c, D: d, E: e, F: f}
}

// Sept creates a new septuplet.
//
// 创建七元组
func Sept[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, f F, g G) *tuples.Sept[A, B, C, D, E, F, G] {
	return &tuples.Sept[A, B, C, D, E, F, G]{A: a, B: b, C: c, D: d, E: e, F: f, G: g}
}

// Oct creates a new octuplet.
//
// 创建八元组
func Oct[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, g G, h H) *tuples.Oct[A, B, C, D, E, F, G, H] {
	return &tuples.Oct[A, B, C, D, E, F, G, H]{A: a, B: b, C: c, D: d, E: e, F: f, G: g, H: h}
}

// Nonet creates a new nonuplet.
//
// 创建九元组
func Nonet[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I) *tuples.Nonet[A, B, C, D, E, F, G, H, I] {
	return &tuples.Nonet[A, B, C, D, E, F, G, H, I]{A: a, B: b, C: c, D: d, E: e, F: f, G: g, H: h, I: i}
}
