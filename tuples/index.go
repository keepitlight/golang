package tuples

// Pair represents a tuple of two values.
//
// 二元组
type Pair[A, B any] struct {
	A A `json:"a,omitempty"`
	B B `json:"b,omitempty"`
}

// Value returns the values of the tuple.
//
// 返回元组的值
func (t *Pair[A, B]) Value() (A, B) {
	return t.A, t.B
}

// Reverse returns a new tuple with the values reversed.
//
// 反转值的顺序并创建新的元组
func (t *Pair[A, B]) Reverse() *Pair[B, A] {
	return &Pair[B, A]{
		A: t.B,
		B: t.A,
	}
}

// Triplet represents a tuple of three values.
//
// 三元组
type Triplet[A, B, C any] struct {
	A A `json:"a,omitempty"`
	B B `json:"b,omitempty"`
	C C `json:"c,omitempty"`
}

// Value returns the values of the tuple.
//
// 返回元组的值
func (t *Triplet[A, B, C]) Value() (A, B, C) {
	return t.A, t.B, t.C
}

// Reverse returns a new tuple with the values reversed.
//
// 反转值的顺序并创建新的元组
func (t *Triplet[A, B, C]) Reverse() *Triplet[C, B, A] {
	return &Triplet[C, B, A]{
		A: t.C,
		B: t.B,
		C: t.A,
	}
}

// Leading returns a new tuple with the leading values.
//
// 将前半部分的值组成新的元组
func (t *Triplet[A, B, C]) Leading() *Pair[A, B] {
	return &Pair[A, B]{
		A: t.A,
		B: t.B,
	}
}

// Trailing returns a new tuple with the trailing values.
//
// 将后半部分的值组成新的元组
func (t *Triplet[A, B, C]) Trailing() *Pair[B, C] {
	return &Pair[B, C]{
		A: t.B,
		B: t.C,
	}
}

// Sides returns a new tuple with the head and tail values.
//
// 将两端的值组成新的元组
func (t *Triplet[A, B, C]) Sides() *Pair[A, C] {
	return &Pair[A, C]{
		A: t.A,
		B: t.C,
	}
}

func (t *Triplet[A, B, C]) AB() *Pair[A, B] {
	return &Pair[A, B]{
		A: t.A,
		B: t.B,
	}
}
func (t *Triplet[A, B, C]) BC() *Pair[B, C] {
	return &Pair[B, C]{
		A: t.B,
		B: t.C,
	}
}
func (t *Triplet[A, B, C]) AC() *Pair[A, C] {
	return &Pair[A, C]{
		A: t.A,
		B: t.C,
	}
}

func (t *Triplet[A, B, C]) BA() *Pair[B, A] {
	return &Pair[B, A]{
		A: t.B,
		B: t.A,
	}
}
func (t *Triplet[A, B, C]) CB() *Pair[C, B] {
	return &Pair[C, B]{
		A: t.C,
		B: t.B,
	}
}
func (t *Triplet[A, B, C]) CA() *Pair[C, A] {
	return &Pair[C, A]{
		A: t.C,
		B: t.A,
	}
}

// Quad (quadruplet) represents a tuple of four values.
//
// 四元组
type Quad[A, B, C, D any] struct {
	A A `json:"a,omitempty"`
	B B `json:"b,omitempty"`
	C C `json:"c,omitempty"`
	D D `json:"d,omitempty"`
}

// Value returns the values of the tuple.
//
// 返回元组的值
func (t *Quad[A, B, C, D]) Value() (A, B, C, D) {
	return t.A, t.B, t.C, t.D
}

// Reverse returns a new tuple with the values reversed.
//
// 反转值的顺序并创建新的元组
func (t *Quad[A, B, C, D]) Reverse() *Quad[D, C, B, A] {
	return &Quad[D, C, B, A]{
		A: t.D,
		B: t.C,
		C: t.B,
		D: t.A,
	}
}

// Leading returns a new tuple with the leading values.
//
// 将前半部分的值组成新的元组
func (t *Quad[A, B, C, D]) Leading() *Triplet[A, B, C] {
	return &Triplet[A, B, C]{
		A: t.A,
		B: t.B,
		C: t.C,
	}
}

// Trailing returns a new tuple with the trailing values.
//
// 将后半部分的值组成新的元组
func (t *Quad[A, B, C, D]) Trailing() *Triplet[B, C, D] {
	return &Triplet[B, C, D]{
		A: t.B,
		B: t.C,
		C: t.D,
	}
}

// Sides returns a new tuple with the head and tail values.
//
// 将两端的值组成新的元组
func (t *Quad[A, B, C, D]) Sides() *Pair[A, D] {
	return &Pair[A, D]{
		A: t.A,
		B: t.D,
	}
}

// Inners return a new tuple with the inner values.
//
// 将中间部分的值组成新的元组
func (t *Quad[A, B, C, D]) Inners() *Pair[B, C] {
	return &Pair[B, C]{
		A: t.B,
		B: t.C,
	}
}

func (t *Quad[A, B, C, D]) AB() *Pair[A, B] {
	return &Pair[A, B]{
		A: t.A,
		B: t.B,
	}
}
func (t *Quad[A, B, C, D]) BC() *Pair[B, C] {
	return &Pair[B, C]{
		A: t.B,
		B: t.C,
	}
}
func (t *Quad[A, B, C, D]) CD() *Pair[C, D] {
	return &Pair[C, D]{
		A: t.C,
		B: t.D,
	}
}

func (t *Quad[A, B, C, D]) AC() *Pair[A, C] {
	return &Pair[A, C]{
		A: t.A,
		B: t.C,
	}
}
func (t *Quad[A, B, C, D]) BD() *Pair[B, D] {
	return &Pair[B, D]{
		A: t.B,
		B: t.D,
	}
}
func (t *Quad[A, B, C, D]) AD() *Pair[A, D] {
	return &Pair[A, D]{
		A: t.A,
		B: t.D,
	}
}

// Quint (quintuplet) represents a tuple of five values.
//
// 五元组
type Quint[A, B, C, D, E any] struct {
	A A `json:"a,omitempty"`
	B B `json:"b,omitempty"`
	C C `json:"c,omitempty"`
	D D `json:"d,omitempty"`
	E E `json:"e,omitempty"`
}

// Value returns the values of the tuple.
//
// 返回元组的值
func (t *Quint[A, B, C, D, E]) Value() (A, B, C, D, E) {
	return t.A, t.B, t.C, t.D, t.E
}

// Reverse returns a new tuple with the values reversed.
//
// 反转值的顺序并创建新的元组
func (t *Quint[A, B, C, D, E]) Reverse() *Quint[E, D, C, B, A] {
	return &Quint[E, D, C, B, A]{
		A: t.E,
		B: t.D,
		C: t.C,
		D: t.B,
		E: t.A,
	}
}

// Leading returns a new tuple with the leading values.
//
// 将前半部分的值组成新的元组
func (t *Quint[A, B, C, D, E]) Leading() *Triplet[A, B, C] {
	return &Triplet[A, B, C]{
		A: t.A,
		B: t.B,
		C: t.C,
	}
}

// Trailing returns a new tuple with the trailing values.
//
// 将后半部分的值组成新的元组
func (t *Quint[A, B, C, D, E]) Trailing() *Triplet[C, D, E] {
	return &Triplet[C, D, E]{
		A: t.C,
		B: t.D,
		C: t.E,
	}
}

// Inners return a new tuple with the inner values.
//
// 将中间部分的值组成新的元组
func (t *Quint[A, B, C, D, E]) Inners() *Triplet[B, C, D] {
	return &Triplet[B, C, D]{
		A: t.B,
		B: t.C,
		C: t.D,
	}
}

// Sides returns a new tuple with the head and tail values.
//
// 将两端的值组成新的元组
func (t *Quint[A, B, C, D, E]) Sides() *Pair[A, E] {
	return &Pair[A, E]{
		A: t.A,
		B: t.E,
	}
}

func (t *Quint[A, B, C, D, E]) AB() *Pair[A, B] {
	return &Pair[A, B]{
		A: t.A,
		B: t.B,
	}
}
func (t *Quint[A, B, C, D, E]) BC() *Pair[B, C] {
	return &Pair[B, C]{
		A: t.B,
		B: t.C,
	}
}
func (t *Quint[A, B, C, D, E]) CD() *Pair[C, D] {
	return &Pair[C, D]{
		A: t.C,
		B: t.D,
	}
}
func (t *Quint[A, B, C, D, E]) DE() *Pair[D, E] {
	return &Pair[D, E]{
		A: t.D,
		B: t.E,
	}
}

func (t *Quint[A, B, C, D, E]) ABC() *Triplet[A, B, C] {
	return &Triplet[A, B, C]{
		A: t.A,
		B: t.B,
		C: t.C,
	}
}
func (t *Quint[A, B, C, D, E]) BCD() *Triplet[B, C, D] {
	return &Triplet[B, C, D]{
		A: t.B,
		B: t.C,
		C: t.D,
	}
}
func (t *Quint[A, B, C, D, E]) CDE() *Triplet[C, D, E] {
	return &Triplet[C, D, E]{
		A: t.C,
		B: t.D,
		C: t.E,
	}
}

// Sextet (sextuplet) represents a tuple of six values.
//
// 六元组
type Sextet[A, B, C, D, E, F any] struct {
	A A `json:"a,omitempty"`
	B B `json:"b,omitempty"`
	C C `json:"c,omitempty"`
	D D `json:"d,omitempty"`
	E E `json:"e,omitempty"`
	F F `json:"f,omitempty"`
}

// Value returns the values of the tuple.
//
// 返回元组的值
func (t *Sextet[A, B, C, D, E, F]) Value() (A, B, C, D, E, F) {
	return t.A, t.B, t.C, t.D, t.E, t.F
}

// Reverse returns a new tuple with the values reversed.
//
// 反转值的顺序并创建新的元组
func (t *Sextet[A, B, C, D, E, F]) Reverse() *Sextet[F, E, D, C, B, A] {
	return &Sextet[F, E, D, C, B, A]{
		A: t.F,
		B: t.E,
		C: t.D,
		D: t.C,
		E: t.B,
		F: t.A,
	}
}

func (t *Sextet[A, B, C, D, E, F]) ABC() *Triplet[A, B, C] {
	return &Triplet[A, B, C]{
		A: t.A,
		B: t.B,
		C: t.C,
	}
}
func (t *Sextet[A, B, C, D, E, F]) BCD() *Triplet[B, C, D] {
	return &Triplet[B, C, D]{
		A: t.B,
		B: t.C,
		C: t.D,
	}
}
func (t *Sextet[A, B, C, D, E, F]) CDE() *Triplet[C, D, E] {
	return &Triplet[C, D, E]{
		A: t.C,
		B: t.D,
		C: t.E,
	}
}
func (t *Sextet[A, B, C, D, E, F]) DEF() *Triplet[D, E, F] {
	return &Triplet[D, E, F]{
		A: t.D,
		B: t.E,
		C: t.F,
	}
}

// Leading returns a new tuple with the leading values.
//
// 将前半部分的值组成新的元组
func (t *Sextet[A, B, C, D, E, F]) Leading() *Quad[A, B, C, D] {
	return &Quad[A, B, C, D]{
		A: t.A,
		B: t.B,
		C: t.C,
		D: t.D,
	}
}

// Trailing returns a new tuple with the trailing values.
//
// 将后半部分的值组成新的元组
func (t *Sextet[A, B, C, D, E, F]) Trailing() *Quad[C, D, E, F] {
	return &Quad[C, D, E, F]{
		A: t.C,
		B: t.D,
		C: t.E,
		D: t.F,
	}
}

// Inners return a new tuple with the inner values.
//
// 将中间部分的值组成新的元组
func (t *Sextet[A, B, C, D, E, F]) Inners() *Quad[B, C, D, E] {
	return &Quad[B, C, D, E]{
		A: t.B,
		B: t.C,
		C: t.D,
		D: t.E,
	}
}

// Sides returns a new tuple with the head and tail values.
//
// 将两端的值组成新的元组
func (t *Sextet[A, B, C, D, E, F]) Sides() *Pair[A, F] {
	return &Pair[A, F]{
		A: t.A,
		B: t.F,
	}
}

// Sept (septuplet) represents a tuple of seven values.
//
// 七元组
type Sept[A, B, C, D, E, F, G any] struct {
	A A `json:"a,omitempty"`
	B B `json:"b,omitempty"`
	C C `json:"c,omitempty"`
	D D `json:"d,omitempty"`
	E E `json:"e,omitempty"`
	F F `json:"f,omitempty"`
	G G `json:"g,omitempty"`
}

// Value returns the values of the tuple.
//
// 返回元组的值
func (t *Sept[A, B, C, D, E, F, G]) Value() (A, B, C, D, E, F, G) {
	return t.A, t.B, t.C, t.D, t.E, t.F, t.G
}

// Reverse returns a new tuple with the values reversed.
//
// 反转值的顺序并创建新的元组
func (t *Sept[A, B, C, D, E, F, G]) Reverse() *Sept[G, F, E, D, C, B, A] {
	return &Sept[G, F, E, D, C, B, A]{
		A: t.G,
		B: t.F,
		C: t.E,
		D: t.D,
		E: t.C,
		F: t.B,
		G: t.A,
	}
}

// Leading returns a new tuple with the leading values.
//
// 将前半部分的值组成新的元组
func (t *Sept[A, B, C, D, E, F, G]) Leading() *Quad[A, B, C, D] {
	return &Quad[A, B, C, D]{
		A: t.A,
		B: t.B,
		C: t.C,
		D: t.D,
	}
}

// Trailing returns a new tuple with the trailing values.
//
// 将后半部分的值组成新的元组
func (t *Sept[A, B, C, D, E, F, G]) Trailing() *Quad[D, E, F, G] {
	return &Quad[D, E, F, G]{
		A: t.D,
		B: t.E,
		C: t.F,
		D: t.G,
	}
}

// Inners return a new tuple with the inner values.
//
// 将中间部分的值组成新的元组
func (t *Sept[A, B, C, D, E, F, G]) Inners() *Quint[B, C, D, E, F] {
	return &Quint[B, C, D, E, F]{
		A: t.B,
		B: t.C,
		C: t.D,
		D: t.E,
		E: t.F,
	}
}

// Sides returns a new tuple with the head and tail values.
//
// 将两端的值组成新的元组
func (t *Sept[A, B, C, D, E, F, G]) Sides() *Pair[A, G] {
	return &Pair[A, G]{
		A: t.A,
		B: t.G,
	}
}

// Oct (octuplet) represents a tuple of eight values.
//
// 八元组
type Oct[A, B, C, D, E, F, G, H any] struct {
	A A `json:"a,omitempty"`
	B B `json:"b,omitempty"`
	C C `json:"c,omitempty"`
	D D `json:"d,omitempty"`
	E E `json:"e,omitempty"`
	F F `json:"f,omitempty"`
	G G `json:"g,omitempty"`
	H H `json:"h,omitempty"`
}

// Value returns the values of the tuple.
//
// 返回元组的值
func (t *Oct[A, B, C, D, E, F, G, H]) Value() (A, B, C, D, E, F, G, H) {
	return t.A, t.B, t.C, t.D, t.E, t.F, t.G, t.H
}

// Reverse returns a new tuple with the values reversed.
//
// 反转值的顺序并创建新的元组
func (t *Oct[A, B, C, D, E, F, G, H]) Reverse() *Oct[H, G, F, E, D, C, B, A] {
	return &Oct[H, G, F, E, D, C, B, A]{
		A: t.H,
		B: t.G,
		C: t.F,
		D: t.E,
		E: t.D,
		F: t.C,
		G: t.B,
		H: t.A,
	}
}

// Leading returns a new tuple with the leading values.
//
// 将前半部分的值组成新的元组
func (t *Oct[A, B, C, D, E, F, G, H]) Leading() *Quint[A, B, C, D, E] {
	return &Quint[A, B, C, D, E]{
		A: t.A,
		B: t.B,
		C: t.C,
		D: t.D,
		E: t.E,
	}
}

// Trailing returns a new tuple with the trailing values.
//
// 将后半部分的值组成新的元组
func (t *Oct[A, B, C, D, E, F, G, H]) Trailing() *Quint[D, E, F, G, H] {
	return &Quint[D, E, F, G, H]{
		A: t.D,
		B: t.E,
		C: t.F,
		D: t.G,
		E: t.H,
	}
}

// Inners return a new tuple with the inner values.
//
// 将中间部分的值组成新的元组
func (t *Oct[A, B, C, D, E, F, G, H]) Inners() *Sextet[B, C, D, E, F, G] {
	return &Sextet[B, C, D, E, F, G]{
		A: t.B,
		B: t.C,
		C: t.D,
		D: t.E,
		E: t.F,
		F: t.G,
	}
}

// Sides returns a new tuple with the head and tail values.
//
// 将两端的值组成新的元组
func (t *Oct[A, B, C, D, E, F, G, H]) Sides() *Pair[A, H] {
	return &Pair[A, H]{
		A: t.A,
		B: t.H,
	}
}

// Nonet (nonuplet) represents a tuple of nine values.
//
// 九元组
type Nonet[A, B, C, D, E, F, G, H, I any] struct {
	A A `json:"a,omitempty"`
	B B `json:"b,omitempty"`
	C C `json:"c,omitempty"`
	D D `json:"d,omitempty"`
	E E `json:"e,omitempty"`
	F F `json:"f,omitempty"`
	G G `json:"g,omitempty"`
	H H `json:"h,omitempty"`
	I I `json:"i,omitempty"`
}

// Value returns the values of the tuple.
//
// 返回元组的值
func (t *Nonet[A, B, C, D, E, F, G, H, I]) Value() (A, B, C, D, E, F, G, H, I) {
	return t.A, t.B, t.C, t.D, t.E, t.F, t.G, t.H, t.I
}

// Reverse returns a new tuple with the values reversed.
//
// 反转值的顺序并创建新的元组
func (t *Nonet[A, B, C, D, E, F, G, H, I]) Reverse() *Nonet[I, H, G, F, E, D, C, B, A] {
	return &Nonet[I, H, G, F, E, D, C, B, A]{
		A: t.I,
		B: t.H,
		C: t.G,
		D: t.F,
		E: t.E,
		F: t.D,
		G: t.C,
		H: t.B,
		I: t.A,
	}
}

// Leading returns a new tuple with the leading values.
//
// 将前半部分的值组成新的元组
func (t *Nonet[A, B, C, D, E, F, G, H, I]) Leading() *Quint[A, B, C, D, E] {
	return &Quint[A, B, C, D, E]{
		A: t.A,
		B: t.B,
		C: t.C,
		D: t.D,
		E: t.E,
	}
}

// Trailing returns a new tuple with the trailing values.
//
// 将后半部分的值组成新的元组
func (t *Nonet[A, B, C, D, E, F, G, H, I]) Trailing() *Quint[E, F, G, H, I] {
	return &Quint[E, F, G, H, I]{
		A: t.E,
		B: t.F,
		C: t.G,
		D: t.H,
		E: t.I,
	}
}

// Inners return a new tuple with the inner values.
//
// 将中间部分的值组成新的元组
func (t *Nonet[A, B, C, D, E, F, G, H, I]) Inners() *Sept[B, C, D, E, F, G, H] {
	return &Sept[B, C, D, E, F, G, H]{
		A: t.B,
		B: t.C,
		C: t.D,
		D: t.E,
		E: t.F,
		F: t.G,
		G: t.H,
	}
}

// Sides returns a new tuple with the head and tail values.
//
// 将两端的值组成新的元组
func (t *Nonet[A, B, C, D, E, F, G, H, I]) Sides() *Pair[A, I] {
	return &Pair[A, I]{
		A: t.A,
		B: t.I,
	}
}
