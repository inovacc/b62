package base62

import (
	"errors"
)

var encodeTable = [62]byte{
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
	'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f',
	'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
	'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
}

var decodeTable = [128]int8{}

func init() {
	for i := range decodeTable {
		decodeTable[i] = -1
	}
	for i, b := range encodeTable {
		decodeTable[b] = int8(i)
	}
}

const (
	compactMask = 0x1E
	mask5Bits   = 0x1F
)

func Encode(data []byte) string {
	bs := newBitInputStream(data)
	result := make([]byte, 0, len(data)*8/5+1)

	for bs.HasMore() {
		rawBits := bs.ReadBits(6)

		var bits int
		if rawBits&compactMask == compactMask {
			bits = rawBits & mask5Bits
			bs.SeekBit(-1)
		} else {
			bits = rawBits
		}

		result = append(result, encodeTable[bits])
	}

	return string(result)
}

func Decode(input string) ([]byte, error) {
	bs := newBitOutputStream(len(input) * 6)

	for i := 0; i < len(input); i++ {
		c := input[i]
		if c >= 128 || decodeTable[c] == -1 {
			return nil, errors.New("invalid Base62 character: " + string(c))
		}

		bits := int(decodeTable[c])
		var bitsCount int

		if bits&compactMask == compactMask {
			bitsCount = 5
		} else if i == len(input)-1 {
			bitsCount = bs.BitsToNextByte()
		} else {
			bitsCount = 6
		}

		bs.WriteBits(bitsCount, bits)
	}

	return bs.Bytes(), nil
}

type bitInputStream struct {
	buf    []byte
	offset int
}

func newBitInputStream(data []byte) *bitInputStream {
	return &bitInputStream{buf: data}
}

func (b *bitInputStream) SeekBit(pos int) {
	b.offset += pos
}

func (b *bitInputStream) ReadBits(bitsCount int) int {
	bitNum := b.offset % 8
	byteNum := b.offset / 8

	firstRead := min(8-bitNum, bitsCount)
	secondRead := bitsCount - firstRead

	result := int((b.buf[byteNum] >> bitNum) & ((1 << firstRead) - 1))
	if secondRead > 0 && byteNum+1 < len(b.buf) {
		result |= int(b.buf[byteNum+1]&((1<<secondRead)-1)) << firstRead
	}

	b.offset += bitsCount
	return result
}

func (b *bitInputStream) HasMore() bool {
	return b.offset < len(b.buf)*8
}

type bitOutputStream struct {
	buf    []byte
	offset int
}

func newBitOutputStream(capacity int) *bitOutputStream {
	return &bitOutputStream{buf: make([]byte, capacity/8)}
}

func (b *bitOutputStream) WriteBits(bitsCount, bits int) {
	bitNum := b.offset % 8
	byteNum := b.offset / 8

	firstWrite := min(8-bitNum, bitsCount)
	secondWrite := bitsCount - firstWrite

	b.buf[byteNum] |= byte(bits&((1<<firstWrite)-1)) << bitNum
	if secondWrite > 0 {
		b.buf[byteNum+1] |= byte(bits >> firstWrite)
	}

	b.offset += bitsCount
}

func (b *bitOutputStream) BitsToNextByte() int {
	if b.offset%8 == 0 {
		return 0
	}
	return 8 - b.offset%8
}

func (b *bitOutputStream) Bytes() []byte {
	size := b.offset / 8
	if b.offset%8 > 0 {
		size++
	}
	return b.buf[:size]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
