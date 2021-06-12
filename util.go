package elti

const maxVarintBytes = 10

//将数据length使用varint编码(小端序排列)，并将编码结果添加入切片ptr，返回新的切片。
func seriLength(length uint32, ptr []byte) []byte {
	var buf [maxVarintBytes]byte
	var n uint32
	for n = 0; length > 127; n++ {
		buf[n] = 0x80 | uint8(length&0x7F)
		length >>= 7
	}
	buf[n] = uint8(length)
	n++
	ptr = append(ptr, buf[0:n]...)
	return ptr
}

// 从ptr中解码varint编码的数据，返回解码结果和新的解码起始位置。
func parseLength(ptr []byte, begin uint32) (uint32, uint32) {
	var result uint32 = 0
	var n uint32 = 0
	for {
		if n > 4 {
			panic("parseLength error, the num is too large to represent in a 32-bit value.")
		}
		b := ptr[begin+n]
		result |= (uint32)(b&0x7F) << (7 * n)
		n++
		if (b & 0x80) == 0 {
			break
		}
	}
	return result, begin + n
}
