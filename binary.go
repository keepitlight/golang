package golang

// BigEndianBytes to convert uint64 to bytes corresponding to the big-endian byte order.
//
// 根据大端字节序将 uint64 转换为字节数组。
func BigEndianBytes(v uint64) [8]byte {
	return [8]byte{
		byte(v >> 56),
		byte(v >> 48),
		byte(v >> 40),
		byte(v >> 32),
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v),
	}
}

// LittleEndianBytes to convert uint64 to bytes corresponding to the little-endian byte order.
//
// 根据小端字节序将 uint64 转换为字节数组。
func LittleEndianBytes(v uint64) [8]byte {
	return [8]byte{
		byte(v),
		byte(v >> 8),
		byte(v >> 16),
		byte(v >> 24),
		byte(v >> 32),
		byte(v >> 40),
		byte(v >> 48),
		byte(v >> 56),
	}
}
