package bitmap

import (
	"fmt"
)

type Bitmap struct {
	len   int
	bytes []byte
}

// NewBitmap: get a new bitmap
func NewBitmap(length int) *Bitmap {
	l := length >> 3
	if length&7 != 0 {
		l++
	}
	return &Bitmap{len: length, bytes: make([]byte, l)}
}

func (b *Bitmap) index(num int) (int, int) {
	return num >> 3, num & 7
}

// add: add a num to the bitmap
func (b *Bitmap) Add(num int) bool {
	if num > b.len {
		return false
	}
	index, bit := b.index(num)
	b.bytes[index] = b.bytes[index] | 1<<bit
	return true
}

// Has: check is a num in the bitmap
func (b *Bitmap) Has(num int) bool {
	if num > b.len {
		return false
	}
	index, bit := b.index(num)
	return b.bytes[index]&(1<<bit) != 0
}

// Del: delete a num from the bitmap
func (b *Bitmap) Del(num int) bool {
	if num > b.len {
		return false
	}
	index, bit := b.index(num)
	b.bytes[index] = b.bytes[index] & ^(1 << bit)
	return true
}

// Len: get the length of the bitmap
func (b *Bitmap) Len() int {
	return b.len
}

// LenRaw: get the memory size of the bitmap array in bytes
func (b *Bitmap) LenRaw() int {
	return len(b.bytes)
}

// String: convert the bitmap to a binary string
func (b *Bitmap) String() string {
	bs := make([]byte, len(b.bytes))
	for i, v := range b.bytes {
		bs[i] = reverse(v)
	}
	s := fmt.Sprintf("%08b", bs)
	l :=b.len&7
	if l!=0{
		l=8-l
	}
	return s[1 : len(s)-1-l]
}

func reverse(c byte) byte {
	c = (c&0xaa)>>1 | (c&0x55)<<1
	c = (c&0xcc)>>2 | (c&0x33)<<2
	c = (c&0xf0)>>4 | (c&0x0f)<<4
	return c
}
