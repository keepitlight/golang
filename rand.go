package golang

import (
	srand "crypto/rand"
	"encoding/binary"
	"math/rand/v2"
	"sync"
	"time"
	"unsafe"
)

var (
	pcgStateSeed  uint64
	chaSeed       [32]byte
	refreshLocker sync.Mutex
	latestRefresh *time.Time
)

// RefreshSeed to refresh the random number generator.
//
// 刷新随机数生成器种子。
func RefreshSeed() {
	refreshLocker.Lock()
	defer refreshLocker.Unlock()

	b1 := make([]byte, 8)
	_, _ = srand.Read(b1)
	pcgStateSeed = binary.LittleEndian.Uint64(b1) // 使用安全随机数初始化 PCG 随机数生成器

	b2 := make([]byte, 32)
	_, _ = srand.Read(b2)

	r := rand.New(rand.NewPCG(pcgStateSeed, uint64(time.Now().UnixNano())))
	r.Shuffle(len(b2), func(i, j int) { b2[i], b2[j] = b2[i], b2[j] })
	chaSeed = *(*[32]byte)(unsafe.Pointer(&b2)) // 使用安全随机数初始化 ChaCha8 随机数生成器

	now := time.Now()
	latestRefresh = &now
}

func init() {
	RefreshSeed()
}

type RandomSourceType int // 随机源类型

const (
	// PCG is a family of simple fast space-efficient statistically good algorithms for random number generation.
	//
	// 来源：PCG 是一类随机数生成器的统称，它基于线性同余方法，并通过使用非线性操作（如XOR和旋转）来提高随机性。
	// PCG 的设计目标是在提供良好随机性的前提下，达到高性能。
	//
	// 特性：PCG 生成器具有周期长、状态空间大、可重复性和可跳转性等优点，这意味着它可以生成大量的随机数序列，
	// 而且这些序列之间不易出现重复。
	//
	// 用途：PCG 适合于需要高吞吐量和低延迟的随机数生成场景，尤其是在多线程环境中，因为 PCG 支持并行生成独立的随机数流。
	PCG RandomSourceType = iota
	// ChaCha8 is a fast, simple, and secure stream cipher.
	//
	// 来源：ChaCha8 是基于 ChaCha 流密码的变体，它被设计用于加密和随机数生成。ChaCha8 使用较少的轮次（8轮），因此与标准 ChaCha20 相比，
	// 它的速度更快，但安全级别较低。
	//
	// 安全性：尽管 ChaCha8 不如 ChaCha20 安全，但它仍然提供了足够的随机性，适用于大多数非加密应用。由于它源自加密算法，
	// ChaCha8 提供了良好的统计特性。
	//
	// 用途：ChaCha8 适合于需要快速生成高质量随机数的场景，例如游戏、模拟和一般用途的随机化需求。
	ChaCha8
	// Crypto is a cryptographically secure random number generator.
	//
	// 来源：Crypto 是一个加密安全的随机数生成器。
	Crypto
)

type CryptoSource struct{}

func (c *CryptoSource) Uint64() uint64 {
	buf := make([]byte, 8)
	n, e := srand.Read(buf)
	if e != nil {
		return 0
	}
	return binary.LittleEndian.Uint64(buf[:n])
}

type randOptions struct {
	pcgIncrementSeed    uint64
	pcgStateSeed        uint64
	refreshSeedDuration time.Duration
	lower, upper        *int64
	t                   RandomSourceType
}

func (r *randOptions) refreshSeed() {
	if r.refreshSeedDuration < 1 {
		return
	}
	if latestRefresh == nil || latestRefresh.Add(r.refreshSeedDuration).Before(time.Now()) {
		RefreshSeed()
	}
}

type RandomOption func(r *randOptions)

func RefreshSeedDuration(d time.Duration) RandomOption {
	return func(r *randOptions) {
		r.refreshSeedDuration = d
	}
}

func CryptoRand() RandomOption {
	return func(r *randOptions) {
		r.t = Crypto
	}
}
func UsePCG(incrementSeed uint64) RandomOption {
	return func(r *randOptions) {
		r.t = PCG
		r.pcgIncrementSeed = incrementSeed
	}
}
func UseChaCha8() RandomOption {
	return func(r *randOptions) {
		r.t = ChaCha8
	}
}

// Rand to generate random bytes corresponding to the options.
//
// 根据选项生成随机字节。
func Rand(buf []byte, options ...RandomOption) {
	random(buf, getOptions(options...))
}

// random to generate random bytes.
func random(buf []byte, opt *randOptions) {
	switch opt.t {
	case ChaCha8:
		r := rand.New(rand.NewChaCha8(chaSeed))
		for i := 0; i < len(buf); i += 8 {
			b := LittleEndianBytes(r.Uint64())
			for j := 0; j < 8; j++ {
				if i+j >= len(buf) {
					return
				}
				buf[i+j] = b[j]
			}
		}
	case PCG:
		r := rand.New(rand.NewPCG(pcgStateSeed, opt.pcgIncrementSeed))
		for i := 0; i < len(buf); i += 8 {
			b := LittleEndianBytes(r.Uint64())
			for j := 0; j < 8; j++ {
				if i+j >= len(buf) {
					return
				}
				buf[i+j] = b[j]
			}
		}
	default:
		_, _ = srand.Read(buf)
	}
}

func getShuffler(opt *randOptions) func(n int, swap func(i, j int)) {
	if opt.t == ChaCha8 {
		return rand.New(rand.NewChaCha8(chaSeed)).Shuffle
	} else if opt.t == PCG {
		if opt.pcgIncrementSeed == 0 {
			opt.pcgIncrementSeed = uint64(time.Now().UnixNano())
		}
		return rand.New(rand.NewPCG(pcgStateSeed, opt.pcgIncrementSeed)).Shuffle
	}
	return func(n int, swap func(i, j int)) {
		if n <= 1 {
			return
		}
		buffer := make([]byte, n*8)
		if _, err := srand.Read(buffer); err != nil {
			return
		}
		for i := 0; i < n; i++ {
			j := i * 8
			r := binary.LittleEndian.Uint64(buffer[j : j+8])
			r %= uint64(n)
			x := int(r)
			if x != i {
				swap(i, x)
			}
		}
	}
}

func getOptions(options ...RandomOption) *randOptions {
	opt := randOptions{}
	for _, o := range options {
		o(&opt)
	}
	opt.refreshSeed()
	if opt.pcgIncrementSeed == 0 {
		opt.pcgIncrementSeed = uint64(time.Now().UnixNano())
	}
	return &opt
}

// Shuffle the order of elements.
// length is the number of elements. Shuffle panics if length < 0.
// swap swaps the elements with indexes i and j.
//
// 乱序排序， length 为元素长度，swap 为交换元素函数, 当 length < 0 时 Shuffle 函数会抛出异常
func Shuffle(length int, swap func(i, j int), options ...RandomOption) {
	getShuffler(getOptions(options...))(length, swap)
}

// ShuffleSlice to shuffle the slice using a cryptographically pseudo-random number generator.
// Argument seeds is optional, used to set the increment seed of PCG random number generator.
//
// 使用随机数，随机打乱切片。
func ShuffleSlice[E any](ss []E, options ...RandomOption) {
	l := len(ss)
	if l <= 1 {
		return
	}
	getShuffler(getOptions(options...))(len(ss), func(i, j int) {
		ss[i], ss[j] = ss[j], ss[i]
	})
	return
}

// ShuffledSlice return a new slice using a random number generator.
// Argument seeds is optional, used to set the increment seed of PCG random number generator.
//
// 返回一个新的使用随机数随机打乱的切片。seeds 为可选参数，用于设置 PCG 随机数生成器的增量种子，其它随机数类型不需要。
func ShuffledSlice[E any](ss []E, options ...RandomOption) (result []E) {
	l := len(ss)
	if l <= 1 {
		return ss
	}
	result = append(result, ss...)

	getShuffler(getOptions(options...))(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return
}
