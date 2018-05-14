package bitVector

type BitVectorType struct {
  BitVector []byte
}

func (v *BitVectorType) toggleBit(pos uint) {
  v.BitVector[pos >> 3] ^= 1 << (pos - 1 % 8)
}

func (v *BitVectorType) getBit(pos uint) bool {
  if v.BitVector[pos >> 3] & (1 << (pos - 1 % 8)) == 0 {
    return false
  }
  return true
}

func (b *BitVectorType) Init(length int) *BitVectorType {
  b.BitVector = make([]byte, length >> 3)
  return b
}

func New(length int) *BitVectorType {
  return new(BitVectorType).Init(length)
}
